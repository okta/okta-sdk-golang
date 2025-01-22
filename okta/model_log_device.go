package okta

import "encoding/json"

// LogDevice struct for LogDevice
type LogDevice struct {
	ID                    *string `json:"id"`
	Name                  *string `json:"name"`
	OsPlatform            *string `json:"os_platform"`
	OsVersion             *string `json:"os_version"`
	Managed               *bool   `json:"managed"`
	Registered            *bool   `json:"registered"`
	DeviceIntegrator      *string `json:"device_integrator"`
	DiskEncryptionType    *string `json:"disk_encryption_type"`
	ScreenLockType        *string `json:"screen_lock_type"`
	Jailbreak             *bool   `json:"jailbreak"`
	SecureHardwarePresent *bool   `json:"secure_hardware_present"`
	AdditionalProperties  map[string]interface{}
}

type _LogDevice LogDevice

// NewLogDevice instantiates a new LogDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogDevice() *LogDevice {
	this := LogDevice{}
	return &this
}

// NewLogDeviceWithDefaults instantiates a new LogDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDeviceWithDefaults() *LogDevice {
	this := LogDevice{}
	return &this
}

func (o LogDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OsPlatform != nil {
		toSerialize["os_platform"] = o.OsPlatform
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	if o.Registered != nil {
		toSerialize["registered"] = o.Registered
	}
	if o.DeviceIntegrator != nil {
		toSerialize["device_integrator"] = o.DeviceIntegrator
	}
	if o.DiskEncryptionType != nil {
		toSerialize["disk_encryption_type"] = o.DiskEncryptionType
	}
	if o.ScreenLockType != nil {
		toSerialize["screen_lock_type"] = o.ScreenLockType
	}
	if o.Jailbreak != nil {
		toSerialize["jailbreak"] = o.Jailbreak
	}
	if o.SecureHardwarePresent != nil {
		toSerialize["secure_hardware_present"] = o.SecureHardwarePresent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogDevice) UnmarshalJSON(bytes []byte) (err error) {
	varLogDevice := _LogDevice{}

	err = json.Unmarshal(bytes, &varLogDevice)
	if err == nil {
		*o = LogDevice(varLogDevice)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "os_platform")
		delete(additionalProperties, "os_version")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "registered")
		delete(additionalProperties, "device_integrator")
		delete(additionalProperties, "disk_encryption_type")
		delete(additionalProperties, "screen_lock_type")
		delete(additionalProperties, "jailbreak")
		delete(additionalProperties, "secure_hardware_present")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogDevice struct {
	value *LogDevice
	isSet bool
}

func (v NullableLogDevice) Get() *LogDevice {
	return v.value
}

func (v *NullableLogDevice) Set(val *LogDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableLogDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableLogDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogDevice(val *LogDevice) *NullableLogDevice {
	return &NullableLogDevice{value: val, isSet: true}
}

func (v NullableLogDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
