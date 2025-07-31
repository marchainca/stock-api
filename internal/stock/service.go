package stock

import (
	"context"
	"sync"
	"time"
)

type Service struct {
	cli          *Client
	mu           sync.Mutex
	token        string
	tokenFetched time.Time
}

func NewService(c *Client) *Service { return &Service{cli: c} }

func (s *Service) ensureToken(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Re-login cada 55 min (o ajusta según exp del JWT)
	if time.Since(s.tokenFetched) < 55*time.Minute && s.token != "" {
		return s.token, nil
	}
	tk, err := s.cli.Login(ctx)
	if err != nil {
		return "", err
	}
	s.token = tk
	s.tokenFetched = time.Now()
	return tk, nil
}

func (s *Service) List(ctx context.Context, cursor string) (listResponse, error) {
	tk, err := s.ensureToken(ctx)
	if err != nil {
		return listResponse{}, err
	}
	return s.cli.List(ctx, tk, cursor)
}
