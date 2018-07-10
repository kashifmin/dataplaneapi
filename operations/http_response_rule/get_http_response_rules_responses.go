// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetHTTPResponseRulesOKCode is the HTTP code returned for type GetHTTPResponseRulesOK
const GetHTTPResponseRulesOKCode int = 200

/*GetHTTPResponseRulesOK Successful operation

swagger:response getHttpResponseRulesOK
*/
type GetHTTPResponseRulesOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetHTTPResponseRulesOKBody `json:"body,omitempty"`
}

// NewGetHTTPResponseRulesOK creates GetHTTPResponseRulesOK with default headers values
func NewGetHTTPResponseRulesOK() *GetHTTPResponseRulesOK {

	return &GetHTTPResponseRulesOK{}
}

// WithPayload adds the payload to the get Http response rules o k response
func (o *GetHTTPResponseRulesOK) WithPayload(payload *models.GetHTTPResponseRulesOKBody) *GetHTTPResponseRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Http response rules o k response
func (o *GetHTTPResponseRulesOK) SetPayload(payload *models.GetHTTPResponseRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPResponseRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetHTTPResponseRulesDefault General Error

swagger:response getHttpResponseRulesDefault
*/
type GetHTTPResponseRulesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHTTPResponseRulesDefault creates GetHTTPResponseRulesDefault with default headers values
func NewGetHTTPResponseRulesDefault(code int) *GetHTTPResponseRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHTTPResponseRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) WithStatusCode(code int) *GetHTTPResponseRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) WithPayload(payload *models.Error) *GetHTTPResponseRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get HTTP response rules default response
func (o *GetHTTPResponseRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHTTPResponseRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
