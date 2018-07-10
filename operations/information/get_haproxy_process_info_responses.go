// Code generated by go-swagger; DO NOT EDIT.

package information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetHaproxyProcessInfoOKCode is the HTTP code returned for type GetHaproxyProcessInfoOK
const GetHaproxyProcessInfoOKCode int = 200

/*GetHaproxyProcessInfoOK Success

swagger:response getHaproxyProcessInfoOK
*/
type GetHaproxyProcessInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProcessInfo `json:"body,omitempty"`
}

// NewGetHaproxyProcessInfoOK creates GetHaproxyProcessInfoOK with default headers values
func NewGetHaproxyProcessInfoOK() *GetHaproxyProcessInfoOK {

	return &GetHaproxyProcessInfoOK{}
}

// WithPayload adds the payload to the get haproxy process info o k response
func (o *GetHaproxyProcessInfoOK) WithPayload(payload *models.ProcessInfo) *GetHaproxyProcessInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get haproxy process info o k response
func (o *GetHaproxyProcessInfoOK) SetPayload(payload *models.ProcessInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHaproxyProcessInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetHaproxyProcessInfoDefault General Error

swagger:response getHaproxyProcessInfoDefault
*/
type GetHaproxyProcessInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHaproxyProcessInfoDefault creates GetHaproxyProcessInfoDefault with default headers values
func NewGetHaproxyProcessInfoDefault(code int) *GetHaproxyProcessInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHaproxyProcessInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get haproxy process info default response
func (o *GetHaproxyProcessInfoDefault) WithStatusCode(code int) *GetHaproxyProcessInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get haproxy process info default response
func (o *GetHaproxyProcessInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get haproxy process info default response
func (o *GetHaproxyProcessInfoDefault) WithPayload(payload *models.Error) *GetHaproxyProcessInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get haproxy process info default response
func (o *GetHaproxyProcessInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHaproxyProcessInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
