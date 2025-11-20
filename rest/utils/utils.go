package utils

import (
	"io/ioutil"
	"net/http"

	brocade_errors "github.com/projectbadger/brocade-go/rest/errors"
	"github.com/projectbadger/brocade-go/session"
	"github.com/projectbadger/brocade-go/utils"
)

func GetResponse(method, url string, body []byte, sess session.Session, contentType utils.ContentType, client utils.RequestClient) (*http.Response, brocade_errors.BrocadeErr) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, brocade_errors.New(err.Error())
	}
	err = sess.HandleRequest(req)
	if err != nil {
		return nil, brocade_errors.New(err.Error())
	}
	contentType.SetRequestHeaders(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, brocade_errors.New(err.Error())
	}
	defer resp.Body.Close()
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, brocade_errors.New(err.Error())
	}
	errs := CheckBodyForErrors(responseBytes, contentType)
	if respErrs := CheckResponseForErrors(resp); len(respErrs.Errors) > 0 {
		for ix := range respErrs.Errors {
			errs.AddError(respErrs.Errors[ix])
		}
	}
	// Errors returned
	return resp, brocade_errors.NewFromErrors(errs.Errors...)
}

func CheckBodyForErrors(body []byte, contentType utils.ContentType) *brocade_errors.Errors {
	var errs brocade_errors.Errors
	err := contentType.Unmarshal(body, &errs)
	if err != nil {
		return nil
	}
	return &errs
}

func CheckResponseForErrors(resp *http.Response) *brocade_errors.Errors {
	var errs brocade_errors.Errors
	if resp.StatusCode >= 400 {
		errs.AddError(brocade_errors.Error{
			ErrorType:    "http-status-error",
			ErrorMessage: brocade_errors.ErrHttpResponseStatus.Error() + ": " + resp.Status,
			ErrorInfo: brocade_errors.ErrorInfo{
				ErrorCode:   resp.StatusCode,
				ErrorModule: "http",
			},
		})
	}
	return &errs
}

func GetResponseBody(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}
