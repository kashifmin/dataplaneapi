// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBackendSwitchingRuleParams creates a new GetBackendSwitchingRuleParams object
// no default values defined in spec.
func NewGetBackendSwitchingRuleParams() GetBackendSwitchingRuleParams {

	return GetBackendSwitchingRuleParams{}
}

// GetBackendSwitchingRuleParams contains all the bound params for the get backend switching rule operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBackendSwitchingRule
type GetBackendSwitchingRuleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Frontend name
	  Required: true
	  In: query
	*/
	Frontend string
	/*Switching Rule ID
	  Required: true
	  In: path
	*/
	ID int64
	/*ID of the transaction where we want to add the operation
	  In: query
	*/
	TransactionID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBackendSwitchingRuleParams() beforehand.
func (o *GetBackendSwitchingRuleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrontend, qhkFrontend, _ := qs.GetOK("frontend")
	if err := o.bindFrontend(qFrontend, qhkFrontend, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTransactionID, qhkTransactionID, _ := qs.GetOK("transaction_id")
	if err := o.bindTransactionID(qTransactionID, qhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendSwitchingRuleParams) bindFrontend(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("frontend", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("frontend", "query", raw); err != nil {
		return err
	}

	o.Frontend = raw

	return nil
}

func (o *GetBackendSwitchingRuleParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}

func (o *GetBackendSwitchingRuleParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TransactionID = &raw

	return nil
}
