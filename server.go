package cistatus

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Server struct {
	Fetcher Fetcher
	Logger  Logger

	mux           *http.ServeMux
	upgrader      websocket.Upgrader
	latestSummary Summary

	JWTAlgorithm string
	JWTSecret    []byte

	websocket struct {
		mutex       sync.RWMutex
		connections map[*websocket.Conn]bool
		broadcast   chan Summary
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.mux == nil {
		s.init()
	}

	s.mux.ServeHTTP(w, r)
}

func (s *Server) init() {
	// Set default color (if not defined)
	if s.latestSummary.Color == "" {
		now := time.Now()
		s.latestSummary.Color = Unkown
		s.latestSummary.LastUpdated = &now
	}

	// Setup for websockets
	s.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	go s.handleBroadcasts()

	// Create servemux with routes to http api
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/api", s.allProjects)
	s.mux.HandleFunc("/api/watch", s.websocketSubscribeHandler)
}

func (s *Server) allProjects(w http.ResponseWriter, r *http.Request) {
	latestSummary := s.latestSummary

	if !s.isAuthorized(r) {
		latestSummary.Projects = []Project{}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Server", "https://github.com/tantalic/cistatus")
	w.Header().Set("X-Server-Version", Version)
	json.NewEncoder(w).Encode(latestSummary)
}

func (s *Server) jwtKey(token *jwt.Token) (interface{}, error) {
	if token.Method.Alg() != s.JWTAlgorithm {
		return nil, errors.Errorf("unexpected jwt algorithm: %v", token.Header["alg"])
	}

	return s.JWTSecret, nil
}

func (s *Server) isAuthorized(r *http.Request) bool {
	// If the JWT algorithm OR secret is not set then all requests are authorized
	if s.JWTSecret == nil || s.JWTAlgorithm == "" {
		return true
	}

	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		// no header present == not authorized
		return false
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "bearer" {
		// malformed header == not authorized, token should be in format of "bearer XXXXXXXXXXXXXXXX"
		return false
	}

	token, err := jwt.Parse(headerParts[1], s.jwtKey)
	if err != nil {
		s.logf("JWT error: %s", err)
		return false
	}

	return token.Valid
}

func (s *Server) websocketSubscribeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		err := errors.Wrap(err, "unable to create websocket connection to subscribe to status updates")
		s.log(err)
		return
	}

	s.addSubscriber(conn)
}

func (s *Server) StartFetching(interval time.Duration) {
	go func() {
		ticker := time.Tick(interval)

		for {
			s.log("Fetching CI server status")

			projects, err := s.Fetcher.FetchStatus()
			if err != nil {
				s.logf("Error fetching status: %s\n", err)
			}

			now := time.Now()
			s.latestSummary.Projects = projects
			s.latestSummary.LastUpdated = &now

			newColor := color(projects)
			if newColor != s.latestSummary.Color {
				s.latestSummary.Color = newColor
				s.websocket.broadcast <- s.latestSummary
			}

			s.logf("Fetched %d projects", len(projects))
			<-ticker
		}

	}()
}

func (s *Server) handleBroadcasts() {
	s.websocket.connections = make(map[*websocket.Conn]bool)
	s.websocket.broadcast = make(chan Summary)

	for {
		// Grab the next message from the broadcast channel
		msg := <-s.websocket.broadcast
		// Send it out to every client that is currently connected
		s.websocket.mutex.RLock()
		for conn := range s.websocket.connections {
			err := conn.WriteJSON(msg)
			if err != nil {
				s.removeSubscriber(conn)
			}
		}
		s.websocket.mutex.RUnlock()
	}
}

func (s *Server) addSubscriber(conn *websocket.Conn) {
	// Send current status immediately
	summary := s.latestSummary
	summary.Projects = nil
	conn.WriteJSON(summary)

	// Add to list of connections
	s.websocket.mutex.Lock()
	s.websocket.connections[conn] = true
	s.websocket.mutex.Unlock()
}

func (s *Server) removeSubscriber(conn *websocket.Conn) {
	s.websocket.mutex.Lock()
	delete(s.websocket.connections, conn)
	s.websocket.mutex.Unlock()

	conn.Close()
}

func (s *Server) log(a ...interface{}) {
	if s.Logger == nil {
		return
	}

	s.Logger.Log(a...)
}

func (s *Server) logf(format string, a ...interface{}) {
	if s.Logger == nil {
		return
	}

	s.Logger.Logf(format, a...)
}

func color(projects []Project) Color {
	color := Green

	for _, project := range projects {
		for _, branch := range project.Branches {
			for _, status := range branch.Statuses {

				// If any status is failed return red immediately
				if status.Status == "failed" {
					return Red
				}

				// If any status is pending or running return yellow
				if status.Status == "pending" || status.Status == "running" {
					color = Yellow
				}

			}
		}
	}

	return color
}
