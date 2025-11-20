package api_interface

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	brocade_errors "github.com/projectbadger/brocade-go/rest/errors"
	"github.com/projectbadger/brocade-go/utils"
)

func GetHTTPResponse(req *http.Request, config *RESTConfig) (*http.Response, brocade_errors.BrocadeErr) {
	err := config.HandleRequest(req)
	if err != nil {
		return nil, brocade_errors.NewFromErr(err)
	}
	// fmt.Printf("Request before being sent out: '%#v' to URL: '%s'\n", req, req.URL.String())
	// return nil, errors.New("stop here before the request")
	resp, err := config.Client().Do(req)
	if err != nil {
		return nil, brocade_errors.NewFromErr(err)
	}
	return resp, nil
}

func GetResponse(req *http.Request, config *RESTConfig) (*http.Response, []byte, brocade_errors.BrocadeErr) {
	err := config.HandleRequest(req)
	if err != nil {
		return nil, nil, brocade_errors.New(err.Error())
	}
	resp, err := config.Client().Do(req)
	if err != nil {
		return nil, nil, brocade_errors.New(err.Error())
	}
	// fmt.Println("Got Response from client:", resp)
	defer resp.Body.Close()
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, brocade_errors.New(err.Error())
	}
	// fmt.Println("Got response body from client:", string(responseBytes))
	errs := CheckBodyForErrors(responseBytes, config.ContentType())
	if respErrs := CheckResponseForErrors(resp); len(respErrs.Errors) > 0 {
		if errs == nil {
			errs = &brocade_errors.Errors{}
		}
		for ix := range respErrs.Errors {
			errs.AddError(respErrs.Errors[ix])
		}
	}
	// if err == nil {
	// Errors returned
	return resp, responseBytes, errs
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

func CheckBodyForErrors(body []byte, contentType utils.ContentType) brocade_errors.BrocadeErr {
	var errs struct {
		XMLName xml.Name              `json:"-" xml:"errors"`
		Errors  brocade_errors.Errors `json:"errors" xml:"errors>error"`
	}
	switch contentType {
	case utils.ContentTypeJSON:
		err := contentType.Unmarshal(body, &errs)
		if err != nil || errs.Errors.Errors == nil {
			// fmt.Println("No errors found")
			return nil
		}
		// fmt.Println("JSON errors parsed:", errs)
		return brocade_errors.NewFromErrors(errs.Errors.Errors...)
	case utils.ContentTypeXML:
		err := contentType.Unmarshal(body, &errs.Errors)
		if err != nil || errs.Errors.Errors == nil {
			// fmt.Println("No errors found")
			return nil
		}
		// fmt.Println("XML errors parsed:", errs)
		return brocade_errors.NewFromErrors(errs.Errors.Errors...)
	}
	return brocade_errors.New("unknown content type: '%#v'", contentType)
}

func GetResponseBody(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}
