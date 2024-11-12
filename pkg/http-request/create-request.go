package httprequest

import "net/http"

var requestURL = "https://api-create.runsystem.info"

type CreateRequest struct {
}

func MakeCreateRequest(path string) {
	req, err := http.NewRequest(http.MethodGet, requestURL+path, nil)
}

func (c CreateRequest) Checkin() {

}
