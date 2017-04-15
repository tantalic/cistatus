package cistatus

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"sync"

	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Server struct {
	Logger *log.Logger

	JWT struct {
		Algorithm string
		Secret    []byte
	}

	*http.ServeMux

	fetcher       Fetcher
	latestSummary Summary
	websocket     struct {
		upgrader    websocket.Upgrader
		mutex       sync.RWMutex
		connections map[*websocket.Conn]bool
		broadcast   chan Summary
	}
}

func NewServer(fetcher Fetcher, fetchInterval time.Duration) *Server {
	now := time.Now()

	s := &Server{
		fetcher: fetcher,
		// Initial status summary is "unkown"
		latestSummary: Summary{
			Color:       Unknown,
			LastUpdated: &now,
		},
		// Default to discarding logs
		Logger: log.New(ioutil.Discard, "", 0),
	}

	// Setup for websockets
	s.websocket.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	s.websocket.connections = make(map[*websocket.Conn]bool)
	s.websocket.broadcast = make(chan Summary)
	go s.handleBroadcasts()

	// Create servemux with routes to http api
	s.ServeMux = http.NewServeMux()
	s.ServeMux.HandleFunc("/api", s.allProjects)
	s.ServeMux.HandleFunc("/api/watch", s.websocketSubscribeHandler)

	// Start fetching
	go s.fetchLoop(fetchInterval)

	return s
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
	if token.Method.Alg() != s.JWT.Algorithm {
		return nil, errors.Errorf("unexpected jwt algorithm: %v", token.Header["alg"])
	}

	return s.JWT.Secret, nil
}

func (s *Server) isAuthorized(r *http.Request) bool {
	// If the JWT algorithm OR secret is not set then all requests are authorized
	if s.JWT.Secret == nil || s.JWT.Algorithm == "" {
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
		s.Logger.Printf("JWT error: %s\n", err)
		return false
	}

	return token.Valid
}

func (s *Server) websocketSubscribeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.websocket.upgrader.Upgrade(w, r, nil)
	if err != nil {
		err := errors.Wrap(err, "unable to create websocket connection to subscribe to status updates")
		s.Logger.Println(err)
		return
	}

	s.addSubscriber(conn)
}

func (s *Server) fetchLoop(interval time.Duration) {
	ticker := time.Tick(interval)
	errCount := 0

	for {
		s.Logger.Println("Fetching CI server status")

		projects, err := s.fetcher.FetchStatus()
		if err != nil {
			s.Logger.Printf("Error fetching status: %s\n", err)
			errCount++
			if errCount >= 10 {

			}
			<-ticker
			continue
		}

		now := time.Now()
		s.latestSummary.Projects = projects
		s.latestSummary.LastUpdated = &now

		newColor := color(projects)
		if newColor != s.latestSummary.Color {
			s.latestSummary.Color = newColor
			s.websocket.broadcast <- s.latestSummary
		}

		s.Logger.Printf("Fetched %d projects\n", len(projects))
		<-ticker
	}
}

func (s *Server) handleBroadcasts() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-s.websocket.broadcast
		msg.Projects = nil

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
