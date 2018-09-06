// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuditLog audit log
// swagger:model auditLog
type AuditLog struct {

	// account id
	AccountID string `json:"account_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// payload
	Payload *AuditLogPayload `json:"payload,omitempty"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AuditLogPayload audit log payload
// swagger:model AuditLogPayload
type AuditLogPayload struct {

	// action
	Action string `json:"action,omitempty"`

	// actor email
	ActorEmail string `json:"actor_email,omitempty"`

	// actor id
	ActorID string `json:"actor_id,omitempty"`

	// actor name
	ActorName string `json:"actor_name,omitempty"`

	// log type
	LogType string `json:"log_type,omitempty"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty"`

	// audit log payload
	AuditLogPayload map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *AuditLogPayload) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// action
		Action string `json:"action,omitempty"`

		// actor email
		ActorEmail string `json:"actor_email,omitempty"`

		// actor id
		ActorID string `json:"actor_id,omitempty"`

		// actor name
		ActorName string `json:"actor_name,omitempty"`

		// log type
		LogType string `json:"log_type,omitempty"`

		// timestamp
		Timestamp string `json:"timestamp,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv AuditLogPayload

	rcv.Action = stage1.Action

	rcv.ActorEmail = stage1.ActorEmail

	rcv.ActorID = stage1.ActorID

	rcv.ActorName = stage1.ActorName

	rcv.LogType = stage1.LogType

	rcv.Timestamp = stage1.Timestamp

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "action")

	delete(stage2, "actor_email")

	delete(stage2, "actor_id")

	delete(stage2, "actor_name")

	delete(stage2, "log_type")

	delete(stage2, "timestamp")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.AuditLogPayload = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m AuditLogPayload) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// action
		Action string `json:"action,omitempty"`

		// actor email
		ActorEmail string `json:"actor_email,omitempty"`

		// actor id
		ActorID string `json:"actor_id,omitempty"`

		// actor name
		ActorName string `json:"actor_name,omitempty"`

		// log type
		LogType string `json:"log_type,omitempty"`

		// timestamp
		Timestamp string `json:"timestamp,omitempty"`
	}

	stage1.Action = m.Action

	stage1.ActorEmail = m.ActorEmail

	stage1.ActorID = m.ActorID

	stage1.ActorName = m.ActorName

	stage1.LogType = m.LogType

	stage1.Timestamp = m.Timestamp

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.AuditLogPayload) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.AuditLogPayload)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this audit log payload
func (m *AuditLogPayload) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditLogPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLogPayload) UnmarshalBinary(b []byte) error {
	var res AuditLogPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
