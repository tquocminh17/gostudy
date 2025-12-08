package user

import (
	"errors"
	"fmt"

	"github.com/tquocminh17/goapi/pkg/models"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(req models.RegisterRequest) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("Username and Password are required")
	}

	if req.Username == "admin" {
		return errors.New("Username is reserved")
	}

	fmt.Printf("Successfully registered user: %s\n", req.Username)

	return nil
}
