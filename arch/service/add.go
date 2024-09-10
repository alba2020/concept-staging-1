package service

import (
	"github.com/alba2020/arch/domain"
	"github.com/alba2020/arch/dto"
)

type Repo interface {
	Get(x domain.First, y domain.Second) (domain.Result, bool)
	Set(x domain.First, y domain.Second, val domain.Result) error
}

type AddService struct {
	repo Repo
}

func New(r Repo) *AddService {
	return &AddService{
		repo: r,
	}
}

func (s *AddService) Add(req dto.AddRequest) dto.AddResponse {
	if val, ok := s.repo.Get(req.X, req.Y); ok {
		return dto.AddResponse{Result: val}
	}

	res := req.X.Add(req.Y)
	s.repo.Set(req.X, req.Y, res)

	return dto.AddResponse{Result: res}
}

