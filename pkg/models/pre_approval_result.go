package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PreApprovalResult pre approval result
//
// swagger:model PreApprovalResult
type PreApprovalResult struct {

	// payer
	Payer *Party `json:"payer,omitempty"`

	// ISO4217 Currency
	PayerCurrency string `json:"payerCurrency,omitempty"`

	// The mesage that is shown to the approver.
	PayerMessage string `json:"payerMessage,omitempty"`

	// reason
	Reason *ErrorReason `json:"reason,omitempty"`

	// status
	// Enum: [PENDING SUCCESSFUL FAILED]
	Status string `json:"status,omitempty"`

	// The request validity time of the pre-approval
	ValidityTime int64 `json:"validityTime,omitempty"`
}

// Validate validates this pre approval result
func (m *PreApprovalResult) Validate(formats strfmt.Registry) error {
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

func (m *PreApprovalResult) validatePayer(formats strfmt.Registry) error {

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

func (m *PreApprovalResult) validateReason(formats strfmt.Registry) error {

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

var preApprovalResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","SUCCESSFUL","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		preApprovalResultTypeStatusPropEnum = append(preApprovalResultTypeStatusPropEnum, v)
	}
}

const (

	// PreApprovalResultStatusPENDING captures enum value "PENDING"
	PreApprovalResultStatusPENDING string = "PENDING"

	// PreApprovalResultStatusSUCCESSFUL captures enum value "SUCCESSFUL"
	PreApprovalResultStatusSUCCESSFUL string = "SUCCESSFUL"

	// PreApprovalResultStatusFAILED captures enum value "FAILED"
	PreApprovalResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *PreApprovalResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, preApprovalResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PreApprovalResult) validateStatus(formats strfmt.Registry) error {

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
func (m *PreApprovalResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreApprovalResult) UnmarshalBinary(b []byte) error {
	var res PreApprovalResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
