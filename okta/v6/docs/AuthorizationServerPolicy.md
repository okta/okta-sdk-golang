# AuthorizationServerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the Policy | [optional] 
**Type** | Pointer to **string** | Indicates that the Policy is an authorization server Policy | [optional] 
**Name** | Pointer to **string** | Name of the Policy | [optional] 
**Conditions** | Pointer to [**AuthorizationServerPolicyConditions**](AuthorizationServerPolicyConditions.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the Policy | [optional] 
**Priority** | Pointer to **int32** | Specifies the order in which this Policy is evaluated in relation to the other Policies in a custom authorization server | [optional] 
**Status** | Pointer to **string** | Specifies whether requests have access to this Policy | [optional] 
**System** | Pointer to **bool** | Specifies whether Okta created this Policy | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the Policy was created | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Policy was last updated | [optional] [readonly] 
**Links** | Pointer to [**AuthorizationServerPolicyAllOfLinks**](AuthorizationServerPolicyAllOfLinks.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicy

`func NewAuthorizationServerPolicy() *AuthorizationServerPolicy`

NewAuthorizationServerPolicy instantiates a new AuthorizationServerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyWithDefaults

`func NewAuthorizationServerPolicyWithDefaults() *AuthorizationServerPolicy`

NewAuthorizationServerPolicyWithDefaults instantiates a new AuthorizationServerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorizationServerPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationServerPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationServerPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationServerPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AuthorizationServerPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationServerPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationServerPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizationServerPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AuthorizationServerPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizationServerPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizationServerPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizationServerPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConditions

`func (o *AuthorizationServerPolicy) GetConditions() AuthorizationServerPolicyConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizationServerPolicy) GetConditionsOk() (*AuthorizationServerPolicyConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizationServerPolicy) SetConditions(v AuthorizationServerPolicyConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthorizationServerPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizationServerPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizationServerPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizationServerPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizationServerPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *AuthorizationServerPolicy) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AuthorizationServerPolicy) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AuthorizationServerPolicy) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AuthorizationServerPolicy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorizationServerPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationServerPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationServerPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorizationServerPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *AuthorizationServerPolicy) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AuthorizationServerPolicy) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AuthorizationServerPolicy) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AuthorizationServerPolicy) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetCreated

`func (o *AuthorizationServerPolicy) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorizationServerPolicy) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorizationServerPolicy) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorizationServerPolicy) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthorizationServerPolicy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthorizationServerPolicy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthorizationServerPolicy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthorizationServerPolicy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *AuthorizationServerPolicy) GetLinks() AuthorizationServerPolicyAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizationServerPolicy) GetLinksOk() (*AuthorizationServerPolicyAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizationServerPolicy) SetLinks(v AuthorizationServerPolicyAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizationServerPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


