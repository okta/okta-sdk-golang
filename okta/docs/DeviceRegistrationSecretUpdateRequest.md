# DeviceRegistrationSecretUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Admin-friendly label for this registration secret. Must be unique per org. | [optional] 
**MaxRegistrations** | Pointer to **int32** | Maximum number of successful device registrations. May be raised or lowered. Lowering below the current &#x60;registrationCount&#x60; is allowed and effectively soft-disables further registrations against this secret.  | [optional] 
**Status** | Pointer to **string** | Status of the registration secret. * &#x60;ACTIVE&#x60; — The secret can be used for device registration. * &#x60;INACTIVE&#x60; — The secret can&#39;t be used for device registration. Any token request presenting a client assertion signed with an inactive secret is rejected.  | [optional] 

## Methods

### NewDeviceRegistrationSecretUpdateRequest

`func NewDeviceRegistrationSecretUpdateRequest() *DeviceRegistrationSecretUpdateRequest`

NewDeviceRegistrationSecretUpdateRequest instantiates a new DeviceRegistrationSecretUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRegistrationSecretUpdateRequestWithDefaults

`func NewDeviceRegistrationSecretUpdateRequestWithDefaults() *DeviceRegistrationSecretUpdateRequest`

NewDeviceRegistrationSecretUpdateRequestWithDefaults instantiates a new DeviceRegistrationSecretUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DeviceRegistrationSecretUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceRegistrationSecretUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceRegistrationSecretUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceRegistrationSecretUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxRegistrations

`func (o *DeviceRegistrationSecretUpdateRequest) GetMaxRegistrations() int32`

GetMaxRegistrations returns the MaxRegistrations field if non-nil, zero value otherwise.

### GetMaxRegistrationsOk

`func (o *DeviceRegistrationSecretUpdateRequest) GetMaxRegistrationsOk() (*int32, bool)`

GetMaxRegistrationsOk returns a tuple with the MaxRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRegistrations

`func (o *DeviceRegistrationSecretUpdateRequest) SetMaxRegistrations(v int32)`

SetMaxRegistrations sets MaxRegistrations field to given value.

### HasMaxRegistrations

`func (o *DeviceRegistrationSecretUpdateRequest) HasMaxRegistrations() bool`

HasMaxRegistrations returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceRegistrationSecretUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceRegistrationSecretUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceRegistrationSecretUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceRegistrationSecretUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


