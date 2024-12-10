// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListOrgMembersParams creates a new ListOrgMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOrgMembersParams() *ListOrgMembersParams {
	return &ListOrgMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOrgMembersParamsWithTimeout creates a new ListOrgMembersParams object
// with the ability to set a timeout on a request.
func NewListOrgMembersParamsWithTimeout(timeout time.Duration) *ListOrgMembersParams {
	return &ListOrgMembersParams{
		timeout: timeout,
	}
}

// NewListOrgMembersParamsWithContext creates a new ListOrgMembersParams object
// with the ability to set a context for a request.
func NewListOrgMembersParamsWithContext(ctx context.Context) *ListOrgMembersParams {
	return &ListOrgMembersParams{
		Context: ctx,
	}
}

// NewListOrgMembersParamsWithHTTPClient creates a new ListOrgMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOrgMembersParamsWithHTTPClient(client *http.Client) *ListOrgMembersParams {
	return &ListOrgMembersParams{
		HTTPClient: client,
	}
}

/*
ListOrgMembersParams contains all the parameters to send to the API endpoint

	for the list org members operation.

	Typically these are written to a http.Request.
*/
type ListOrgMembersParams struct {

	// OrganizationToken.
	OrganizationToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list org members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrgMembersParams) WithDefaults() *ListOrgMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list org members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOrgMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list org members params
func (o *ListOrgMembersParams) WithTimeout(timeout time.Duration) *ListOrgMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list org members params
func (o *ListOrgMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list org members params
func (o *ListOrgMembersParams) WithContext(ctx context.Context) *ListOrgMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list org members params
func (o *ListOrgMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list org members params
func (o *ListOrgMembersParams) WithHTTPClient(client *http.Client) *ListOrgMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list org members params
func (o *ListOrgMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationToken adds the organizationToken to the list org members params
func (o *ListOrgMembersParams) WithOrganizationToken(organizationToken string) *ListOrgMembersParams {
	o.SetOrganizationToken(organizationToken)
	return o
}

// SetOrganizationToken adds the organizationToken to the list org members params
func (o *ListOrgMembersParams) SetOrganizationToken(organizationToken string) {
	o.OrganizationToken = organizationToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrgMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationToken
	if err := r.SetPathParam("organizationToken", o.OrganizationToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
