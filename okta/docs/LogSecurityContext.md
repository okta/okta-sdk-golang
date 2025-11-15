# LogSecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to **NullableInt32** | The [Autonomous system](https://docs.telemetry.mozilla.org/datasets/other/asn_aggregates/reference) number that&#39;s associated with the autonomous system the event request was sourced to | [optional] [readonly] 
**AsOrg** | Pointer to **NullableString** | The organization that is associated with the autonomous system that the event request is sourced to | [optional] [readonly] 
**Domain** | Pointer to **NullableString** | The domain name that&#39;s associated with the IP address of the inbound event request | [optional] [readonly] 
**Isp** | Pointer to **NullableString** | The Internet service provider that&#39;s used to send the event&#39;s request | [optional] [readonly] 
**IsProxy** | Pointer to **NullableBool** | Specifies whether an event&#39;s request is from a known proxy | [optional] [readonly] 
**UserBehaviors** | Pointer to **[]string** | The result of the user behavior detection models associated with the event | [optional] [readonly] 

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

### SetAsNumberNil

`func (o *LogSecurityContext) SetAsNumberNil(b bool)`

 SetAsNumberNil sets the value for AsNumber to be an explicit nil

### UnsetAsNumber
`func (o *LogSecurityContext) UnsetAsNumber()`

UnsetAsNumber ensures that no value is present for AsNumber, not even an explicit nil
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

### SetAsOrgNil

`func (o *LogSecurityContext) SetAsOrgNil(b bool)`

 SetAsOrgNil sets the value for AsOrg to be an explicit nil

### UnsetAsOrg
`func (o *LogSecurityContext) UnsetAsOrg()`

UnsetAsOrg ensures that no value is present for AsOrg, not even an explicit nil
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

### SetDomainNil

`func (o *LogSecurityContext) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *LogSecurityContext) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
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

### SetIspNil

`func (o *LogSecurityContext) SetIspNil(b bool)`

 SetIspNil sets the value for Isp to be an explicit nil

### UnsetIsp
`func (o *LogSecurityContext) UnsetIsp()`

UnsetIsp ensures that no value is present for Isp, not even an explicit nil
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

### SetIsProxyNil

`func (o *LogSecurityContext) SetIsProxyNil(b bool)`

 SetIsProxyNil sets the value for IsProxy to be an explicit nil

### UnsetIsProxy
`func (o *LogSecurityContext) UnsetIsProxy()`

UnsetIsProxy ensures that no value is present for IsProxy, not even an explicit nil
### GetUserBehaviors

`func (o *LogSecurityContext) GetUserBehaviors() []string`

GetUserBehaviors returns the UserBehaviors field if non-nil, zero value otherwise.

### GetUserBehaviorsOk

`func (o *LogSecurityContext) GetUserBehaviorsOk() (*[]string, bool)`

GetUserBehaviorsOk returns a tuple with the UserBehaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBehaviors

`func (o *LogSecurityContext) SetUserBehaviors(v []string)`

SetUserBehaviors sets UserBehaviors field to given value.

### HasUserBehaviors

`func (o *LogSecurityContext) HasUserBehaviors() bool`

HasUserBehaviors returns a boolean if a field has been set.

### SetUserBehaviorsNil

`func (o *LogSecurityContext) SetUserBehaviorsNil(b bool)`

 SetUserBehaviorsNil sets the value for UserBehaviors to be an explicit nil

### UnsetUserBehaviors
`func (o *LogSecurityContext) UnsetUserBehaviors()`

UnsetUserBehaviors ensures that no value is present for UserBehaviors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


