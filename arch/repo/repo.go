package repo

import (
	"fmt"

	"github.com/alba2020/arch/domain"
)

type NumbersRepo struct {
	cache map[string]domain.Result
}

func NewNumbersRepo() *NumbersRepo {
	return &NumbersRepo{
		cache: make(map[string]domain.Result),
	}
}

func key(x domain.First, y domain.Second) string {
	return fmt.Sprintf("%d %d", x, y)
}


func (r *NumbersRepo) Get(x domain.First, y domain.Second) (domain.Result, bool) {
	val, ok := r.cache[key(x, y)]

	return val, ok
}

func (r *NumbersRepo) Set(x domain.First, y domain.Second, val domain.Result) error {
	r.cache[key(x, y)] = val
	return nil
}
