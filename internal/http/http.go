package http

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/micheam/wiseman/internal/usecase"
)

var Version = "v0.0.1"

type Container struct {
	UseCaseListPersons *usecase.UseCaseListPersons
	UseCaseGetPerson   *usecase.UseCaseGetPerson
}

func Router(_ context.Context, c Container) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "text/plain")
		_, _ = w.Write(banner(Version))
	})

	// person
	ctrl := &PersonController{c.UseCaseListPersons, c.UseCaseGetPerson}
	r.HandleFunc("/person", ctrl.ListPersons)
	r.HandleFunc("/person/{id:[0-9]+-[0-9]+-[0-9]}", ctrl.GetPerson)

	return r
}

func dumpRequest(r *http.Request) (data map[string]interface{}, err error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"body":   string(b),
		"header": r.Header,
	}, nil
}

func banner(version string) []byte {
	sb := new(strings.Builder)
	sb.WriteString(`
██╗    ██╗██╗███████╗███████╗███╗   ███╗ █████╗ ███╗   ██╗
██║    ██║██║██╔════╝██╔════╝████╗ ████║██╔══██╗████╗  ██║
██║ █╗ ██║██║███████╗█████╗  ██╔████╔██║███████║██╔██╗ ██║
██║███╗██║██║╚════██║██╔══╝  ██║╚██╔╝██║██╔══██║██║╚██╗██║
╚███╔███╔╝██║███████║███████╗██║ ╚═╝ ██║██║  ██║██║ ╚████║
 ╚══╝╚══╝ ╚═╝╚══════╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝
`)
	for i := 0; i < (58 - len(version)); i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(version)
	return []byte(sb.String())
}
