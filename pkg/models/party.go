package models

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Party Party identifies a account holder in the wallet platform. Party consists of two parameters, type and partyId. Each type have its own validation of the partyId<br> MSISDN - Mobile Number validated according to ITU-T E.164. Validated with IsMSISDN<br> EMAIL - Validated to be a valid e-mail format. Validated with IsEmail<br> PARTY_CODE - UUID of the party. Validated with IsUuid
//
// swagger:model Party
type Party struct {

	// party Id
	PartyID string `json:"partyId,omitempty"`

	// party Id type
	// Enum: [MSISDN EMAIL PARTY_CODE]
	PartyIDType string `json:"partyIdType,omitempty"`
}

// Validate validates this party
func (m *Party) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartyIDType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var partyTypePartyIDTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MSISDN","EMAIL","PARTY_CODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyTypePartyIDTypePropEnum = append(partyTypePartyIDTypePropEnum, v)
	}
}

const (

	// PartyPartyIDTypeMSISDN captures enum value "MSISDN"
	PartyPartyIDTypeMSISDN string = "MSISDN"

	// PartyPartyIDTypeEMAIL captures enum value "EMAIL"
	PartyPartyIDTypeEMAIL string = "EMAIL"

	// PartyPartyIDTypePARTYCODE captures enum value "PARTY_CODE"
	PartyPartyIDTypePARTYCODE string = "PARTY_CODE"
)

// prop value enum
func (m *Party) validatePartyIDTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, partyTypePartyIDTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Party) validatePartyIDType(formats strfmt.Registry) error {

	if swag.IsZero(m.PartyIDType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePartyIDTypeEnum("partyIdType", "body", m.PartyIDType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Party) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Party) UnmarshalBinary(b []byte) error {
	var res Party
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
