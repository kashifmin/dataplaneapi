// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetBackendOKCode is the HTTP code returned for type GetBackendOK
const GetBackendOKCode int = 200

/*GetBackendOK Successful operation

swagger:response getBackendOK
*/
type GetBackendOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetBackendOKBody `json:"body,omitempty"`
}

// NewGetBackendOK creates GetBackendOK with default headers values
func NewGetBackendOK() *GetBackendOK {

	return &GetBackendOK{}
}

// WithPayload adds the payload to the get backend o k response
func (o *GetBackendOK) WithPayload(payload *models.GetBackendOKBody) *GetBackendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend o k response
func (o *GetBackendOK) SetPayload(payload *models.GetBackendOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackendNotFoundCode is the HTTP code returned for type GetBackendNotFound
const GetBackendNotFoundCode int = 404

/*GetBackendNotFound The specified resource was not found

swagger:response getBackendNotFound
*/
type GetBackendNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBackendNotFound creates GetBackendNotFound with default headers values
func NewGetBackendNotFound() *GetBackendNotFound {

	return &GetBackendNotFound{}
}

// WithPayload adds the payload to the get backend not found response
func (o *GetBackendNotFound) WithPayload(payload *models.Error) *GetBackendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend not found response
func (o *GetBackendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBackendDefault General Error

swagger:response getBackendDefault
*/
type GetBackendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBackendDefault creates GetBackendDefault with default headers values
func NewGetBackendDefault(code int) *GetBackendDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBackendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get backend default response
func (o *GetBackendDefault) WithStatusCode(code int) *GetBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get backend default response
func (o *GetBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get backend default response
func (o *GetBackendDefault) WithPayload(payload *models.Error) *GetBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backend default response
func (o *GetBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
