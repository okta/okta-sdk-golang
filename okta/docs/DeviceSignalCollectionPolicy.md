# DeviceSignalCollectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to **NullableString** | Policy conditions aren&#39;t supported. Conditions are applied at the rule level for this policy type. | [optional] 
**Embedded** | Pointer to [**AccessPolicyAllOfEmbedded**](AccessPolicyAllOfEmbedded.md) |  | [optional] 

## Methods

### NewDeviceSignalCollectionPolicy

`func NewDeviceSignalCollectionPolicy() *DeviceSignalCollectionPolicy`

NewDeviceSignalCollectionPolicy instantiates a new DeviceSignalCollectionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSignalCollectionPolicyWithDefaults

`func NewDeviceSignalCollectionPolicyWithDefaults() *DeviceSignalCollectionPolicy`

NewDeviceSignalCollectionPolicyWithDefaults instantiates a new DeviceSignalCollectionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *DeviceSignalCollectionPolicy) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DeviceSignalCollectionPolicy) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DeviceSignalCollectionPolicy) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DeviceSignalCollectionPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *DeviceSignalCollectionPolicy) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *DeviceSignalCollectionPolicy) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetEmbedded

`func (o *DeviceSignalCollectionPolicy) GetEmbedded() AccessPolicyAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DeviceSignalCollectionPolicy) GetEmbeddedOk() (*AccessPolicyAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DeviceSignalCollectionPolicy) SetEmbedded(v AccessPolicyAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DeviceSignalCollectionPolicy) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


