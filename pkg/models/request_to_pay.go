package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RequestToPay request to pay
//
// swagger:model RequestToPay
type RequestToPay struct {

	// Amount that will be debited from the payer account.
	Amount string `json:"amount,omitempty"`

	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`

	// External id is used as a reference to the transaction. External id is used for reconciliation. The external id will be included in transaction history report. <br>External id is not required to be unique.
	ExternalID string `json:"externalId,omitempty"`

	// Message that will be written in the payee transaction history note field.
	PayeeNote string `json:"payeeNote,omitempty"`

	// payer
	Payer *Party `json:"payer,omitempty"`

	// Message that will be written in the payer transaction history message field.
	PayerMessage string `json:"payerMessage,omitempty"`
}

// Validate validates this request to pay
func (m *RequestToPay) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RequestToPay) validatePayer(formats strfmt.Registry) error {

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
func (m *RequestToPay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestToPay) UnmarshalBinary(b []byte) error {
	var res RequestToPay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
