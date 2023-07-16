package backend

import (
	"net/http"
)

type ParentMapper struct {}

func (pm ParentMapper) getOriginURL(req http.Request) (*http.Request, error) {
	/*  receiver: ParentMapper
	*	parameters: 
	*		req -> http.Request
	*	returns:
	*		http.Request
	*		error
	*/
	var url string
	
	// fetch origin url from delivery service 
	// assign url to var::url

	newRequest, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return newRequest, nil
}
