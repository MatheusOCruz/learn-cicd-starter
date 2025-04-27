package auth

import (
	"net/http"
	"testing"
)

func TestAuthNoHeader(t *testing.T) {
	h := make(http.Header)
	_, err := GetAPIKey((h))
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("not returning error")
	}
}

func TestNoApiKey(t *testing.T) {
	h := make(http.Header)
	h["Authorization"] = []string{"Rusbe da Silva"}
	_, err := GetAPIKey((h))
	if err == nil {
		t.Errorf("malformed header passed")
	}

}

func TestApiKey(t *testing.T) {
	h := make(http.Header)
	h["Authorization"] = []string{"ApiKey junin"}
	res, err := GetAPIKey((h))
	if err != nil || res != "junin" {
		t.Errorf("junin foi esquecido")
	}

}

