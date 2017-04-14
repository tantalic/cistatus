package cistatus

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Client struct {
	HTTPClient http.Client
	Token      string

	Hostname string
	Port     int
	UseTLS   bool
}

func (c *Client) Summary(ctx context.Context) (*Summary, error) {
	URL, err := c.summaryURL()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	if c.Token != "" {
		auth := fmt.Sprintf("bearer %s", c.Token)
		req.Header.Add("Authorization", auth)
	}
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var summary Summary
	err = json.NewDecoder(resp.Body).Decode(&summary)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (c *Client) summaryURL() (string, error) {
	if c.Hostname == "" {
		return "", errors.New("hostname must be set on cistatus.Client")
	}

	var scheme string
	if c.UseTLS {
		scheme = "https"
	} else {
		scheme = "http"
	}

	if c.Port != 0 {
		return fmt.Sprintf("%s://%s:%d/api", scheme, c.Hostname, c.Port), nil
	}

	return fmt.Sprintf("%s://%s/api", scheme, c.Hostname), nil
}

func (c *Client) Watch(summChan chan Summary) error {
	URL, err := c.watchURL()
	if err != nil {
		return err
	}

	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(URL, nil)
	if err != nil {
		return err
	}

	for {
		var summary Summary
		err := conn.ReadJSON(&summary)
		if err != nil {
			conn.Close()
			return err
		}

		summChan <- summary
	}

}

func (c *Client) watchURL() (string, error) {
	if c.Hostname == "" {
		return "", errors.New("hostname must be set on cistatus.Client")
	}

	var scheme string
	if c.UseTLS {
		scheme = "wss"
	} else {
		scheme = "ws"
	}

	if c.Port != 0 {
		return fmt.Sprintf("%s://%s:%d/api/watch", scheme, c.Hostname, c.Port), nil
	}

	return fmt.Sprintf("%s://%s/api/watch", scheme, c.Hostname), nil
}
