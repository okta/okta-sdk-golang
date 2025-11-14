# PolicyCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the policy was created | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the Policy | [optional] 
**Id** | Pointer to **string** | Identifier of the Policy | [optional] [readonly] [default to "Assigned"]
**LastUpdated** | Pointer to **time.Time** | Timestamp when the policy was last modified | [optional] [readonly] 
**Name** | **string** | Name of the policy | 
**Priority** | Pointer to **int32** | Specifies the order in which this policy is evaluated in relation to the other policies | [optional] 
**Status** | Pointer to **string** | Whether or not the policy is active. Use the &#x60;activate&#x60; query parameter to set the status of a policy. | [optional] 
**System** | Pointer to **bool** | Specifies whether Okta created the Policy | [optional] [default to false]
**Type** | **string** | All Okta orgs contain only one IdP discovery policy with an immutable default rule routing to your org&#39;s sign-in page, one entity risk policy, and one session protection policy. Creating or replacing a policy with the &#x60;IDP_DISCOVERY&#x60; type, the &#x60;ENTITY_RISK&#x60; type, or the &#x60;POST_AUTH_SESSION&#x60; type isn&#39;t supported. The following policy types are available with Identity Engine: &#x60;ACCESS_POLICY&#x60;, &#x60;PROFILE_ENROLLMENT&#x60;, &#x60;POST_AUTH_SESSION&#x60;, &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &#x60;DEVICE_SIGNAL_COLLECTION&#x60;, &#x60;ENTITY_RISK&#x60;. | 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**PolicyLinks**](PolicyLinks.md) |  | [optional] 

## Methods

### NewPolicyCommon

`func NewPolicyCommon(name string, type_ string, ) *PolicyCommon`

NewPolicyCommon instantiates a new PolicyCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCommonWithDefaults

`func NewPolicyCommonWithDefaults() *PolicyCommon`

NewPolicyCommonWithDefaults instantiates a new PolicyCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PolicyCommon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PolicyCommon) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PolicyCommon) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PolicyCommon) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyCommon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyCommon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyCommon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyCommon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PolicyCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PolicyCommon) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PolicyCommon) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PolicyCommon) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PolicyCommon) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *PolicyCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyCommon) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *PolicyCommon) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PolicyCommon) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PolicyCommon) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PolicyCommon) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *PolicyCommon) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyCommon) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyCommon) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PolicyCommon) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *PolicyCommon) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *PolicyCommon) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *PolicyCommon) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *PolicyCommon) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *PolicyCommon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyCommon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyCommon) SetType(v string)`

SetType sets Type field to given value.


### GetEmbedded

`func (o *PolicyCommon) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PolicyCommon) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PolicyCommon) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PolicyCommon) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *PolicyCommon) GetLinks() PolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PolicyCommon) GetLinksOk() (*PolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PolicyCommon) SetLinks(v PolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PolicyCommon) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


