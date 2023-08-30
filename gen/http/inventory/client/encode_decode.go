// Code generated by goa v3.12.4, DO NOT EDIT.
//
// inventory HTTP client encoders and decoders
//
// Command:
// $ goa gen game-service/design

package client

import (
	"bytes"
	"context"
	inventory "game-service/gen/inventory"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "inventory" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*inventory.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("inventory", "show", "*inventory.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowInventoryPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("inventory", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeShowResponse returns a decoder for responses returned by the inventory
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeShowResponse may return the following errors:
//   - "not_found" (type *inventory.NotFound): http.StatusNotFound
//   - error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("inventory", "show", err)
			}
			return body, nil
		case http.StatusNotFound:
			var (
				body ShowNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("inventory", "show", err)
			}
			err = ValidateShowNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("inventory", "show", err)
			}
			return nil, NewShowNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("inventory", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "inventory" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddInventoryPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("inventory", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the inventory add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*inventory.AddPayload)
		if !ok {
			return goahttp.ErrInvalidType("inventory", "add", "*inventory.AddPayload", v)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("inventory", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the inventory
// add endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("inventory", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveRequest instantiates a HTTP request object with method and path
// set to call the "inventory" service "remove" endpoint
func (c *Client) BuildRemoveRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveInventoryPath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("inventory", "remove", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRemoveRequest returns an encoder for requests sent to the inventory
// remove server.
func EncodeRemoveRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*inventory.RemovePayload)
		if !ok {
			return goahttp.ErrInvalidType("inventory", "remove", "*inventory.RemovePayload", v)
		}
		body := NewRemoveRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("inventory", "remove", err)
		}
		return nil
	}
}

// DecodeRemoveResponse returns a decoder for responses returned by the
// inventory remove endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRemoveResponse may return the following errors:
//   - "not_found" (type *inventory.NotFound): http.StatusNotFound
//   - error: internal error
func DecodeRemoveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusNotFound:
			var (
				body RemoveNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("inventory", "remove", err)
			}
			err = ValidateRemoveNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("inventory", "remove", err)
			}
			return nil, NewRemoveNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("inventory", "remove", resp.StatusCode, string(body))
		}
	}
}