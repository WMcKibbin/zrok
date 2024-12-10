// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListOrgMembersReader is a Reader for the ListOrgMembers structure.
type ListOrgMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrgMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrgMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListOrgMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListOrgMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /members/{organizationToken}] listOrgMembers", response, response.Code())
	}
}

// NewListOrgMembersOK creates a ListOrgMembersOK with default headers values
func NewListOrgMembersOK() *ListOrgMembersOK {
	return &ListOrgMembersOK{}
}

/*
ListOrgMembersOK describes a response with status code 200, with default header values.

ok
*/
type ListOrgMembersOK struct {
	Payload *ListOrgMembersOKBody
}

// IsSuccess returns true when this list org members o k response has a 2xx status code
func (o *ListOrgMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list org members o k response has a 3xx status code
func (o *ListOrgMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list org members o k response has a 4xx status code
func (o *ListOrgMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list org members o k response has a 5xx status code
func (o *ListOrgMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list org members o k response a status code equal to that given
func (o *ListOrgMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list org members o k response
func (o *ListOrgMembersOK) Code() int {
	return 200
}

func (o *ListOrgMembersOK) Error() string {
	return fmt.Sprintf("[GET /members/{organizationToken}][%d] listOrgMembersOK  %+v", 200, o.Payload)
}

func (o *ListOrgMembersOK) String() string {
	return fmt.Sprintf("[GET /members/{organizationToken}][%d] listOrgMembersOK  %+v", 200, o.Payload)
}

func (o *ListOrgMembersOK) GetPayload() *ListOrgMembersOKBody {
	return o.Payload
}

func (o *ListOrgMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListOrgMembersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrgMembersNotFound creates a ListOrgMembersNotFound with default headers values
func NewListOrgMembersNotFound() *ListOrgMembersNotFound {
	return &ListOrgMembersNotFound{}
}

/*
ListOrgMembersNotFound describes a response with status code 404, with default header values.

not found
*/
type ListOrgMembersNotFound struct {
}

// IsSuccess returns true when this list org members not found response has a 2xx status code
func (o *ListOrgMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list org members not found response has a 3xx status code
func (o *ListOrgMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list org members not found response has a 4xx status code
func (o *ListOrgMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list org members not found response has a 5xx status code
func (o *ListOrgMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list org members not found response a status code equal to that given
func (o *ListOrgMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list org members not found response
func (o *ListOrgMembersNotFound) Code() int {
	return 404
}

func (o *ListOrgMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /members/{organizationToken}][%d] listOrgMembersNotFound ", 404)
}

func (o *ListOrgMembersNotFound) String() string {
	return fmt.Sprintf("[GET /members/{organizationToken}][%d] listOrgMembersNotFound ", 404)
}

func (o *ListOrgMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListOrgMembersInternalServerError creates a ListOrgMembersInternalServerError with default headers values
func NewListOrgMembersInternalServerError() *ListOrgMembersInternalServerError {
	return &ListOrgMembersInternalServerError{}
}

/*
ListOrgMembersInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type ListOrgMembersInternalServerError struct {
}

// IsSuccess returns true when this list org members internal server error response has a 2xx status code
func (o *ListOrgMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list org members internal server error response has a 3xx status code
func (o *ListOrgMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list org members internal server error response has a 4xx status code
func (o *ListOrgMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list org members internal server error response has a 5xx status code
func (o *ListOrgMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list org members internal server error response a status code equal to that given
func (o *ListOrgMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list org members internal server error response
func (o *ListOrgMembersInternalServerError) Code() int {
	return 500
}

func (o *ListOrgMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /members/{organizationToken}][%d] listOrgMembersInternalServerError ", 500)
}

func (o *ListOrgMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /members/{organizationToken}][%d] listOrgMembersInternalServerError ", 500)
}

func (o *ListOrgMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
ListOrgMembersOKBody list org members o k body
swagger:model ListOrgMembersOKBody
*/
type ListOrgMembersOKBody struct {

	// members
	Members []*ListOrgMembersOKBodyMembersItems0 `json:"members"`
}

// Validate validates this list org members o k body
func (o *ListOrgMembersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListOrgMembersOKBody) validateMembers(formats strfmt.Registry) error {
	if swag.IsZero(o.Members) { // not required
		return nil
	}

	for i := 0; i < len(o.Members); i++ {
		if swag.IsZero(o.Members[i]) { // not required
			continue
		}

		if o.Members[i] != nil {
			if err := o.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listOrgMembersOK" + "." + "members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listOrgMembersOK" + "." + "members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list org members o k body based on the context it is used
func (o *ListOrgMembersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListOrgMembersOKBody) contextValidateMembers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Members); i++ {

		if o.Members[i] != nil {

			if swag.IsZero(o.Members[i]) { // not required
				return nil
			}

			if err := o.Members[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listOrgMembersOK" + "." + "members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listOrgMembersOK" + "." + "members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListOrgMembersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOrgMembersOKBody) UnmarshalBinary(b []byte) error {
	var res ListOrgMembersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListOrgMembersOKBodyMembersItems0 list org members o k body members items0
swagger:model ListOrgMembersOKBodyMembersItems0
*/
type ListOrgMembersOKBodyMembersItems0 struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// email
	Email string `json:"email,omitempty"`
}

// Validate validates this list org members o k body members items0
func (o *ListOrgMembersOKBodyMembersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list org members o k body members items0 based on context it is used
func (o *ListOrgMembersOKBodyMembersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOrgMembersOKBodyMembersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOrgMembersOKBodyMembersItems0) UnmarshalBinary(b []byte) error {
	var res ListOrgMembersOKBodyMembersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
