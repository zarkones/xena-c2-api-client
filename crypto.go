package c2api

import (
	"io"
	"net/http"
)

func GetC2PublicKey() (pubKeyPEM []byte, err error) {
	if len(*AuthToken) == 0 {
		return nil, ErrNilAuthToken
	}

	req, err := http.NewRequest(http.MethodGet, *BaseURL+"/v1/public-key", nil)
	if err != nil {
		return nil, err
	}

	setAuth(req)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}
