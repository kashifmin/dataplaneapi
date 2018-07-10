// Code generated by go-swagger; DO NOT EDIT.

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFiltersHandlerFunc turns a function with the right signature into a get filters handler
type GetFiltersHandlerFunc func(GetFiltersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFiltersHandlerFunc) Handle(params GetFiltersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFiltersHandler interface for that can handle valid get filters params
type GetFiltersHandler interface {
	Handle(GetFiltersParams, interface{}) middleware.Responder
}

// NewGetFilters creates a new http.Handler for the get filters operation
func NewGetFilters(ctx *middleware.Context, handler GetFiltersHandler) *GetFilters {
	return &GetFilters{Context: ctx, Handler: handler}
}

/*GetFilters swagger:route GET /services/haproxy/configuration/filters Filter getFilters

Return an array of all Filters

Returns all Filters that are configured in specified parent.

*/
type GetFilters struct {
	Context *middleware.Context
	Handler GetFiltersHandler
}

func (o *GetFilters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFiltersParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
