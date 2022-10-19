package totp

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
)

const secret = "I65VU7K5ZQL7WB4E"

func TestTOTP(t *testing.T) {
	url := url.URL{
		Scheme: "https",
		Host:   "authenticationtest.com",
		Path:   "/totp",
	}

	q := url.Query()
	q.Set("secret", secret)
	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		t.Error(err.Error())
	}

	tot := TOTP{}
	err = json.NewDecoder(resp.Body).Decode(&tot)
	if err != nil {
		t.Error(err.Error())
	}

	code, err := GenerateTotp(secret)
	if err != nil {
		t.Error(err.Error())
	}

	if code != tot.Code {
		t.Error(err.Error())
	}
}

type TOTP struct {
	Code int32 `json:"code,string"`
}
