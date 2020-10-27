// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostCohortsOKCode is the HTTP code returned for type PostCohortsOK
const PostCohortsOKCode int = 200

/*PostCohortsOK Updated cohort

swagger:response postCohortsOK
*/
type PostCohortsOK struct {
}

// NewPostCohortsOK creates PostCohortsOK with default headers values
func NewPostCohortsOK() *PostCohortsOK {

	return &PostCohortsOK{}
}

// WriteResponse to the client
func (o *PostCohortsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostCohortsConflictCode is the HTTP code returned for type PostCohortsConflict
const PostCohortsConflictCode int = 409

/*PostCohortsConflict The cohort already exists. Try PUT /node/explore/cohorts to update and existing cohort.

swagger:response postCohortsConflict
*/
type PostCohortsConflict struct {
}

// NewPostCohortsConflict creates PostCohortsConflict with default headers values
func NewPostCohortsConflict() *PostCohortsConflict {

	return &PostCohortsConflict{}
}

// WriteResponse to the client
func (o *PostCohortsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

/*PostCohortsDefault Error response.

swagger:response postCohortsDefault
*/
type PostCohortsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *PostCohortsDefaultBody `json:"body,omitempty"`
}

// NewPostCohortsDefault creates PostCohortsDefault with default headers values
func NewPostCohortsDefault(code int) *PostCohortsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostCohortsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post cohorts default response
func (o *PostCohortsDefault) WithStatusCode(code int) *PostCohortsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post cohorts default response
func (o *PostCohortsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post cohorts default response
func (o *PostCohortsDefault) WithPayload(payload *PostCohortsDefaultBody) *PostCohortsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post cohorts default response
func (o *PostCohortsDefault) SetPayload(payload *PostCohortsDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCohortsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
