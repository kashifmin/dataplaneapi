// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSitesHandlerFunc turns a function with the right signature into a get sites handler
type GetSitesHandlerFunc func(GetSitesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSitesHandlerFunc) Handle(params GetSitesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSitesHandler interface for that can handle valid get sites params
type GetSitesHandler interface {
	Handle(GetSitesParams, interface{}) middleware.Responder
}

// NewGetSites creates a new http.Handler for the get sites operation
func NewGetSites(ctx *middleware.Context, handler GetSitesHandler) *GetSites {
	return &GetSites{Context: ctx, Handler: handler}
}

/*GetSites swagger:route GET /services/haproxy/sites Sites getSites

Return an array of sites

Returns an array of all configured sites.

*/
type GetSites struct {
	Context *middleware.Context
	Handler GetSitesHandler
}

func (o *GetSites) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSitesParams()

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
