// Code generated by go-swagger; DO NOT EDIT.

package share

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/openziti/zrok/rest_model_zrok"
)

// AccessHandlerFunc turns a function with the right signature into a access handler
type AccessHandlerFunc func(AccessParams, *rest_model_zrok.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn AccessHandlerFunc) Handle(params AccessParams, principal *rest_model_zrok.Principal) middleware.Responder {
	return fn(params, principal)
}

// AccessHandler interface for that can handle valid access params
type AccessHandler interface {
	Handle(AccessParams, *rest_model_zrok.Principal) middleware.Responder
}

// NewAccess creates a new http.Handler for the access operation
func NewAccess(ctx *middleware.Context, handler AccessHandler) *Access {
	return &Access{Context: ctx, Handler: handler}
}

/*
	Access swagger:route POST /access share access

Access access API
*/
type Access struct {
	Context *middleware.Context
	Handler AccessHandler
}

func (o *Access) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAccessParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *rest_model_zrok.Principal
	if uprinc != nil {
		principal = uprinc.(*rest_model_zrok.Principal) // this is really a rest_model_zrok.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}