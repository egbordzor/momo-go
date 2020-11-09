package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ErrorReason error reason
//
// swagger:model ErrorReason
type ErrorReason struct {

	// code
	// Enum: [PAYEE_NOT_FOUND PAYER_NOT_FOUND NOT_ALLOWED NOT_ALLOWED_TARGET_ENVIRONMENT INVALID_CALLBACK_URL_HOST INVALID_CURRENCY SERVICE_UNAVAILABLE INTERNAL_PROCESSING_ERROR NOT_ENOUGH_FUNDS PAYER_LIMIT_REACHED PAYEE_NOT_ALLOWED_TO_RECEIVE PAYMENT_NOT_APPROVED RESOURCE_NOT_FOUND APPROVAL_REJECTED EXPIRED TRANSACTION_CANCELED RESOURCE_ALREADY_EXIST]
	Code string `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this error reason
func (m *ErrorReason) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var errorReasonTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PAYEE_NOT_FOUND","PAYER_NOT_FOUND","NOT_ALLOWED","NOT_ALLOWED_TARGET_ENVIRONMENT","INVALID_CALLBACK_URL_HOST","INVALID_CURRENCY","SERVICE_UNAVAILABLE","INTERNAL_PROCESSING_ERROR","NOT_ENOUGH_FUNDS","PAYER_LIMIT_REACHED","PAYEE_NOT_ALLOWED_TO_RECEIVE","PAYMENT_NOT_APPROVED","RESOURCE_NOT_FOUND","APPROVAL_REJECTED","EXPIRED","TRANSACTION_CANCELED","RESOURCE_ALREADY_EXIST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorReasonTypeCodePropEnum = append(errorReasonTypeCodePropEnum, v)
	}
}

const (

	// ErrorReasonCodePAYEENOTFOUND captures enum value "PAYEE_NOT_FOUND"
	ErrorReasonCodePAYEENOTFOUND string = "PAYEE_NOT_FOUND"

	// ErrorReasonCodePAYERNOTFOUND captures enum value "PAYER_NOT_FOUND"
	ErrorReasonCodePAYERNOTFOUND string = "PAYER_NOT_FOUND"

	// ErrorReasonCodeNOTALLOWED captures enum value "NOT_ALLOWED"
	ErrorReasonCodeNOTALLOWED string = "NOT_ALLOWED"

	// ErrorReasonCodeNOTALLOWEDTARGETENVIRONMENT captures enum value "NOT_ALLOWED_TARGET_ENVIRONMENT"
	ErrorReasonCodeNOTALLOWEDTARGETENVIRONMENT string = "NOT_ALLOWED_TARGET_ENVIRONMENT"

	// ErrorReasonCodeINVALIDCALLBACKURLHOST captures enum value "INVALID_CALLBACK_URL_HOST"
	ErrorReasonCodeINVALIDCALLBACKURLHOST string = "INVALID_CALLBACK_URL_HOST"

	// ErrorReasonCodeINVALIDCURRENCY captures enum value "INVALID_CURRENCY"
	ErrorReasonCodeINVALIDCURRENCY string = "INVALID_CURRENCY"

	// ErrorReasonCodeSERVICEUNAVAILABLE captures enum value "SERVICE_UNAVAILABLE"
	ErrorReasonCodeSERVICEUNAVAILABLE string = "SERVICE_UNAVAILABLE"

	// ErrorReasonCodeINTERNALPROCESSINGERROR captures enum value "INTERNAL_PROCESSING_ERROR"
	ErrorReasonCodeINTERNALPROCESSINGERROR string = "INTERNAL_PROCESSING_ERROR"

	// ErrorReasonCodeNOTENOUGHFUNDS captures enum value "NOT_ENOUGH_FUNDS"
	ErrorReasonCodeNOTENOUGHFUNDS string = "NOT_ENOUGH_FUNDS"

	// ErrorReasonCodePAYERLIMITREACHED captures enum value "PAYER_LIMIT_REACHED"
	ErrorReasonCodePAYERLIMITREACHED string = "PAYER_LIMIT_REACHED"

	// ErrorReasonCodePAYEENOTALLOWEDTORECEIVE captures enum value "PAYEE_NOT_ALLOWED_TO_RECEIVE"
	ErrorReasonCodePAYEENOTALLOWEDTORECEIVE string = "PAYEE_NOT_ALLOWED_TO_RECEIVE"

	// ErrorReasonCodePAYMENTNOTAPPROVED captures enum value "PAYMENT_NOT_APPROVED"
	ErrorReasonCodePAYMENTNOTAPPROVED string = "PAYMENT_NOT_APPROVED"

	// ErrorReasonCodeRESOURCENOTFOUND captures enum value "RESOURCE_NOT_FOUND"
	ErrorReasonCodeRESOURCENOTFOUND string = "RESOURCE_NOT_FOUND"

	// ErrorReasonCodeAPPROVALREJECTED captures enum value "APPROVAL_REJECTED"
	ErrorReasonCodeAPPROVALREJECTED string = "APPROVAL_REJECTED"

	// ErrorReasonCodeEXPIRED captures enum value "EXPIRED"
	ErrorReasonCodeEXPIRED string = "EXPIRED"

	// ErrorReasonCodeTRANSACTIONCANCELED captures enum value "TRANSACTION_CANCELED"
	ErrorReasonCodeTRANSACTIONCANCELED string = "TRANSACTION_CANCELED"

	// ErrorReasonCodeRESOURCEALREADYEXIST captures enum value "RESOURCE_ALREADY_EXIST"
	ErrorReasonCodeRESOURCEALREADYEXIST string = "RESOURCE_ALREADY_EXIST"
)

// prop value enum
func (m *ErrorReason) validateCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, errorReasonTypeCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ErrorReason) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorReason) UnmarshalBinary(b []byte) error {
	var res ErrorReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
