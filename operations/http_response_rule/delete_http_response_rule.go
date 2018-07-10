// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteHTTPResponseRuleHandlerFunc turns a function with the right signature into a delete HTTP response rule handler
type DeleteHTTPResponseRuleHandlerFunc func(DeleteHTTPResponseRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHTTPResponseRuleHandlerFunc) Handle(params DeleteHTTPResponseRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteHTTPResponseRuleHandler interface for that can handle valid delete HTTP response rule params
type DeleteHTTPResponseRuleHandler interface {
	Handle(DeleteHTTPResponseRuleParams, interface{}) middleware.Responder
}

// NewDeleteHTTPResponseRule creates a new http.Handler for the delete HTTP response rule operation
func NewDeleteHTTPResponseRule(ctx *middleware.Context, handler DeleteHTTPResponseRuleHandler) *DeleteHTTPResponseRule {
	return &DeleteHTTPResponseRule{Context: ctx, Handler: handler}
}

/*DeleteHTTPResponseRule swagger:route DELETE /services/haproxy/configuration/http_response_rules/{id} HTTPResponseRule deleteHttpResponseRule

Delete a HTTP Response Rule

Deletes a HTTP Response Rule configuration by it's ID from the specified parent.

*/
type DeleteHTTPResponseRule struct {
	Context *middleware.Context
	Handler DeleteHTTPResponseRuleHandler
}

func (o *DeleteHTTPResponseRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteHTTPResponseRuleParams()

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
