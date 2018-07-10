// Code generated by go-swagger; DO NOT EDIT.

package listener

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// CreateListenerCreatedCode is the HTTP code returned for type CreateListenerCreated
const CreateListenerCreatedCode int = 201

/*CreateListenerCreated Listener created

swagger:response createListenerCreated
*/
type CreateListenerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Listener `json:"body,omitempty"`
}

// NewCreateListenerCreated creates CreateListenerCreated with default headers values
func NewCreateListenerCreated() *CreateListenerCreated {

	return &CreateListenerCreated{}
}

// WithPayload adds the payload to the create listener created response
func (o *CreateListenerCreated) WithPayload(payload *models.Listener) *CreateListenerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create listener created response
func (o *CreateListenerCreated) SetPayload(payload *models.Listener) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateListenerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateListenerBadRequestCode is the HTTP code returned for type CreateListenerBadRequest
const CreateListenerBadRequestCode int = 400

/*CreateListenerBadRequest Bad request

swagger:response createListenerBadRequest
*/
type CreateListenerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateListenerBadRequest creates CreateListenerBadRequest with default headers values
func NewCreateListenerBadRequest() *CreateListenerBadRequest {

	return &CreateListenerBadRequest{}
}

// WithPayload adds the payload to the create listener bad request response
func (o *CreateListenerBadRequest) WithPayload(payload *models.Error) *CreateListenerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create listener bad request response
func (o *CreateListenerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateListenerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateListenerConflictCode is the HTTP code returned for type CreateListenerConflict
const CreateListenerConflictCode int = 409

/*CreateListenerConflict The specified resource already exists

swagger:response createListenerConflict
*/
type CreateListenerConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateListenerConflict creates CreateListenerConflict with default headers values
func NewCreateListenerConflict() *CreateListenerConflict {

	return &CreateListenerConflict{}
}

// WithPayload adds the payload to the create listener conflict response
func (o *CreateListenerConflict) WithPayload(payload *models.Error) *CreateListenerConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create listener conflict response
func (o *CreateListenerConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateListenerConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateListenerDefault General Error

swagger:response createListenerDefault
*/
type CreateListenerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateListenerDefault creates CreateListenerDefault with default headers values
func NewCreateListenerDefault(code int) *CreateListenerDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateListenerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create listener default response
func (o *CreateListenerDefault) WithStatusCode(code int) *CreateListenerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create listener default response
func (o *CreateListenerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create listener default response
func (o *CreateListenerDefault) WithPayload(payload *models.Error) *CreateListenerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create listener default response
func (o *CreateListenerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateListenerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
