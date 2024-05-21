# LogSecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to **int32** |  | [optional] [readonly] 
**AsOrg** | Pointer to **string** |  | [optional] [readonly] 
**Domain** | Pointer to **string** |  | [optional] [readonly] 
**Isp** | Pointer to **string** |  | [optional] [readonly] 
**IsProxy** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewLogSecurityContext

`func NewLogSecurityContext() *LogSecurityContext`

NewLogSecurityContext instantiates a new LogSecurityContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSecurityContextWithDefaults

`func NewLogSecurityContextWithDefaults() *LogSecurityContext`

NewLogSecurityContextWithDefaults instantiates a new LogSecurityContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *LogSecurityContext) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *LogSecurityContext) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *LogSecurityContext) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *LogSecurityContext) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetAsOrg

`func (o *LogSecurityContext) GetAsOrg() string`

GetAsOrg returns the AsOrg field if non-nil, zero value otherwise.

### GetAsOrgOk

`func (o *LogSecurityContext) GetAsOrgOk() (*string, bool)`

GetAsOrgOk returns a tuple with the AsOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOrg

`func (o *LogSecurityContext) SetAsOrg(v string)`

SetAsOrg sets AsOrg field to given value.

### HasAsOrg

`func (o *LogSecurityContext) HasAsOrg() bool`

HasAsOrg returns a boolean if a field has been set.

### GetDomain

`func (o *LogSecurityContext) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LogSecurityContext) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LogSecurityContext) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LogSecurityContext) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIsp

`func (o *LogSecurityContext) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *LogSecurityContext) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *LogSecurityContext) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *LogSecurityContext) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIsProxy

`func (o *LogSecurityContext) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *LogSecurityContext) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *LogSecurityContext) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *LogSecurityContext) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


