// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteSiteHandlerFunc turns a function with the right signature into a delete site handler
type DeleteSiteHandlerFunc func(DeleteSiteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSiteHandlerFunc) Handle(params DeleteSiteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteSiteHandler interface for that can handle valid delete site params
type DeleteSiteHandler interface {
	Handle(DeleteSiteParams, interface{}) middleware.Responder
}

// NewDeleteSite creates a new http.Handler for the delete site operation
func NewDeleteSite(ctx *middleware.Context, handler DeleteSiteHandler) *DeleteSite {
	return &DeleteSite{Context: ctx, Handler: handler}
}

/*DeleteSite swagger:route DELETE /services/haproxy/sites/{name} Sites deleteSite

Delete a site

Deletes a site from the configuration by it's name.

*/
type DeleteSite struct {
	Context *middleware.Context
	Handler DeleteSiteHandler
}

func (o *DeleteSite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteSiteParams()

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
