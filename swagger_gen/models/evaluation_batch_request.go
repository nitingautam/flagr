// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EvaluationBatchRequest evaluation batch request
//
// swagger:model evaluationBatchRequest
type EvaluationBatchRequest struct {

	// enable debug
	EnableDebug bool `json:"enableDebug,omitempty"`

	// entities
	// Required: true
	// Min Items: 1
	Entities []*EvaluationEntity `json:"entities"`

	// flagIDs
	// Min Items: 1
	FlagIDs []int64 `json:"flagIDs"`

	// flagKeys. Either flagIDs, flagKeys or flagTags works. If pass in multiples, Flagr may return duplicate results.
	// Min Items: 1
	FlagKeys []string `json:"flagKeys"`

	// flagTags. Either flagIDs, flagKeys or flagTags works. If pass in multiples, Flagr may return duplicate results.
	// Min Items: 1
	FlagTags []string `json:"flagTags"`
}

// Validate validates this evaluation batch request
func (m *EvaluationBatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagIDs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvaluationBatchRequest) validateEntities(formats strfmt.Registry) error {

	if err := validate.Required("entities", "body", m.Entities); err != nil {
		return err
	}

	iEntitiesSize := int64(len(m.Entities))

	if err := validate.MinItems("entities", "body", iEntitiesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EvaluationBatchRequest) validateFlagIDs(formats strfmt.Registry) error {

	if swag.IsZero(m.FlagIDs) { // not required
		return nil
	}

	iFlagIDsSize := int64(len(m.FlagIDs))

	if err := validate.MinItems("flagIDs", "body", iFlagIDsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.FlagIDs); i++ {

		if err := validate.MinimumInt("flagIDs"+"."+strconv.Itoa(i), "body", int64(m.FlagIDs[i]), 1, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *EvaluationBatchRequest) validateFlagKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.FlagKeys) { // not required
		return nil
	}

	iFlagKeysSize := int64(len(m.FlagKeys))

	if err := validate.MinItems("flagKeys", "body", iFlagKeysSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.FlagKeys); i++ {

		if err := validate.MinLength("flagKeys"+"."+strconv.Itoa(i), "body", string(m.FlagKeys[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *EvaluationBatchRequest) validateFlagTags(formats strfmt.Registry) error {

	if swag.IsZero(m.FlagTags) { // not required
		return nil
	}

	iFlagTagsSize := int64(len(m.FlagTags))

	if err := validate.MinItems("flagTags", "body", iFlagTagsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.FlagTags); i++ {

		if err := validate.MinLength("flagTags"+"."+strconv.Itoa(i), "body", string(m.FlagTags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationBatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationBatchRequest) UnmarshalBinary(b []byte) error {
	var res EvaluationBatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
