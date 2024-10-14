package resource

import (
	"context"
	"errors"
)

type Service struct {
}

func NewService() *Service {
	return new(Service)
}

func (s Service) GetData(ctx context.Context) (map[string]any, error) {
	if ctx == nil {
		return nil, errors.New("context is nil")
	}
	return map[string]any{"message": "Hello World"}, nil
}
