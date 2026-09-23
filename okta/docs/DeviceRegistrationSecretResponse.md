# DeviceRegistrationSecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the registration secret was created | [optional] [readonly] 
**Description** | Pointer to **string** | Admin-friendly label for this registration secret | [optional] 
**Id** | Pointer to **string** | Unique identifier for the registration secret (prefix &#x60;drs&#x60;) | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the registration secret was last updated | [optional] [readonly] 
**MaxRegistrations** | Pointer to **int32** | Maximum number of successful device registrations allowed for this secret | [optional] 
**RegistrationCount** | Pointer to **int32** | Number of successful HMAC token issuances against this secret. Monotonically increasing for the lifetime of the secret.  &gt; **Note:** A device that obtains a token but doesn&#39;t complete &#x60;/idp/v2/device/register&#x60; leaves its grant &#x60;PENDING&#x60; while still consuming a slot of the cap.  | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the registration secret. * &#x60;ACTIVE&#x60; — The secret can be used for device registration. * &#x60;INACTIVE&#x60; — The secret can&#39;t be used for device registration. Any token request presenting a client assertion signed with an inactive secret is rejected.  | [optional] 

## Methods

### NewDeviceRegistrationSecretResponse

`func NewDeviceRegistrationSecretResponse() *DeviceRegistrationSecretResponse`

NewDeviceRegistrationSecretResponse instantiates a new DeviceRegistrationSecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRegistrationSecretResponseWithDefaults

`func NewDeviceRegistrationSecretResponseWithDefaults() *DeviceRegistrationSecretResponse`

NewDeviceRegistrationSecretResponseWithDefaults instantiates a new DeviceRegistrationSecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeviceRegistrationSecretResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceRegistrationSecretResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceRegistrationSecretResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeviceRegistrationSecretResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceRegistrationSecretResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceRegistrationSecretResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceRegistrationSecretResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceRegistrationSecretResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DeviceRegistrationSecretResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceRegistrationSecretResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceRegistrationSecretResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceRegistrationSecretResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DeviceRegistrationSecretResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceRegistrationSecretResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceRegistrationSecretResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeviceRegistrationSecretResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMaxRegistrations

`func (o *DeviceRegistrationSecretResponse) GetMaxRegistrations() int32`

GetMaxRegistrations returns the MaxRegistrations field if non-nil, zero value otherwise.

### GetMaxRegistrationsOk

`func (o *DeviceRegistrationSecretResponse) GetMaxRegistrationsOk() (*int32, bool)`

GetMaxRegistrationsOk returns a tuple with the MaxRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRegistrations

`func (o *DeviceRegistrationSecretResponse) SetMaxRegistrations(v int32)`

SetMaxRegistrations sets MaxRegistrations field to given value.

### HasMaxRegistrations

`func (o *DeviceRegistrationSecretResponse) HasMaxRegistrations() bool`

HasMaxRegistrations returns a boolean if a field has been set.

### GetRegistrationCount

`func (o *DeviceRegistrationSecretResponse) GetRegistrationCount() int32`

GetRegistrationCount returns the RegistrationCount field if non-nil, zero value otherwise.

### GetRegistrationCountOk

`func (o *DeviceRegistrationSecretResponse) GetRegistrationCountOk() (*int32, bool)`

GetRegistrationCountOk returns a tuple with the RegistrationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationCount

`func (o *DeviceRegistrationSecretResponse) SetRegistrationCount(v int32)`

SetRegistrationCount sets RegistrationCount field to given value.

### HasRegistrationCount

`func (o *DeviceRegistrationSecretResponse) HasRegistrationCount() bool`

HasRegistrationCount returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceRegistrationSecretResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceRegistrationSecretResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceRegistrationSecretResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceRegistrationSecretResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


