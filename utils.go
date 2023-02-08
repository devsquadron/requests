package requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/devsquadron/models"
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
