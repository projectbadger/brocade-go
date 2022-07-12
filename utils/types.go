package utils

import (
	brocade_errors "brocade/rest/errors"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type ContentType string

func (ct *ContentType) SetRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", ct.String())
	req.Header.Set("Accept", ct.String())
}

func (ct *ContentType) Valid() bool {
	if ct == nil {
		return false
	}
	str := string(*ct)
	switch str {
	case "application/yang-data+json", "application/yang-data+xml":
		return true
	}
	return false
}

func (ct *ContentType) String() string {
	if ct != nil && ct.Valid() {
		return string(*ct)
	}
	return ""
}

func (ct *ContentType) Unmarshal(data []byte, v interface{}) brocade_errors.BrocadeErr {
	if ct == nil {
		return ErrContentTypeNil
	}
	switch *ct {
	case ContentTypeJSON:
		return brocade_errors.NewFromErr(json.Unmarshal(data, v))
	case ContentTypeXML:
		return brocade_errors.NewFromErr(xml.Unmarshal(data, v))
	}
	return ErrContentTypeInvalid
}

func (ct *ContentType) UnmarshalResponse(data []byte, v interface{}) brocade_errors.BrocadeErr {
	if ct == nil {
		return ErrContentTypeNil
	}
	// fmt.Printf("*ct: '%#v' xml: '%#v', json: '%#v'\n", *ct, ContentTypeXML, ContentTypeJSON)
	switch *ct {
	case ContentTypeJSON:
		// err := json.Unmarshal(data, v)
		// fmt.Printf("resp before parsing: '%v'\nerror: '%#v'\n", v, err)
		// if err != nil && v != nil {
		var resp struct {
			Response map[string]*json.RawMessage `json:"Response,omitempty"`
			Errors   brocade_errors.Errors       `json:"errors,omitempty"`
		}
		err := json.Unmarshal(data, &resp.Errors)
		if err == nil && resp.Errors.Errors != nil {
			return brocade_errors.NewFromErrors(resp.Errors.Errors...)
		}
		_ = json.Unmarshal(data, &resp)
		if resp.Errors.Errors != nil {
			return brocade_errors.NewFromErrors(resp.Errors.Errors...)
		}
		// fmt.Printf("resp after unmarshal: '%v'\n", resp)
		if resp.Response == nil {
			_ = json.Unmarshal(data, &resp.Response)
		}
		if resp.Response != nil {
			for _, val := range resp.Response {
				err = json.Unmarshal(*val, v)
				if err != nil {
					return brocade_errors.NewFromErr(err)
				}
				return nil
			}
			return ErrNoParsableContentFound
		}
		// fmt.Printf("Data: '%#v' struct: '%#v'\n", string(data), resp)
		// return brocade_errors.New("no parsable content found")
		// }
		err = json.Unmarshal(data, v)
		// fmt.Printf("resp before parsing: '%v'\nerror: '%#v'\n", v, err)
		return brocade_errors.NewFromErr(err)
	case ContentTypeXML:
		err := xml.Unmarshal(data, v)
		if err != nil {
			var errs brocade_errors.Errors
			var resp struct {
				Response struct {
					InnerXML string `xml:",innerxml"`
				}
				Errors brocade_errors.Errors `xml:"errors"`
			}
			if err = xml.Unmarshal(data, &errs); err == nil && errs.Errors != nil {
				return brocade_errors.NewFromErrors(errs.Errors...)
			}
			err = xml.Unmarshal(data, &resp)
			// if err != nil {
			// 	fmt.Println("Error: ", err)
			// }
			// fmt.Printf("resp: '%#v'\n", resp)
			if resp.Response.InnerXML != "" {
				data = []byte(resp.Response.InnerXML)
				err = xml.Unmarshal(data, v)
				if err != nil {
					return brocade_errors.NewFromErr(err)
				}
				return nil
			}
			if err = xml.Unmarshal(data, &resp.Response); err == nil {
				data = []byte(resp.Response.InnerXML)
				err = xml.Unmarshal(data, v)
				if err != nil {
					return brocade_errors.NewFromErr(err)
				}
				return nil
			}
			return ErrNoParsableContentFound
		}
		return ErrNoParsableContentFound
		// return nil
	}
	return ErrContentTypeInvalid
}

func (ct *ContentType) Marshal(v interface{}) ([]byte, error) {
	if ct == nil {
		return nil, ErrContentTypeNil
	}
	switch *ct {
	case ContentTypeJSON:
		return json.Marshal(v)
	case ContentTypeXML:
		return xml.Marshal(v)
	}
	return nil, ErrContentTypeInvalid
}

type RequestClient interface {
	Do(*http.Request) (*http.Response, error)
}
