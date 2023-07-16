package backend

import (
	"net/http"
)

type ParentFetcher struct {}

func (pf ParentFetcher) getObject(req *http.Request) (*http.Response, error) {
	/*  receiver: ParentFetcher
	*	parameters:
	*		req -> *http.Request
	*	returns:
	*		*http.Response
	*		error
	*/
	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		return nil, err
	}
	
	return resp, err
}
