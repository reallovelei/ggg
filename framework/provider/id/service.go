package id

import (
	"github.com/rs/xid"
)

type GGGIDService struct {
}

func NewGGGIDService(params ...interface{}) (interface{}, error) {
	return &GGGIDService{}, nil
}

func (s *GGGIDService) NewID() string {
	return xid.New().String()
}
