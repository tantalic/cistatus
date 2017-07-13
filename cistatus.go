package cistatus

import (
	"time"
)

const (
	Version = "0.2.1"
)

type Fetcher interface {
	FetchStatus() ([]Project, error)
}

type Project struct {
	Name     string   `json:"name"`
	Branches []Branch `json:"branches,omitempty"`
}

func (p Project) String() string {
	return p.Name
}

type Branch struct {
	Name     string   `json:"name"`
	Commit   string   `json:"commit"`
	Statuses []Status `json:"statuses,omitempty"`
}

func (b Branch) String() string {
	return b.Name
}

type Status struct {
	Name    string    `json:"name"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
	Author  string    `json:"author"`
}

func (s Status) String() string {
	return s.Name
}

type Summary struct {
	Projects    []Project  `json:"projects,omitempty"`
	Color       Color      `json:"color"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

type Color string

const (
	Red     = Color("red")
	Yellow  = Color("yellow")
	Green   = Color("green")
	Unknown = Color("question")
)
