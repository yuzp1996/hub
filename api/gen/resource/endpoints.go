// Code generated by goa v3.16.0, DO NOT EDIT.
//
// resource endpoints
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "resource" service endpoints.
type Endpoints struct {
	Query                    goa.Endpoint
	List                     goa.Endpoint
	VersionsByID             goa.Endpoint
	ByCatalogKindNameVersion goa.Endpoint
	ByVersionID              goa.Endpoint
	ByCatalogKindName        goa.Endpoint
	ByID                     goa.Endpoint
}

// NewEndpoints wraps the methods of the "resource" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Query:                    NewQueryEndpoint(s),
		List:                     NewListEndpoint(s),
		VersionsByID:             NewVersionsByIDEndpoint(s),
		ByCatalogKindNameVersion: NewByCatalogKindNameVersionEndpoint(s),
		ByVersionID:              NewByVersionIDEndpoint(s),
		ByCatalogKindName:        NewByCatalogKindNameEndpoint(s),
		ByID:                     NewByIDEndpoint(s),
	}
}

// Use applies the given middleware to all the "resource" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Query = m(e.Query)
	e.List = m(e.List)
	e.VersionsByID = m(e.VersionsByID)
	e.ByCatalogKindNameVersion = m(e.ByCatalogKindNameVersion)
	e.ByVersionID = m(e.ByVersionID)
	e.ByCatalogKindName = m(e.ByCatalogKindName)
	e.ByID = m(e.ByID)
}

// NewQueryEndpoint returns an endpoint function that calls the method "Query"
// of service "resource".
func NewQueryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*QueryPayload)
		return s.Query(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "List" of
// service "resource".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResources(res, "default")
		return vres, nil
	}
}

// NewVersionsByIDEndpoint returns an endpoint function that calls the method
// "VersionsByID" of service "resource".
func NewVersionsByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*VersionsByIDPayload)
		return s.VersionsByID(ctx, p)
	}
}

// NewByCatalogKindNameVersionEndpoint returns an endpoint function that calls
// the method "ByCatalogKindNameVersion" of service "resource".
func NewByCatalogKindNameVersionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByCatalogKindNameVersionPayload)
		return s.ByCatalogKindNameVersion(ctx, p)
	}
}

// NewByVersionIDEndpoint returns an endpoint function that calls the method
// "ByVersionId" of service "resource".
func NewByVersionIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByVersionIDPayload)
		return s.ByVersionID(ctx, p)
	}
}

// NewByCatalogKindNameEndpoint returns an endpoint function that calls the
// method "ByCatalogKindName" of service "resource".
func NewByCatalogKindNameEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByCatalogKindNamePayload)
		return s.ByCatalogKindName(ctx, p)
	}
}

// NewByIDEndpoint returns an endpoint function that calls the method "ById" of
// service "resource".
func NewByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ByIDPayload)
		return s.ByID(ctx, p)
	}
}
