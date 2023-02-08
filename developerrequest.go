package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/devsquadron/models"
)

type DeveloperClient struct {
	UrlString   string
	ContentType string
}

func NewDeveloperClient(url string) *DeveloperClient {
	return &DeveloperClient{
		UrlString:   url,
		ContentType: "application/json;charset=UTF-8",
	}
}

func getTokenFromResponse(res *http.Response) (string, error) {
	var (
		err    error
		tknRes models.TokenResponse
	)
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		err = json.NewDecoder(res.Body).Decode(&tknRes)
		if err != nil {
			return "", err
		}
		return tknRes.Token, nil
	}
	return "", errors.New(fmt.Sprintf("The response failed with status %s", res.Status))
}

func (clnt *DeveloperClient) CreateNewDeveloper(dev *models.Developer) (string, error) {
	var (
		err     error
		devData []byte
		res     *http.Response
		newTkn  string
	)

	devData, err = json.Marshal(dev)
	if err != nil {
		return "", err
	}
	devCreateUrl, err := getUrl(clnt.UrlString, "/developer/")
	if err != nil {
		return "", err
	}
	res, err = http.Post(devCreateUrl.String(), clnt.ContentType, bytes.NewBuffer(devData))
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", getErrorFromResponse(res)
	}
	newTkn, err = getTokenFromResponse(res)
	if err != nil {
		return "", err
	}
	return newTkn, nil
}

func (clnt *DeveloperClient) LoginDeveloper(dev *models.Developer) (string, error) {
	var (
		err     error
		devData []byte
		res     *http.Response
		newTkn  string
	)

	devData, err = json.Marshal(dev)
	if err != nil {
		return "", err
	}
	devLoginUrl, err := getUrl(clnt.UrlString, "/developer/login/")
	if err != nil {
		return "", err
	}
	res, err = http.Post(devLoginUrl.String(), clnt.ContentType, bytes.NewBuffer(devData))
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", getErrorFromResponse(res)
	}
	newTkn, err = getTokenFromResponse(res)
	if err != nil {
		return "", err
	}
	return newTkn, nil
}
