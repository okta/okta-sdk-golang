# TenantSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceProperties** | Pointer to [**[]AppInstanceProperty**](AppInstanceProperty.md) |  | [optional] 

## Methods

### NewTenantSettings

`func NewTenantSettings() *TenantSettings`

NewTenantSettings instantiates a new TenantSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSettingsWithDefaults

`func NewTenantSettingsWithDefaults() *TenantSettings`

NewTenantSettingsWithDefaults instantiates a new TenantSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceProperties

`func (o *TenantSettings) GetAppInstanceProperties() []AppInstanceProperty`

GetAppInstanceProperties returns the AppInstanceProperties field if non-nil, zero value otherwise.

### GetAppInstancePropertiesOk

`func (o *TenantSettings) GetAppInstancePropertiesOk() (*[]AppInstanceProperty, bool)`

GetAppInstancePropertiesOk returns a tuple with the AppInstanceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceProperties

`func (o *TenantSettings) SetAppInstanceProperties(v []AppInstanceProperty)`

SetAppInstanceProperties sets AppInstanceProperties field to given value.

### HasAppInstanceProperties

`func (o *TenantSettings) HasAppInstanceProperties() bool`

HasAppInstanceProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


