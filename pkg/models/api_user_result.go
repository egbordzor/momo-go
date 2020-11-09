package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIUserResult The API user information
//
// swagger:model ApiUserResult
type APIUserResult struct {

	// payment server Url
	PaymentServerURL *PaymentServerURL `json:"paymentServerUrl,omitempty"`

	// The provider callback host
	ProviderCallbackHost string `json:"providerCallbackHost,omitempty"`

	// target environment
	TargetEnvironment *TargetEnvironment `json:"targetEnvironment,omitempty"`
}

// Validate validates this Api user result
func (m *APIUserResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentServerURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIUserResult) validatePaymentServerURL(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentServerURL) { // not required
		return nil
	}

	if m.PaymentServerURL != nil {
		if err := m.PaymentServerURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paymentServerUrl")
			}
			return err
		}
	}

	return nil
}

func (m *APIUserResult) validateTargetEnvironment(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetEnvironment) { // not required
		return nil
	}

	if m.TargetEnvironment != nil {
		if err := m.TargetEnvironment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEnvironment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIUserResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIUserResult) UnmarshalBinary(b []byte) error {
	var res APIUserResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
