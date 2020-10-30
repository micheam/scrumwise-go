package scrumwise

import "fmt"

const (
	baseURL    string = "https://api.scrumwise.com/service/api"
	apiVersion string = "v1"
)

func endpoint(resource string) string {
	return fmt.Sprintf("$s/$s/$s", baseURL, apiVersion, resource)
}
