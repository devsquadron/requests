package requests

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/devsquadron/models"
)

type TeamClient struct {
	UrlString   string
	ContentType string
}

func NewTeamClient(url string) *TeamClient {
	return &TeamClient{
		UrlString:   url,
		ContentType: "application/json;charset=UTF-8",
	}
}

func (clnt *TeamClient) CreateNewTeam(tm *models.Team, tkn string) error {
	var (
		err    error
		req    *http.Request
		tmData []byte
		res    *http.Response
	)

	tmData, err = json.Marshal(tm)
	if err != nil {
		return err
	}
	teamUrl, err := getUrl(clnt.UrlString, "/team/")
	if err != nil {
		return err
	}

	req, err = http.NewRequest(http.MethodPost, teamUrl.String(), bytes.NewBuffer(tmData))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", tkn)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return getErrorFromResponse(res)
	}
	return nil
}

func (clnt *TeamClient) GrowTeam(tkn string, tm string, dev *models.Developer) error {
	var (
		err     error
		req     *http.Request
		devData []byte
		res     *http.Response
	)

	devData, err = json.Marshal(dev)
	if err != nil {
		return err
	}
	growTmUrl, err := getUrl(clnt.UrlString, "/team/grow/")
	if err != nil {
		return err
	}
	q := growTmUrl.Query()
	q.Set("team", tm)
	growTmUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodPost, growTmUrl.String(), bytes.NewBuffer(devData))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", tkn)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return getErrorFromResponse(res)
	}
	return nil
}

func (clnt *TeamClient) InfoTeam(tkn string, tm string) (*models.Team, error) {
	var (
		err    error
		req    *http.Request
		res    *http.Response
		recdTm *models.Team
	)

	infoTmUrl, err := getUrl(clnt.UrlString, "/team/info/")
	if err != nil {
		return nil, err
	}
	q := infoTmUrl.Query()
	q.Set("team", tm)
	infoTmUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodGet, infoTmUrl.String(), bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", tkn)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, getErrorFromResponse(res)
	}

	recdTm, err = fromResponse[models.Team](res)
	return recdTm, nil
}
