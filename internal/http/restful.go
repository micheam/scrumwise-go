package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micheam/wiseman/internal/usecase"
	"github.com/micheam/wiseman/scrumwise"
)

type PersonController struct {
	UsecaseListPersons *usecase.UseCaseListPersons
	UsecaseGetPerson   *usecase.UseCaseGetPerson
}

func (p *PersonController) ListPersons(w http.ResponseWriter, r *http.Request) {
	err := p.UsecaseListPersons.Execute(r.Context(), func(_ context.Context, ps []*scrumwise.Person) error {
		b, err := json.Marshal(ListPersonsResponse{
			Persons: ps,
		})
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}
		w.WriteHeader(200)
		w.Header().Add("content-type", "application/json")
		_, err = w.Write(b)
		return err
	})
	if err != nil {
		log.Println("failed to execute list-persons: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *PersonController) GetPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	log.Println("id: ", vars["id"])
}

// ------
// Models
// ------

type ListPersonsResponse struct {
	Persons []*scrumwise.Person `json:"persons"`
}
