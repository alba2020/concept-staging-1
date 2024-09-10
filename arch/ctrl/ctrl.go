package ctrl

import (
	"github.com/alba2020/arch/domain"
	"github.com/alba2020/arch/dto"
)

type Adder interface {
	Add(dto.AddRequest) dto.AddResponse
}

type Controller struct {
	svc Adder
}

func New(a Adder) *Controller {
	return &Controller{
		svc: a,
	}
}

func (c *Controller) Handle(x, y int) int {
	req := dto.AddRequest{
		X: domain.First(x),
		Y: domain.Second(y),
	}

	res := c.svc.Add(req)
	
	return int(res.Result)
}
