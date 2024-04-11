// Code generated by goa v3.16.0, DO NOT EDIT.
//
// rating HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	rating "github.com/tektoncd/hub/api/gen/rating"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "rating" service "Get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint
	)
	{
		p, ok := v.(*rating.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("rating", "Get", "*rating.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetRatingPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("rating", "Get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the rating Get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*rating.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("rating", "Get", "*rating.GetPayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the rating Get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeGetResponse may return the following errors:
//   - "not-found" (type *goa.ServiceError): http.StatusNotFound
//   - "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//   - "invalid-token" (type *goa.ServiceError): http.StatusUnauthorized
//   - "invalid-scopes" (type *goa.ServiceError): http.StatusForbidden
//   - error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Get", err)
			}
			res := NewGetResultOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Get", err)
			}
			err = ValidateGetNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Get", err)
			}
			return nil, NewGetNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body GetInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Get", err)
			}
			err = ValidateGetInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Get", err)
			}
			return nil, NewGetInternalError(&body)
		case http.StatusUnauthorized:
			var (
				body GetInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Get", err)
			}
			err = ValidateGetInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Get", err)
			}
			return nil, NewGetInvalidToken(&body)
		case http.StatusForbidden:
			var (
				body GetInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Get", err)
			}
			err = ValidateGetInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Get", err)
			}
			return nil, NewGetInvalidScopes(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("rating", "Get", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "rating" service "Update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id uint
	)
	{
		p, ok := v.(*rating.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("rating", "Update", "*rating.UpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateRatingPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("rating", "Update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the rating
// Update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*rating.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("rating", "Update", "*rating.UpdatePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("rating", "Update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the rating
// Update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "not-found" (type *goa.ServiceError): http.StatusNotFound
//   - "internal-error" (type *goa.ServiceError): http.StatusInternalServerError
//   - "invalid-token" (type *goa.ServiceError): http.StatusUnauthorized
//   - "invalid-scopes" (type *goa.ServiceError): http.StatusForbidden
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
			return nil, nil
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Update", err)
			}
			return nil, NewUpdateNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body UpdateInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Update", err)
			}
			err = ValidateUpdateInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Update", err)
			}
			return nil, NewUpdateInternalError(&body)
		case http.StatusUnauthorized:
			var (
				body UpdateInvalidTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Update", err)
			}
			err = ValidateUpdateInvalidTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Update", err)
			}
			return nil, NewUpdateInvalidToken(&body)
		case http.StatusForbidden:
			var (
				body UpdateInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("rating", "Update", err)
			}
			err = ValidateUpdateInvalidScopesResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("rating", "Update", err)
			}
			return nil, NewUpdateInvalidScopes(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("rating", "Update", resp.StatusCode, string(body))
		}
	}
}
