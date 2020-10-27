// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ldsec/medco/connector/restapi/models"
)

// PutCohortsHandlerFunc turns a function with the right signature into a put cohorts handler
type PutCohortsHandlerFunc func(PutCohortsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PutCohortsHandlerFunc) Handle(params PutCohortsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PutCohortsHandler interface for that can handle valid put cohorts params
type PutCohortsHandler interface {
	Handle(PutCohortsParams, *models.User) middleware.Responder
}

// NewPutCohorts creates a new http.Handler for the put cohorts operation
func NewPutCohorts(ctx *middleware.Context, handler PutCohortsHandler) *PutCohorts {
	return &PutCohorts{Context: ctx, Handler: handler}
}

/*PutCohorts swagger:route PUT /node/explore/cohorts/{name} medco-node putCohorts

Update a cohort

*/
type PutCohorts struct {
	Context *middleware.Context
	Handler PutCohortsHandler
}

func (o *PutCohorts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutCohortsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutCohortsBody put cohorts body
//
// swagger:model PutCohortsBody
type PutCohortsBody struct {

	// creation date
	CreationDate string `json:"creationDate,omitempty"`

	// patient set ID
	PatientSetID int64 `json:"patientSetID,omitempty"`

	// update date
	UpdateDate string `json:"updateDate,omitempty"`
}

// Validate validates this put cohorts body
func (o *PutCohortsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCohortsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCohortsBody) UnmarshalBinary(b []byte) error {
	var res PutCohortsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutCohortsDefaultBody put cohorts default body
//
// swagger:model PutCohortsDefaultBody
type PutCohortsDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put cohorts default body
func (o *PutCohortsDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCohortsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCohortsDefaultBody) UnmarshalBinary(b []byte) error {
	var res PutCohortsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
