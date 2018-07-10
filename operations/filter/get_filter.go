// Code generated by go-swagger; DO NOT EDIT.

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFilterHandlerFunc turns a function with the right signature into a get filter handler
type GetFilterHandlerFunc func(GetFilterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFilterHandlerFunc) Handle(params GetFilterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFilterHandler interface for that can handle valid get filter params
type GetFilterHandler interface {
	Handle(GetFilterParams, interface{}) middleware.Responder
}

// NewGetFilter creates a new http.Handler for the get filter operation
func NewGetFilter(ctx *middleware.Context, handler GetFilterHandler) *GetFilter {
	return &GetFilter{Context: ctx, Handler: handler}
}

/*GetFilter swagger:route GET /services/haproxy/configuration/filters/{id} Filter getFilter

Return one Filter

Returns one Filter configuration by it's ID in the specified parent.

*/
type GetFilter struct {
	Context *middleware.Context
	Handler GetFilterHandler
}

func (o *GetFilter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFilterParams()

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
