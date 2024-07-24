# PolicyContextDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** | The platform of the device, for example, IOS. | [optional] 
**Registered** | Pointer to **bool** | If the device is registered | [optional] 
**Managed** | Pointer to **bool** | If the device is managed | [optional] 
**AssuranceId** | Pointer to **string** | The device assurance policy ID for the simulation | [optional] 

## Methods

### NewPolicyContextDevice

`func NewPolicyContextDevice() *PolicyContextDevice`

NewPolicyContextDevice instantiates a new PolicyContextDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyContextDeviceWithDefaults

`func NewPolicyContextDeviceWithDefaults() *PolicyContextDevice`

NewPolicyContextDeviceWithDefaults instantiates a new PolicyContextDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *PolicyContextDevice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PolicyContextDevice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PolicyContextDevice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PolicyContextDevice) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRegistered

`func (o *PolicyContextDevice) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *PolicyContextDevice) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *PolicyContextDevice) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *PolicyContextDevice) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetManaged

`func (o *PolicyContextDevice) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PolicyContextDevice) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PolicyContextDevice) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PolicyContextDevice) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetAssuranceId

`func (o *PolicyContextDevice) GetAssuranceId() string`

GetAssuranceId returns the AssuranceId field if non-nil, zero value otherwise.

### GetAssuranceIdOk

`func (o *PolicyContextDevice) GetAssuranceIdOk() (*string, bool)`

GetAssuranceIdOk returns a tuple with the AssuranceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceId

`func (o *PolicyContextDevice) SetAssuranceId(v string)`

SetAssuranceId sets AssuranceId field to given value.

### HasAssuranceId

`func (o *PolicyContextDevice) HasAssuranceId() bool`

HasAssuranceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


