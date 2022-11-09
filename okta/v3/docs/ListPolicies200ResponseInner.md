# ListPolicies200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**PolicyType**](PolicyType.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**AccountLink** | Pointer to [**PolicyAccountLink**](PolicyAccountLink.md) |  | [optional] 
**Conditions** | Pointer to [**PolicyRuleConditions**](PolicyRuleConditions.md) |  | [optional] 
**MaxClockSkew** | Pointer to **int32** |  | [optional] 
**Provisioning** | Pointer to [**Provisioning**](Provisioning.md) |  | [optional] 
**Subject** | Pointer to [**PolicySubject**](PolicySubject.md) |  | [optional] 
**Settings** | Pointer to [**PasswordPolicySettings**](PasswordPolicySettings.md) |  | [optional] 

## Methods

### NewListPolicies200ResponseInner

`func NewListPolicies200ResponseInner() *ListPolicies200ResponseInner`

NewListPolicies200ResponseInner instantiates a new ListPolicies200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicies200ResponseInnerWithDefaults

`func NewListPolicies200ResponseInnerWithDefaults() *ListPolicies200ResponseInner`

NewListPolicies200ResponseInnerWithDefaults instantiates a new ListPolicies200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListPolicies200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListPolicies200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListPolicies200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListPolicies200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ListPolicies200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListPolicies200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListPolicies200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListPolicies200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ListPolicies200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPolicies200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPolicies200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListPolicies200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListPolicies200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListPolicies200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListPolicies200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListPolicies200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ListPolicies200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPolicies200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPolicies200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPolicies200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *ListPolicies200ResponseInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListPolicies200ResponseInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListPolicies200ResponseInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ListPolicies200ResponseInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *ListPolicies200ResponseInner) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPolicies200ResponseInner) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPolicies200ResponseInner) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListPolicies200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *ListPolicies200ResponseInner) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ListPolicies200ResponseInner) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ListPolicies200ResponseInner) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ListPolicies200ResponseInner) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *ListPolicies200ResponseInner) GetType() PolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListPolicies200ResponseInner) GetTypeOk() (*PolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListPolicies200ResponseInner) SetType(v PolicyType)`

SetType sets Type field to given value.

### HasType

`func (o *ListPolicies200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *ListPolicies200ResponseInner) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListPolicies200ResponseInner) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListPolicies200ResponseInner) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ListPolicies200ResponseInner) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ListPolicies200ResponseInner) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListPolicies200ResponseInner) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListPolicies200ResponseInner) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListPolicies200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccountLink

`func (o *ListPolicies200ResponseInner) GetAccountLink() PolicyAccountLink`

GetAccountLink returns the AccountLink field if non-nil, zero value otherwise.

### GetAccountLinkOk

`func (o *ListPolicies200ResponseInner) GetAccountLinkOk() (*PolicyAccountLink, bool)`

GetAccountLinkOk returns a tuple with the AccountLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLink

`func (o *ListPolicies200ResponseInner) SetAccountLink(v PolicyAccountLink)`

SetAccountLink sets AccountLink field to given value.

### HasAccountLink

`func (o *ListPolicies200ResponseInner) HasAccountLink() bool`

HasAccountLink returns a boolean if a field has been set.

### GetConditions

`func (o *ListPolicies200ResponseInner) GetConditions() PolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ListPolicies200ResponseInner) GetConditionsOk() (*PolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ListPolicies200ResponseInner) SetConditions(v PolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ListPolicies200ResponseInner) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetMaxClockSkew

`func (o *ListPolicies200ResponseInner) GetMaxClockSkew() int32`

GetMaxClockSkew returns the MaxClockSkew field if non-nil, zero value otherwise.

### GetMaxClockSkewOk

`func (o *ListPolicies200ResponseInner) GetMaxClockSkewOk() (*int32, bool)`

GetMaxClockSkewOk returns a tuple with the MaxClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClockSkew

`func (o *ListPolicies200ResponseInner) SetMaxClockSkew(v int32)`

SetMaxClockSkew sets MaxClockSkew field to given value.

### HasMaxClockSkew

`func (o *ListPolicies200ResponseInner) HasMaxClockSkew() bool`

HasMaxClockSkew returns a boolean if a field has been set.

### GetProvisioning

`func (o *ListPolicies200ResponseInner) GetProvisioning() Provisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *ListPolicies200ResponseInner) GetProvisioningOk() (*Provisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *ListPolicies200ResponseInner) SetProvisioning(v Provisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *ListPolicies200ResponseInner) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetSubject

`func (o *ListPolicies200ResponseInner) GetSubject() PolicySubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ListPolicies200ResponseInner) GetSubjectOk() (*PolicySubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ListPolicies200ResponseInner) SetSubject(v PolicySubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ListPolicies200ResponseInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSettings

`func (o *ListPolicies200ResponseInner) GetSettings() PasswordPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ListPolicies200ResponseInner) GetSettingsOk() (*PasswordPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ListPolicies200ResponseInner) SetSettings(v PasswordPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ListPolicies200ResponseInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


