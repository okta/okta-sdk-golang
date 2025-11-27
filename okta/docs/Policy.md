# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the policy was created | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the policy | [optional] 
**Id** | Pointer to **string** | Identifier of the policy | [optional] [readonly] [default to "Assigned"]
**LastUpdated** | Pointer to **time.Time** | Timestamp when the policy was last modified | [optional] [readonly] 
**Name** | **string** | Name of the policy | 
**Priority** | Pointer to **int32** | Specifies the order in which this policy is evaluated in relation to the other policies | [optional] 
**Status** | Pointer to **string** | Whether or not the policy is active. Use the &#x60;activate&#x60; query parameter to set the status of a policy. | [optional] 
**System** | Pointer to **bool** | Specifies whether Okta created the policy | [optional] [default to false]
**Type** | **string** | All Okta orgs contain only one IdP discovery policy with an immutable default rule routing to your org&#39;s sign-in page, one entity risk policy, and one session protection policy. Creating or replacing a policy with the &#x60;IDP_DISCOVERY&#x60; type, the &#x60;ENTITY_RISK&#x60; type, or the &#x60;POST_AUTH_SESSION&#x60; type isn&#39;t supported. The following policy types are available with Identity Engine: &#x60;ACCESS_POLICY&#x60;, &#x60;PROFILE_ENROLLMENT&#x60;, &#x60;POST_AUTH_SESSION&#x60;, &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &#x60;DEVICE_SIGNAL_COLLECTION&#x60;, &#x60;ENTITY_RISK&#x60;. | 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**PolicyLinks**](PolicyLinks.md) |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy(name string, type_ string, ) *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Policy) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Policy) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Policy) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Policy) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Policy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Policy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Policy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Policy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *Policy) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Policy) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Policy) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Policy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *Policy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Policy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Policy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Policy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *Policy) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Policy) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Policy) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Policy) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *Policy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Policy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Policy) SetType(v string)`

SetType sets Type field to given value.


### GetEmbedded

`func (o *Policy) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Policy) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Policy) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *Policy) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *Policy) GetLinks() PolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Policy) GetLinksOk() (*PolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Policy) SetLinks(v PolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Policy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


