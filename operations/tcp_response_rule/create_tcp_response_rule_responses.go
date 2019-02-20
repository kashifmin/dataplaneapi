// Code generated by go-swagger; DO NOT EDIT.

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// CreateTCPResponseRuleCreatedCode is the HTTP code returned for type CreateTCPResponseRuleCreated
const CreateTCPResponseRuleCreatedCode int = 201

/*CreateTCPResponseRuleCreated TCP Response Rule created

swagger:response createTcpResponseRuleCreated
*/
type CreateTCPResponseRuleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.TCPResponseRule `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleCreated creates CreateTCPResponseRuleCreated with default headers values
func NewCreateTCPResponseRuleCreated() *CreateTCPResponseRuleCreated {

	return &CreateTCPResponseRuleCreated{}
}

// WithPayload adds the payload to the create Tcp response rule created response
func (o *CreateTCPResponseRuleCreated) WithPayload(payload *models.TCPResponseRule) *CreateTCPResponseRuleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp response rule created response
func (o *CreateTCPResponseRuleCreated) SetPayload(payload *models.TCPResponseRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTCPResponseRuleBadRequestCode is the HTTP code returned for type CreateTCPResponseRuleBadRequest
const CreateTCPResponseRuleBadRequestCode int = 400

/*CreateTCPResponseRuleBadRequest Bad request

swagger:response createTcpResponseRuleBadRequest
*/
type CreateTCPResponseRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleBadRequest creates CreateTCPResponseRuleBadRequest with default headers values
func NewCreateTCPResponseRuleBadRequest() *CreateTCPResponseRuleBadRequest {

	return &CreateTCPResponseRuleBadRequest{}
}

// WithPayload adds the payload to the create Tcp response rule bad request response
func (o *CreateTCPResponseRuleBadRequest) WithPayload(payload *models.Error) *CreateTCPResponseRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp response rule bad request response
func (o *CreateTCPResponseRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTCPResponseRuleConflictCode is the HTTP code returned for type CreateTCPResponseRuleConflict
const CreateTCPResponseRuleConflictCode int = 409

/*CreateTCPResponseRuleConflict The specified resource already exists

swagger:response createTcpResponseRuleConflict
*/
type CreateTCPResponseRuleConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleConflict creates CreateTCPResponseRuleConflict with default headers values
func NewCreateTCPResponseRuleConflict() *CreateTCPResponseRuleConflict {

	return &CreateTCPResponseRuleConflict{}
}

// WithPayload adds the payload to the create Tcp response rule conflict response
func (o *CreateTCPResponseRuleConflict) WithPayload(payload *models.Error) *CreateTCPResponseRuleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Tcp response rule conflict response
func (o *CreateTCPResponseRuleConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTCPResponseRuleDefault General Error

swagger:response createTcpResponseRuleDefault
*/
type CreateTCPResponseRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTCPResponseRuleDefault creates CreateTCPResponseRuleDefault with default headers values
func NewCreateTCPResponseRuleDefault(code int) *CreateTCPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTCPResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) WithStatusCode(code int) *CreateTCPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) WithPayload(payload *models.Error) *CreateTCPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create TCP response rule default response
func (o *CreateTCPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTCPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}