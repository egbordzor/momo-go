package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RequestToPayResult request to pay result
//
// swagger:model RequestToPayResult
type RequestToPayResult struct {

	// Amount that will be debited from the payer account.
	Amount string `json:"amount,omitempty"`

	// ISO4217 Currency
	Currency string `json:"currency,omitempty"`

	// External id provided in the creation of the requestToPay transaction.
	ExternalID string `json:"externalId,omitempty"`

	// Financial transactionIdd from mobile money manager.<br> Used to connect to the specific financial transaction made in the account
	FinancialTransactionID string `json:"financialTransactionId,omitempty"`

	// Message that will be written in the payee transaction history note field.
	PayeeNote string `json:"payeeNote,omitempty"`

	// payer
	Payer *Party `json:"payer,omitempty"`

	// Message that will be written in the payer transaction history message field.
	PayerMessage string `json:"payerMessage,omitempty"`

	// reason
	Reason *ErrorReason `json:"reason,omitempty"`

	// status
	// Enum: [PENDING SUCCESSFUL FAILED]
	Status string `json:"status,omitempty"`
}

// Validate validates this request to pay result
func (m *RequestToPayResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayer(formats); err != nil {
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

func (m *RequestToPayResult) validatePayer(formats strfmt.Registry) error {

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

func (m *RequestToPayResult) validateReason(formats strfmt.Registry) error {

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

var requestToPayResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","SUCCESSFUL","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestToPayResultTypeStatusPropEnum = append(requestToPayResultTypeStatusPropEnum, v)
	}
}

const (

	// RequestToPayResultStatusPENDING captures enum value "PENDING"
	RequestToPayResultStatusPENDING string = "PENDING"

	// RequestToPayResultStatusSUCCESSFUL captures enum value "SUCCESSFUL"
	RequestToPayResultStatusSUCCESSFUL string = "SUCCESSFUL"

	// RequestToPayResultStatusFAILED captures enum value "FAILED"
	RequestToPayResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *RequestToPayResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, requestToPayResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RequestToPayResult) validateStatus(formats strfmt.Registry) error {

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
func (m *RequestToPayResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestToPayResult) UnmarshalBinary(b []byte) error {
	var res RequestToPayResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
