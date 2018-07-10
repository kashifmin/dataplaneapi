// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// CreateHTTPResponseRuleCreatedCode is the HTTP code returned for type CreateHTTPResponseRuleCreated
const CreateHTTPResponseRuleCreatedCode int = 201

/*CreateHTTPResponseRuleCreated HTTP Response Rule created

swagger:response createHttpResponseRuleCreated
*/
type CreateHTTPResponseRuleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.HTTPResponseRule `json:"body,omitempty"`
}

// NewCreateHTTPResponseRuleCreated creates CreateHTTPResponseRuleCreated with default headers values
func NewCreateHTTPResponseRuleCreated() *CreateHTTPResponseRuleCreated {

	return &CreateHTTPResponseRuleCreated{}
}

// WithPayload adds the payload to the create Http response rule created response
func (o *CreateHTTPResponseRuleCreated) WithPayload(payload *models.HTTPResponseRule) *CreateHTTPResponseRuleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Http response rule created response
func (o *CreateHTTPResponseRuleCreated) SetPayload(payload *models.HTTPResponseRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPResponseRuleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateHTTPResponseRuleBadRequestCode is the HTTP code returned for type CreateHTTPResponseRuleBadRequest
const CreateHTTPResponseRuleBadRequestCode int = 400

/*CreateHTTPResponseRuleBadRequest Bad request

swagger:response createHttpResponseRuleBadRequest
*/
type CreateHTTPResponseRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateHTTPResponseRuleBadRequest creates CreateHTTPResponseRuleBadRequest with default headers values
func NewCreateHTTPResponseRuleBadRequest() *CreateHTTPResponseRuleBadRequest {

	return &CreateHTTPResponseRuleBadRequest{}
}

// WithPayload adds the payload to the create Http response rule bad request response
func (o *CreateHTTPResponseRuleBadRequest) WithPayload(payload *models.Error) *CreateHTTPResponseRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Http response rule bad request response
func (o *CreateHTTPResponseRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPResponseRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateHTTPResponseRuleConflictCode is the HTTP code returned for type CreateHTTPResponseRuleConflict
const CreateHTTPResponseRuleConflictCode int = 409

/*CreateHTTPResponseRuleConflict The specified resource already exists

swagger:response createHttpResponseRuleConflict
*/
type CreateHTTPResponseRuleConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateHTTPResponseRuleConflict creates CreateHTTPResponseRuleConflict with default headers values
func NewCreateHTTPResponseRuleConflict() *CreateHTTPResponseRuleConflict {

	return &CreateHTTPResponseRuleConflict{}
}

// WithPayload adds the payload to the create Http response rule conflict response
func (o *CreateHTTPResponseRuleConflict) WithPayload(payload *models.Error) *CreateHTTPResponseRuleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Http response rule conflict response
func (o *CreateHTTPResponseRuleConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPResponseRuleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateHTTPResponseRuleDefault General Error

swagger:response createHttpResponseRuleDefault
*/
type CreateHTTPResponseRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateHTTPResponseRuleDefault creates CreateHTTPResponseRuleDefault with default headers values
func NewCreateHTTPResponseRuleDefault(code int) *CreateHTTPResponseRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateHTTPResponseRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create HTTP response rule default response
func (o *CreateHTTPResponseRuleDefault) WithStatusCode(code int) *CreateHTTPResponseRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create HTTP response rule default response
func (o *CreateHTTPResponseRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create HTTP response rule default response
func (o *CreateHTTPResponseRuleDefault) WithPayload(payload *models.Error) *CreateHTTPResponseRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create HTTP response rule default response
func (o *CreateHTTPResponseRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateHTTPResponseRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
