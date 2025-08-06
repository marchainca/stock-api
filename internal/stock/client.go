package stock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseURL  string
	user     string
	password string
	httpc    *http.Client
}

func NewClient(baseURL, user, pass string, hc *http.Client) *Client {
	if hc == nil {
		hc = &http.Client{Timeout: 10 * time.Second}
	}
	baseURL = strings.TrimRight(baseURL, "/")
	return &Client{baseURL: baseURL, user: user, password: pass, httpc: hc}
}

// Login obtiene un JWT y lo devuelve.
func (c *Client) Login(ctx context.Context) (string, error) {
	body, _ := json.Marshal(map[string]string{
		"username": c.user,
		"password": c.password,
	})
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("%s/login", c.baseURL),
		bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	var res *http.Response
	var err error
	for i := 0; i < 3; i++ { // retry simple
		res, err = c.httpc.Do(req)
		if err == nil && res.StatusCode < 500 {
			break
		}
		time.Sleep(time.Duration(i+1) * 300 * time.Millisecond)
	}
	if err != nil {
		return "", fmt.Errorf("login request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed: %s", res.Status)
	}

	var lr loginResponse
	if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
		return "", fmt.Errorf("decode login: %w", err)
	}
	return lr.Token, nil
}

// List trae los ítems.
func (c *Client) List(ctx context.Context, token, cursor string) (listResponse, error) {
	url := fmt.Sprintf("%s/list?next_page=%s", c.baseURL, cursor)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpc.Do(req)
	if err != nil {
		return listResponse{}, fmt.Errorf("list request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return listResponse{}, fmt.Errorf("list failed: %s", res.Status)
	}

	var lr listResponse
	if err := json.NewDecoder(res.Body).Decode(&lr); err != nil {
		return listResponse{}, fmt.Errorf("decode list: %w", err)
	}
	return lr, nil
}
