package greeting

import (
	"context"
	"fmt"
)

type Service interface {
	HelloMessage(ctx context.Context, name string) string
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) HelloMessage(ctx context.Context, name string) string {
	return fmt.Sprintf("Hello there, %s!", name)
}
