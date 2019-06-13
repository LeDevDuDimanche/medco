// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/lca1/medco-connector/restapi/models"
)

// QueryResultHandlerFunc turns a function with the right signature into a query result handler
type QueryResultHandlerFunc func(QueryResultParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn QueryResultHandlerFunc) Handle(params QueryResultParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// QueryResultHandler interface for that can handle valid query result params
type QueryResultHandler interface {
	Handle(QueryResultParams, *models.User) middleware.Responder
}

// NewQueryResult creates a new http.Handler for the query result operation
func NewQueryResult(ctx *middleware.Context, handler QueryResultHandler) *QueryResult {
	return &QueryResult{Context: ctx, Handler: handler}
}

/*QueryResult swagger:route POST /picsure2/query/{queryId}/result picsure2 queryResult

Get result of query.

*/
type QueryResult struct {
	Context *middleware.Context
	Handler QueryResultHandler
}

func (o *QueryResult) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewQueryResultParams()

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

// QueryResultBody query result body
// swagger:model QueryResultBody
type QueryResultBody struct {

	// resource credentials
	ResourceCredentials *models.ResourceCredentials `json:"resourceCredentials,omitempty"`
}

// Validate validates this query result body
func (o *QueryResultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResourceCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryResultBody) validateResourceCredentials(formats strfmt.Registry) error {

	if swag.IsZero(o.ResourceCredentials) { // not required
		return nil
	}

	if o.ResourceCredentials != nil {
		if err := o.ResourceCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "resourceCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryResultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryResultBody) UnmarshalBinary(b []byte) error {
	var res QueryResultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// QueryResultDefaultBody query result default body
// swagger:model QueryResultDefaultBody
type QueryResultDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this query result default body
func (o *QueryResultDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *QueryResultDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryResultDefaultBody) UnmarshalBinary(b []byte) error {
	var res QueryResultDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
