package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/devsquadron/models"
)

type TaskClient struct {
	UrlString   string
	ContentType string
}

func NewTaskClient(url string) *TaskClient {
	return &TaskClient{
		UrlString:   url,
		ContentType: "application/json;charset=UTF-8",
	}
}

func fromResponse[T any](res *http.Response) (*T, error) {
	var (
		err error
		tg  T
	)
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		err = json.NewDecoder(res.Body).Decode(&tg)
		if err != nil {
			return nil, err
		}
		return &tg, nil
	}
	return nil, errors.New(fmt.Sprintf("The response failed with status %s", res.Status))
}

func (clnt *TaskClient) CreateNewTask(tkn string, tsk *models.Task, tm string) error {
	var (
		err     error
		tskData []byte
		res     *http.Response
		req     *http.Request
	)

	tskData, err = json.Marshal(tsk)
	if err != nil {
		return err
	}
	tskCreateUrl, err := getUrl(clnt.UrlString, "/task/")
	if err != nil {
		return err
	}

	q := tskCreateUrl.Query()
	q.Set("team", tm)
	tskCreateUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodPost, tskCreateUrl.String(), bytes.NewBuffer(tskData))
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

func (clnt *TaskClient) UpdateTask(tkn string, tsk *models.Task, tm string) error {
	var (
		err     error
		tskData []byte
		res     *http.Response
		req     *http.Request
	)

	tskData, err = json.Marshal(tsk)
	if err != nil {
		return err
	}
	tskCreateUrl, err := getUrl(clnt.UrlString, "/task/")
	if err != nil {
		return err
	}

	q := tskCreateUrl.Query()
	q.Set("team", tm)
	tskCreateUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodPut, tskCreateUrl.String(), bytes.NewBuffer(tskData))
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

func (clnt *TaskClient) GetTasks(tkn string, tm string, tag string, statuses []string, dev string) (*[]models.Task, error) {
	var (
		err      error
		tsksUrl  *url.URL
		req      *http.Request
		res      *http.Response
		recdTsks *[]models.Task
	)

	tsksUrl, err = getUrl(clnt.UrlString, "/tasks/")
	if err != nil {
		return nil, err
	}
	q := tsksUrl.Query()
	q.Set("team", tm)
	q.Set("tag", tag)
	q.Set("developer", dev)
	for _, st := range statuses {
		q.Add("status", st)
	}
	tsksUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodGet, tsksUrl.String(), bytes.NewBuffer([]byte{}))
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

	recdTsks, err = fromResponse[[]models.Task](res)
	if err != nil {
		return nil, err
	}
	return recdTsks, nil
}

func (clnt *TaskClient) GetTaskById(tkn string, id uint64, tm string) (*models.Task, error) {
	var (
		err     error
		res     *http.Response
		req     *http.Request
		recdTsk *models.Task
	)

	tskUrl, err := getUrl(clnt.UrlString, "/task/")
	if err != nil {
		return nil, err
	}

	q := tskUrl.Query()
	q.Set("id", fmt.Sprintf("%d", id))
	q.Set("team", tm)
	tskUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodGet, tskUrl.String(), bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", tkn)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	recdTsk, err = fromResponse[models.Task](res)
	if err != nil {
		return nil, err
	}
	return recdTsk, nil
}

func (clnt *TaskClient) GetTagTaskDistribution(tkn string, tm string) (
	*[]models.TagDistribution, error,
) {
	var (
		err      error
		tsksUrl  *url.URL
		req      *http.Request
		res      *http.Response
		tgTskDst *[]models.TagDistribution
	)

	tsksUrl, err = getUrl(clnt.UrlString, "/tasks/tags/")
	if err != nil {
		return nil, err
	}
	q := tsksUrl.Query()
	q.Set("team", tm)
	tsksUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(http.MethodGet, tsksUrl.String(), bytes.NewBuffer([]byte{}))
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

	tgTskDst, err = fromResponse[[]models.TagDistribution](res)
	if err != nil {
		return nil, err
	}
	return tgTskDst, nil
}
