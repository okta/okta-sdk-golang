# AccessPolicyRuleConditionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**DeviceAccessPolicyRuleCondition**](DeviceAccessPolicyRuleCondition.md) |  | [optional] 
**ElCondition** | Pointer to [**AccessPolicyRuleCustomCondition**](AccessPolicyRuleCustomCondition.md) |  | [optional] 
**UserType** | Pointer to [**UserTypeCondition**](UserTypeCondition.md) |  | [optional] 

## Methods

### NewAccessPolicyRuleConditionsAllOf

`func NewAccessPolicyRuleConditionsAllOf() *AccessPolicyRuleConditionsAllOf`

NewAccessPolicyRuleConditionsAllOf instantiates a new AccessPolicyRuleConditionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyRuleConditionsAllOfWithDefaults

`func NewAccessPolicyRuleConditionsAllOfWithDefaults() *AccessPolicyRuleConditionsAllOf`

NewAccessPolicyRuleConditionsAllOfWithDefaults instantiates a new AccessPolicyRuleConditionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *AccessPolicyRuleConditionsAllOf) GetDevice() DeviceAccessPolicyRuleCondition`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AccessPolicyRuleConditionsAllOf) GetDeviceOk() (*DeviceAccessPolicyRuleCondition, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AccessPolicyRuleConditionsAllOf) SetDevice(v DeviceAccessPolicyRuleCondition)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AccessPolicyRuleConditionsAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetElCondition

`func (o *AccessPolicyRuleConditionsAllOf) GetElCondition() AccessPolicyRuleCustomCondition`

GetElCondition returns the ElCondition field if non-nil, zero value otherwise.

### GetElConditionOk

`func (o *AccessPolicyRuleConditionsAllOf) GetElConditionOk() (*AccessPolicyRuleCustomCondition, bool)`

GetElConditionOk returns a tuple with the ElCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElCondition

`func (o *AccessPolicyRuleConditionsAllOf) SetElCondition(v AccessPolicyRuleCustomCondition)`

SetElCondition sets ElCondition field to given value.

### HasElCondition

`func (o *AccessPolicyRuleConditionsAllOf) HasElCondition() bool`

HasElCondition returns a boolean if a field has been set.

### GetUserType

`func (o *AccessPolicyRuleConditionsAllOf) GetUserType() UserTypeCondition`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *AccessPolicyRuleConditionsAllOf) GetUserTypeOk() (*UserTypeCondition, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *AccessPolicyRuleConditionsAllOf) SetUserType(v UserTypeCondition)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *AccessPolicyRuleConditionsAllOf) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


