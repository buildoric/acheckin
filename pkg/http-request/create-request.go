package httprequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/buildoric/acheckin/pkg/config"
)

var requestURL = "https://api-create.runsystem.info"

type CreateRequest struct {
	Config *config.Config
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Data    struct {
		UserObjId string `json:"userObjId"`
	} `json:"data"`
}

type CanCheckinResponse struct {
	Data struct {
		TimeKeepingMonth struct {
			TimeCheckIn  string `json:"timeCheckIn"`
			TimeCheckOut string `json:"timeCheckOut"`
		} `json:"timeKeepingMonth"`
	} `json:"data"`
}

func MakeCreateRequest(method string, path string, body []byte, accessToken string) (*http.Response, error) {
	var bodyReader *bytes.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)

	}

	// var requestURL

	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(method, requestURL+path, nil)
	} else {
		req, err = http.NewRequest(method, requestURL+path, bodyReader)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic ZHhpbnRlcm5hbF9wbDpnb0R4QDIwMjE=")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Access-Token", accessToken)
	return http.DefaultClient.Do(req)

}

func (c CreateRequest) Login() *LoginResponse {
	loginRequest := &LoginRequest{
		Username: c.Config.Create.Username,
		Password: c.Config.Create.Password,
	}

	jsonBody, err := json.Marshal(loginRequest)

	if err != nil {
		return nil
	}

	res, err := MakeCreateRequest(http.MethodPost, "/signIn", jsonBody, "")

	if err != nil {
		return nil
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		return nil
	}

	var user *LoginResponse
	json.Unmarshal(resBody, &user)

	return user
}

func (c CreateRequest) Checkin(accessToken string) {
	now := time.Now()
	monthAt := now.Format("012006")
	recordTime := now.Format("2006-01-02 15:04:05")
	d := make(map[string]string)

	d["monthAt"] = monthAt
	d["recordTime"] = recordTime
	jsonBody, _ := json.Marshal(d)
	res, err := MakeCreateRequest(http.MethodPost, "/auth/time-keepings/checkin", jsonBody, accessToken)

	if err != nil {
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
	}

	log.Printf("data = %s", resBody)
}

// https://api-create.runsystem.info/auth/time-keepings/canCheckIn?userObjId=60af0cdc3ca32b4f1ce7d46e

func (c CreateRequest) CanCheckIn(accessToken, userObjId string) (*CanCheckinResponse, error) {
	path := fmt.Sprintf("/auth/time-keepings/canCheckIn?userObjId=%s", userObjId)

	res, err := MakeCreateRequest(http.MethodGet, path, nil, accessToken)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var d *CanCheckinResponse
	json.Unmarshal(resBody, &d)

	return d, nil
}
