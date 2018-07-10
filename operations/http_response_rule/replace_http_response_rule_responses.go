// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceHTTPResponseRuleOKCode is the HTTP code returned for type ReplaceHTTPResponseRuleOK
const ReplaceHTTPResponseRuleOKCode int = 200

/*ReplaceHTTPResponseRuleOK HTTP Response Rule replaced

swagger:response replaceHttpResponseRuleOK
*/
type ReplaceHTTPResponseRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.HTTPResponseRule `json:"body,omitempty"`
}

// NewReplaceHTTPResponseRuleOK creates ReplaceHTTPResponseRuleOK with default headers values
func NewReplaceHTTPResponseRuleOK() *ReplaceHTTPResponseRuleOK {

	return &ReplaceHTTPResponseRuleOK{}
}

// WithPayload adds the payload to the replace Http response rule o k response
func (o *ReplaceHTTPResponseRuleOK) WithPayload(payload *models.HTTPResponseRule) *ReplaceHTTPResponseRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http response rule o k response
func (o *ReplaceHTTPResponseRuleOK) SetPayload(payload *models.HTTPResponseRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPResponseRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPResponseRuleBadRequestCode is the HTTP code returned for type ReplaceHTTPResponseRuleBadRequest
const ReplaceHTTPResponseRuleBadRequestCode int = 400

/*ReplaceHTTPResponseRuleBadRequest Bad request

swagger:response replaceHttpResponseRuleBadRequest
*/
type ReplaceHTTPResponseRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPResponseRuleBadRequest creates ReplaceHTTPResponseRuleBadRequest with default headers values
func NewReplaceHTTPResponseRuleBadRequest() *ReplaceHTTPResponseRuleBadRequest {

	return &ReplaceHTTPResponseRuleBadRequest{}
}

// WithPayload adds the payload to the replace Http response rule bad request response
func (o *ReplaceHTTPResponseRuleBadRequest) WithPayload(payload *models.Error) *ReplaceHTTPResponseRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http response rule bad request response
func (o *ReplaceHTTPResponseRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPResponseRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceHTTPResponseRuleNotFoundCode is the HTTP code returned for type ReplaceHTTPResponseRuleNotFound
const ReplaceHTTPResponseRuleNotFoundCode int = 404

/*ReplaceHTTPResponseRuleNotFound The specified resource was not found

swagger:response replaceHttpResponseRuleNotFound
*/
type ReplaceHTTPResponseRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPResponseRuleNotFound creates ReplaceHTTPResponseRuleNotFound with default headers values
func NewReplaceHTTPResponseRuleNotFound() *ReplaceHTTPResponseRuleNotFound {

	return &ReplaceHTTPResponseRuleNotFound{}
}

// WithPayload adds the payload to the replace Http response rule not found response
func (o *ReplaceHTTPResponseRuleNotFound) WithPayload(payload *models.Error) *ReplaceHTTPResponseRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Http response rule not found response
func (o *ReplaceHTTPResponseRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPResponseRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceHTTPResponseRuleDefault General Error

swagger:response replaceHttpResponseRuleDefault
*/
type ReplaceHTTPResponseRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceHTTPResponseRuleDefault creates ReplaceHTTPResponseRuleDefault with default headers values
func NewReplaceHTTPResponseRuleDefault(code int) *ReplaceHTTPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceHTTPResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace HTTP response rule default response
func (o *ReplaceHTTPResponseRuleDefault) WithStatusCode(code int) *ReplaceHTTPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace HTTP response rule default response
func (o *ReplaceHTTPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace HTTP response rule default response
func (o *ReplaceHTTPResponseRuleDefault) WithPayload(payload *models.Error) *ReplaceHTTPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace HTTP response rule default response
func (o *ReplaceHTTPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceHTTPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
