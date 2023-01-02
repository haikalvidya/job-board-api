package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetRequestJob(url string, urlJobDetail string, params map[string]string) (string, error) {
	var requestURL string
	if params["job_id"] != "" {
		requestURL = urlJobDetail
	} else {
		delete(params, "job_id")
		delete(params, "full_time")
		requestURL = url
	}
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", fmt.Errorf("client: could not create request: %s", err)
	}

	if params["job_id"] != "" {
		req.URL.Path += "/" + params["job_id"]
	} else {
		for key, value := range params {
			q := req.URL.Query()
			q.Add(key, value)
			req.URL.RawQuery = q.Encode()
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("client: error making http request: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("client: unexpected status code: %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("client: could not read response body: %s", err)
	}
	return string(resBody), nil
}
