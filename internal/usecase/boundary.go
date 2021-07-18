package usecase

import (
	"context"
	"errors"

	"github.com/micheam/wiseman/scrumwise"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrIllegalArgument = errors.New("illegal argument")
)

//go:generate mockgen -source boundary.go -destination boundary_mocks.go -package usecase

type PersonGateway interface {
	List(ctx context.Context) ([]*scrumwise.Person, error)
	Get(ctx context.Context, id string) (*scrumwise.Person, error)
}
