// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// NewReplaceServerSwitchingRuleParams creates a new ReplaceServerSwitchingRuleParams object
// no default values defined in spec.
func NewReplaceServerSwitchingRuleParams() ReplaceServerSwitchingRuleParams {

	return ReplaceServerSwitchingRuleParams{}
}

// ReplaceServerSwitchingRuleParams contains all the bound params for the replace server switching rule operation
// typically these are obtained from a http.Request
//
// swagger:parameters replaceServerSwitchingRule
type ReplaceServerSwitchingRuleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Backend name
	  Required: true
	  In: query
	*/
	Backend string
	/*
	  Required: true
	  In: body
	*/
	Data *models.ServerSwitchingRule
	/*Switching Rule ID
	  Required: true
	  In: path
	*/
	ID int64
	/*ID of the transaction where we want to add the operation
	  In: query
	*/
	TransactionID *string
	/*Version used for checking configuration version
	  In: query
	*/
	Version *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReplaceServerSwitchingRuleParams() beforehand.
func (o *ReplaceServerSwitchingRuleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBackend, qhkBackend, _ := qs.GetOK("backend")
	if err := o.bindBackend(qBackend, qhkBackend, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ServerSwitchingRule
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body"))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body"))
	}
	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTransactionID, qhkTransactionID, _ := qs.GetOK("transaction_id")
	if err := o.bindTransactionID(qTransactionID, qhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplaceServerSwitchingRuleParams) bindBackend(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("backend", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("backend", "query", raw); err != nil {
		return err
	}

	o.Backend = raw

	return nil
}

func (o *ReplaceServerSwitchingRuleParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *ReplaceServerSwitchingRuleParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *ReplaceServerSwitchingRuleParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("version", "query", "int64", raw)
	}
	o.Version = &value

	return nil
}
