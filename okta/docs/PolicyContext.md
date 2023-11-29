# PolicyContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**PolicyContextDevice**](PolicyContextDevice.md) |  | [optional] 
**Groups** | [**PolicyContextGroups**](PolicyContextGroups.md) |  | 
**Ip** | Pointer to **string** | The network rule condition, zone, or IP address | [optional] 
**Risk** | Pointer to [**PolicyContextRisk**](PolicyContextRisk.md) |  | [optional] 
**User** | [**PolicyContextUser**](PolicyContextUser.md) |  | 
**Zones** | Pointer to [**PolicyContextZones**](PolicyContextZones.md) |  | [optional] 

## Methods

### NewPolicyContext

`func NewPolicyContext(groups PolicyContextGroups, user PolicyContextUser, ) *PolicyContext`

NewPolicyContext instantiates a new PolicyContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyContextWithDefaults

`func NewPolicyContextWithDefaults() *PolicyContext`

NewPolicyContextWithDefaults instantiates a new PolicyContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PolicyContext) GetDevice() PolicyContextDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PolicyContext) GetDeviceOk() (*PolicyContextDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PolicyContext) SetDevice(v PolicyContextDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PolicyContext) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetGroups

`func (o *PolicyContext) GetGroups() PolicyContextGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PolicyContext) GetGroupsOk() (*PolicyContextGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PolicyContext) SetGroups(v PolicyContextGroups)`

SetGroups sets Groups field to given value.


### GetIp

`func (o *PolicyContext) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PolicyContext) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PolicyContext) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *PolicyContext) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetRisk

`func (o *PolicyContext) GetRisk() PolicyContextRisk`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *PolicyContext) GetRiskOk() (*PolicyContextRisk, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *PolicyContext) SetRisk(v PolicyContextRisk)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *PolicyContext) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetUser

`func (o *PolicyContext) GetUser() PolicyContextUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PolicyContext) GetUserOk() (*PolicyContextUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PolicyContext) SetUser(v PolicyContextUser)`

SetUser sets User field to given value.


### GetZones

`func (o *PolicyContext) GetZones() PolicyContextZones`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *PolicyContext) GetZonesOk() (*PolicyContextZones, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *PolicyContext) SetZones(v PolicyContextZones)`

SetZones sets Zones field to given value.

### HasZones

`func (o *PolicyContext) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


