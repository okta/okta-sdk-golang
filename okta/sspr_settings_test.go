package okta

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// selfServicePasswordResetActionJSON is the action as the Management API returns
// it for a PASSWORD policy rule that carries the settings object.
const selfServicePasswordResetActionJSON = `{
  "access": "ALLOW",
  "requirement": {
    "primary": { "methods": ["email"] },
    "stepUp": { "required": false },
    "accessControl": "LEGACY"
  },
  "settings": { "allowRecoveryEmailWithoutEnrollment": true }
}`

func TestSelfServicePasswordResetActionUnmarshalsSettings(t *testing.T) {
	var action SelfServicePasswordResetAction
	require.NoError(t, json.Unmarshal([]byte(selfServicePasswordResetActionJSON), &action))

	require.NotNil(t, action.Settings, "settings must unmarshal into the typed field")
	assert.True(t, action.Settings.GetAllowRecoveryEmailWithoutEnrollment())
	assert.NotContains(t, action.AdditionalProperties, "settings",
		"settings is modelled, so it must not fall through to AdditionalProperties")
}

func TestSelfServicePasswordResetActionMarshalsSettings(t *testing.T) {
	settings := NewSsprSettings()
	settings.SetAllowRecoveryEmailWithoutEnrollment(true)

	action := NewSelfServicePasswordResetAction()
	action.SetAccess("ALLOW")
	action.SetSettings(*settings)

	body, err := json.Marshal(action)
	require.NoError(t, err)

	var wire map[string]interface{}
	require.NoError(t, json.Unmarshal(body, &wire))
	require.Contains(t, wire, "settings")
	assert.Equal(t, map[string]interface{}{"allowRecoveryEmailWithoutEnrollment": true}, wire["settings"])
}

// TestSelfServicePasswordResetActionRoundTripsSettings holds the guarantee that
// callers depend on: read the action, write it back, and the API sees what it
// sent.
func TestSelfServicePasswordResetActionRoundTripsSettings(t *testing.T) {
	var action SelfServicePasswordResetAction
	require.NoError(t, json.Unmarshal([]byte(selfServicePasswordResetActionJSON), &action))

	body, err := json.Marshal(action)
	require.NoError(t, err)

	var want, got interface{}
	require.NoError(t, json.Unmarshal([]byte(selfServicePasswordResetActionJSON), &want))
	require.NoError(t, json.Unmarshal(body, &got))
	assert.Equal(t, want, got)
}

func TestSelfServicePasswordResetActionOmitsUnsetSettings(t *testing.T) {
	action := NewSelfServicePasswordResetAction()
	action.SetAccess("ALLOW")

	body, err := json.Marshal(action)
	require.NoError(t, err)
	assert.NotContains(t, string(body), "settings")
}

func TestSsprSettingsAccessors(t *testing.T) {
	settings := NewSsprSettings()

	assert.False(t, settings.HasAllowRecoveryEmailWithoutEnrollment())
	assert.False(t, settings.GetAllowRecoveryEmailWithoutEnrollment(), "an unset property reads as the zero value")
	value, ok := settings.GetAllowRecoveryEmailWithoutEnrollmentOk()
	assert.Nil(t, value)
	assert.False(t, ok)

	settings.SetAllowRecoveryEmailWithoutEnrollment(false)

	assert.True(t, settings.HasAllowRecoveryEmailWithoutEnrollment(), "false is a set value, not an absent one")
	value, ok = settings.GetAllowRecoveryEmailWithoutEnrollmentOk()
	require.True(t, ok)
	require.NotNil(t, value)
	assert.False(t, *value)

	body, err := json.Marshal(settings)
	require.NoError(t, err)
	assert.JSONEq(t, `{"allowRecoveryEmailWithoutEnrollment":false}`, string(body),
		"an explicit false must reach the API, because the property is a pointer")
}

// TestSsprSettingsKeepsUnknownProperties covers the other properties that the
// API can hold under settings and that this SDK does not model yet.
func TestSsprSettingsKeepsUnknownProperties(t *testing.T) {
	const withUnknown = `{"allowRecoveryEmailWithoutEnrollment":true,"unmodeledProperty":"keep me"}`

	var settings SsprSettings
	require.NoError(t, json.Unmarshal([]byte(withUnknown), &settings))
	assert.Equal(t, "keep me", settings.AdditionalProperties["unmodeledProperty"])

	body, err := json.Marshal(settings)
	require.NoError(t, err)
	assert.JSONEq(t, withUnknown, string(body))
}
