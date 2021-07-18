package usecase

import (
	"context"
	"fmt"

	"github.com/micheam/wiseman/scrumwise"
)

type (
	ListPersonsHandler func(ctx context.Context, ps []*scrumwise.Person) error
	PersonHandler      func(ctx context.Context, p scrumwise.Person) error
)

type UseCaseListPersons struct{ persons PersonGateway }

func NewUseCaseListPersons(persons PersonGateway) *UseCaseListPersons {
	return &UseCaseListPersons{persons}
}

// Execute retrieves active persons from scrumwise, then outputs them.
// This may return an empty array if any persons are found.
func (u *UseCaseListPersons) Execute(ctx context.Context, output ListPersonsHandler) error {
	ps, err := u.persons.List(ctx)
	if err != nil {
		return fmt.Errorf("failed to retreive existing persons: %w", err)
	}
	return output(ctx, ps)
}

type (
	UseCaseGetPerson struct{ persons PersonGateway }
	GetPersonInput   struct{ PersonID string }
)

// Execute gets Person by id, then outputs it.
// This may return ErrNotFound if no person is found with a specified id.
func (u *UseCaseGetPerson) Execute(ctx context.Context, input GetPersonInput, output PersonHandler) error {
	id := input.PersonID
	if len(id) == 0 {
		return fmt.Errorf("empty PersionID: %w", ErrIllegalArgument)
	}

	p, err := u.persons.Get(ctx, input.PersonID)
	if err != nil {
		return fmt.Errorf("failed to retreive existing persons: %w", err)
	}
	return output(ctx, *p)
}
