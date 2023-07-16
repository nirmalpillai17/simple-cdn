package backend

import (
	"net/http"
)

type BEWorker interface {}

func Do(req http.Request) (http.Response, error) {

	return resp, nil
}
