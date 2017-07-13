package cistatus

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type Server struct {
	Logger *log.Logger

	JWT struct {
		Algorithm string
		Secret    []byte
	}

	*http.ServeMux
	wsHub *wsHub

	fetcher       Fetcher
	latestSummary Summary
}

func NewServer(fetcher Fetcher, fetchInterval time.Duration) *Server {
	now := time.Now()

	s := &Server{
		fetcher: fetcher,
		// Initial status summary is "unknown"
		latestSummary: Summary{
			Color:       Unknown,
			LastUpdated: &now,
		},
		// Default to discarding logs
		Logger: log.New(ioutil.Discard, "", 0),
		wsHub:  newWSHub(),
	}

	// Start WebSocket Hub
	go s.wsHub.run()

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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		err := errors.Wrap(err, "unable to create websocket connection to subscribe to status updates")
		s.Logger.Println(err)
		return
	}

	subscriber := newWSSubscriber(conn)
	s.wsHub.register <- subscriber
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
			s.wsHub.broadcast <- s.latestSummary
		}

		s.Logger.Printf("Fetched %d projects\n", len(projects))
		<-ticker
	}
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
