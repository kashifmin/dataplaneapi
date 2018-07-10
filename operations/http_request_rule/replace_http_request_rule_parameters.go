// Code generated by go-swagger; DO NOT EDIT.

package http_request_rule

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

// NewReplaceHTTPRequestRuleParams creates a new ReplaceHTTPRequestRuleParams object
// no default values defined in spec.
func NewReplaceHTTPRequestRuleParams() ReplaceHTTPRequestRuleParams {

	return ReplaceHTTPRequestRuleParams{}
}

// ReplaceHTTPRequestRuleParams contains all the bound params for the replace HTTP request rule operation
// typically these are obtained from a http.Request
//
// swagger:parameters replaceHTTPRequestRule
type ReplaceHTTPRequestRuleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Data *models.HTTPRequestRule
	/*HTTP Request Rule ID
	  Required: true
	  In: path
	*/
	ID int64
	/*Parent name
	  Required: true
	  In: query
	*/
	ParentName string
	/*Parent type
	  Required: true
	  In: query
	*/
	ParentType string
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
// To ensure default values, the struct must have been initialized with NewReplaceHTTPRequestRuleParams() beforehand.
func (o *ReplaceHTTPRequestRuleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.HTTPRequestRule
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

	qParentName, qhkParentName, _ := qs.GetOK("parent_name")
	if err := o.bindParentName(qParentName, qhkParentName, route.Formats); err != nil {
		res = append(res, err)
	}

	qParentType, qhkParentType, _ := qs.GetOK("parent_type")
	if err := o.bindParentType(qParentType, qhkParentType, route.Formats); err != nil {
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

func (o *ReplaceHTTPRequestRuleParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *ReplaceHTTPRequestRuleParams) bindParentName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("parent_name", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("parent_name", "query", raw); err != nil {
		return err
	}

	o.ParentName = raw

	return nil
}

func (o *ReplaceHTTPRequestRuleParams) bindParentType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("parent_type", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("parent_type", "query", raw); err != nil {
		return err
	}

	o.ParentType = raw

	if err := o.validateParentType(formats); err != nil {
		return err
	}

	return nil
}

func (o *ReplaceHTTPRequestRuleParams) validateParentType(formats strfmt.Registry) error {

	if err := validate.Enum("parent_type", "query", o.ParentType, []interface{}{"frontend", "backend"}); err != nil {
		return err
	}

	return nil
}

func (o *ReplaceHTTPRequestRuleParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *ReplaceHTTPRequestRuleParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
