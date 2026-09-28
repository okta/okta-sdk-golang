# DeviceRegistrationSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Admin-friendly label for this registration secret. Must be unique per org. | 
**MaxRegistrations** | Pointer to **int32** | Maximum number of successful device registrations this secret may be used for before further registrations are rejected. The cap bounds the blast radius of a leaked secret. Admins must size it to the fleet they intend to register with this secret.  | [optional] [default to 10000]

## Methods

### NewDeviceRegistrationSecretRequest

`func NewDeviceRegistrationSecretRequest(description string, ) *DeviceRegistrationSecretRequest`

NewDeviceRegistrationSecretRequest instantiates a new DeviceRegistrationSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRegistrationSecretRequestWithDefaults

`func NewDeviceRegistrationSecretRequestWithDefaults() *DeviceRegistrationSecretRequest`

NewDeviceRegistrationSecretRequestWithDefaults instantiates a new DeviceRegistrationSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DeviceRegistrationSecretRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceRegistrationSecretRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceRegistrationSecretRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMaxRegistrations

`func (o *DeviceRegistrationSecretRequest) GetMaxRegistrations() int32`

GetMaxRegistrations returns the MaxRegistrations field if non-nil, zero value otherwise.

### GetMaxRegistrationsOk

`func (o *DeviceRegistrationSecretRequest) GetMaxRegistrationsOk() (*int32, bool)`

GetMaxRegistrationsOk returns a tuple with the MaxRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRegistrations

`func (o *DeviceRegistrationSecretRequest) SetMaxRegistrations(v int32)`

SetMaxRegistrations sets MaxRegistrations field to given value.

### HasMaxRegistrations

`func (o *DeviceRegistrationSecretRequest) HasMaxRegistrations() bool`

HasMaxRegistrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


