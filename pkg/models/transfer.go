package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Transfer transfer
//
// swagger:model Transfer
type Transfer struct {

	// Amount that will be debited from the payer account.
	Amount string `json:"amount,omitempty"`

	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`

	// External id is used as a reference to the transaction. External id is used for reconciliation. The external id will be included in transaction history report. <br>External id is not required to be unique.
	ExternalID string `json:"externalId,omitempty"`

	// payee
	Payee *Party `json:"payee,omitempty"`

	// Message that will be written in the payee transaction history note field.
	PayeeNote string `json:"payeeNote,omitempty"`

	// Message that will be written in the payer transaction history message field.
	PayerMessage string `json:"payerMessage,omitempty"`
}

// Validate validates this transfer
func (m *Transfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transfer) validatePayee(formats strfmt.Registry) error {

	if swag.IsZero(m.Payee) { // not required
		return nil
	}

	if m.Payee != nil {
		if err := m.Payee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payee")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transfer) UnmarshalBinary(b []byte) error {
	var res Transfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
