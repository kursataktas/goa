// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// calc HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design

package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	calcsvc "goa.design/goa/examples/calc/gen/calc"
	goahttp "goa.design/goa/http"
)

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "calc" service "add" endpoint
func (c *Client) BuildAddRequest(v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calcsvc.AddPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "add", "*calcsvc.AddPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "add", u.String(), err)
	}

	return req, nil
}

// DecodeAddResponse returns a decoder for responses returned by the calc add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "add", err)
			}

			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}
