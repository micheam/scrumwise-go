package localdata

import (
	"context"
	"errors"

	"github.com/micheam/wiseman/internal/usecase"
	"github.com/micheam/wiseman/scrumwise"
)

type PersonRepository struct {
	data *DataCache
}

var _ usecase.PersonGateway = (*PersonRepository)(nil)

func NewPersonRepository(data *DataCache) *PersonRepository {
	return &PersonRepository{data: data}
}

func (p *PersonRepository) List(ctx context.Context) ([]*scrumwise.Person, error) {
	data := p.data.Current()
	if data == nil {
		return nil, errors.New("failed to get Current Data")
	}
	if data.Persons == nil {
		return nil, errors.New("data.Persons is nill")
	}
	return data.Persons, nil
}

func (p *PersonRepository) Get(ctx context.Context, id string) (*scrumwise.Person, error) {
	panic("Not Implemented Yet")
}
