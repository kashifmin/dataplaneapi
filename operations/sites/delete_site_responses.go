// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// DeleteSiteNoContentCode is the HTTP code returned for type DeleteSiteNoContent
const DeleteSiteNoContentCode int = 204

/*DeleteSiteNoContent Site deleted

swagger:response deleteSiteNoContent
*/
type DeleteSiteNoContent struct {
}

// NewDeleteSiteNoContent creates DeleteSiteNoContent with default headers values
func NewDeleteSiteNoContent() *DeleteSiteNoContent {

	return &DeleteSiteNoContent{}
}

// WriteResponse to the client
func (o *DeleteSiteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteSiteNotFoundCode is the HTTP code returned for type DeleteSiteNotFound
const DeleteSiteNotFoundCode int = 404

/*DeleteSiteNotFound The specified resource was not found

swagger:response deleteSiteNotFound
*/
type DeleteSiteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSiteNotFound creates DeleteSiteNotFound with default headers values
func NewDeleteSiteNotFound() *DeleteSiteNotFound {

	return &DeleteSiteNotFound{}
}

// WithPayload adds the payload to the delete site not found response
func (o *DeleteSiteNotFound) WithPayload(payload *models.Error) *DeleteSiteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete site not found response
func (o *DeleteSiteNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSiteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSiteDefault General Error

swagger:response deleteSiteDefault
*/
type DeleteSiteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSiteDefault creates DeleteSiteDefault with default headers values
func NewDeleteSiteDefault(code int) *DeleteSiteDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSiteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete site default response
func (o *DeleteSiteDefault) WithStatusCode(code int) *DeleteSiteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete site default response
func (o *DeleteSiteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete site default response
func (o *DeleteSiteDefault) WithPayload(payload *models.Error) *DeleteSiteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete site default response
func (o *DeleteSiteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSiteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
