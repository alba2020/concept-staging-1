package dto

import "github.com/alba2020/arch/domain"

type AddRequest struct {
	X domain.First
	Y domain.Second
}

type AddResponse struct {
	Result domain.Result
}
