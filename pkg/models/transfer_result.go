package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransferResult transfer result
//
// swagger:model TransferResult
type TransferResult struct {

	// Amount that will be debited from the payer account.
	Amount string `json:"amount,omitempty"`

	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`

	// External id is used as a reference to the transaction. External id is used for reconciliation. The external id will be included in transaction history report. <br>External id is not required to be unique.
	ExternalID string `json:"externalId,omitempty"`

	// Financial transactionIdd from mobile money manager.<br> Used to connect to the specific financial transaction made in the account
	FinancialTransactionID string `json:"financialTransactionId,omitempty"`

	// payee
	Payee *Party `json:"payee,omitempty"`

	// Message that will be written in the payee transaction history note field.
	PayeeNote string `json:"payeeNote,omitempty"`

	// Message that will be written in the payer transaction history message field.
	PayerMessage string `json:"payerMessage,omitempty"`

	// reason
	Reason *ErrorReason `json:"reason,omitempty"`

	// status
	// Enum: [PENDING SUCCESSFUL FAILED]
	Status string `json:"status,omitempty"`
}

// Validate validates this transfer result
func (m *TransferResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransferResult) validatePayee(formats strfmt.Registry) error {

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

func (m *TransferResult) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if m.Reason != nil {
		if err := m.Reason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

var transferResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","SUCCESSFUL","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transferResultTypeStatusPropEnum = append(transferResultTypeStatusPropEnum, v)
	}
}

const (

	// TransferResultStatusPENDING captures enum value "PENDING"
	TransferResultStatusPENDING string = "PENDING"

	// TransferResultStatusSUCCESSFUL captures enum value "SUCCESSFUL"
	TransferResultStatusSUCCESSFUL string = "SUCCESSFUL"

	// TransferResultStatusFAILED captures enum value "FAILED"
	TransferResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *TransferResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transferResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransferResult) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransferResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransferResult) UnmarshalBinary(b []byte) error {
	var res TransferResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
