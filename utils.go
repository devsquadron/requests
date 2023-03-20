package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/devsquadron/models"
)

var (
	EMPTY_BYTE_ARRAY = bytes.NewBuffer([]byte{})
)

func getUrl(base string, ext string) (*url.URL, error) {
	url, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	extUrl := url.JoinPath(ext)
	return extUrl, nil
}

func getErrorFromResponse(res *http.Response) error {
	var (
		err    error
		resErr models.ErrorResponse
	)
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&resErr)
	if err != nil {
		return err
	}
	return errors.New(resErr.Error)
}
