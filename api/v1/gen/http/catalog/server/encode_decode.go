// Code generated by goa v3.14.5, DO NOT EDIT.
//
// catalog HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/v1/design

package server

import (
	"context"
	"errors"
	"net/http"

	catalog "github.com/tektoncd/hub/api/v1/gen/catalog"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the catalog
// List endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*catalog.ListResult)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeListError returns an encoder for errors returned by the List catalog
// endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "internal-error":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalCatalogCatalogToCatalogResponseBody builds a value of type
// *CatalogResponseBody from a value of type *catalog.Catalog.
func marshalCatalogCatalogToCatalogResponseBody(v *catalog.Catalog) *CatalogResponseBody {
	if v == nil {
		return nil
	}
	res := &CatalogResponseBody{
		ID:       v.ID,
		Name:     v.Name,
		Type:     v.Type,
		URL:      v.URL,
		Provider: v.Provider,
	}

	return res
}
