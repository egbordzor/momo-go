package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PreApproval pre approval
//
// swagger:model PreApproval
type PreApproval struct {

	// payer
	Payer *Party `json:"payer,omitempty"`

	// ISO4217 Currency
	PayerCurrency string `json:"payerCurrency,omitempty"`

	// The mesage that is shown to the approver.
	PayerMessage string `json:"payerMessage,omitempty"`

	// The request validity time of the pre-approval
	ValidityTime int64 `json:"validityTime,omitempty"`
}

// Validate validates this pre approval
func (m *PreApproval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreApproval) validatePayer(formats strfmt.Registry) error {

	if swag.IsZero(m.Payer) { // not required
		return nil
	}

	if m.Payer != nil {
		if err := m.Payer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreApproval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreApproval) UnmarshalBinary(b []byte) error {
	var res PreApproval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
