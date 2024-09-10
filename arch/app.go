package main

import (
	"fmt"

	"github.com/alba2020/arch/ctrl"
	"github.com/alba2020/arch/repo"
	"github.com/alba2020/arch/service"
)

type Applicaton struct {
	ctrl *ctrl.Controller
}

func NewApplication() *Applicaton {
	repo := repo.NewNumbersRepo()
	svc := service.New(repo)
	ctrl := ctrl.New(svc)

	return &Applicaton{ctrl}
}

func (a *Applicaton) Run() {
	x := 21
	y := 2

	result := a.ctrl.Handle(x, y)

	fmt.Println(result)
}
