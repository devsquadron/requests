package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

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

func (clnt *TeamClient) JoinTeam(tkn string, tm string) error {
	var (
		err       error
		req       *http.Request
		res       *http.Response
		joinTmUrl *url.URL
	)

	joinTmUrl, err = getUrl(clnt.UrlString, "/team/join/")
	if err != nil {
		return err
	}
	q := joinTmUrl.Query()
	q.Set("team", tm)
	joinTmUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodPost, joinTmUrl.String(), EMPTY_BYTE_ARRAY)
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

func (clnt *TeamClient) RespondJoinRequest(tkn string, tm string, resJnReq *models.RespondJoinRequestReq) error {
	var (
		err          error
		req          *http.Request
		res          *http.Response
		resJnUrl     *url.URL
		resJnReqData []byte
	)

	resJnReqData, err = json.Marshal(resJnReq)
	if err != nil {
		return err
	}

	resJnUrl, err = getUrl(clnt.UrlString, "/team/respond-join/")
	if err != nil {
		return err
	}
	q := resJnUrl.Query()
	q.Set("team", tm)
	resJnUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodPost, resJnUrl.String(), bytes.NewBuffer(resJnReqData))
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

	req, err = http.NewRequest(http.MethodGet, infoTmUrl.String(), EMPTY_BYTE_ARRAY)
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
