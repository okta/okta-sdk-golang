# PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **NullableTime** | Timestamp when the rule was created | [optional] [readonly] 
**Id** | Pointer to **string** | Identifier for the rule | [optional] 
**LastUpdated** | Pointer to **NullableTime** | Timestamp when the rule was last modified | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the rule | [optional] 
**Priority** | Pointer to **int32** | Priority of the rule | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**System** | Pointer to **bool** | Specifies whether Okta created the Policy Rule (&#x60;system&#x3D;true&#x60;). You can&#39;t delete Policy Rules that have &#x60;system&#x60; set to &#x60;true&#x60;. | [optional] [default to false]
**Type** | Pointer to **string** | Rule type | [optional] 

## Methods

### NewPolicyRule

`func NewPolicyRule() *PolicyRule`

NewPolicyRule instantiates a new PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleWithDefaults

`func NewPolicyRuleWithDefaults() *PolicyRule`

NewPolicyRuleWithDefaults instantiates a new PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PolicyRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PolicyRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PolicyRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PolicyRule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *PolicyRule) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PolicyRule) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetId

`func (o *PolicyRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyRule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *PolicyRule) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PolicyRule) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetName

`func (o *PolicyRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *PolicyRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PolicyRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PolicyRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PolicyRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *PolicyRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PolicyRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *PolicyRule) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *PolicyRule) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *PolicyRule) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *PolicyRule) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *PolicyRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


