# RealmAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**Actions**](Actions.md) |  | [optional] 
**Conditions** | Pointer to [**Conditions**](Conditions.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the realm assignment was created | [optional] [readonly] 
**Domains** | Pointer to **[]string** | Array of allowed domains. No user in this realm can be created or updated unless they have a username and email from one of these domains.  The following characters aren&#39;t allowed in the domain name: &#x60;!$%^&amp;()&#x3D;*+,:;&lt;&gt;&#39;[]|/?\\&#x60; | [optional] 
**Id** | Pointer to **string** | Unique ID of the realm assignment | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Indicates the default realm. Existing users will start out in the default realm and can be moved individually to other realms. | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp of when the realm assignment was updated | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the realm | [optional] 
**Priority** | Pointer to **int32** | The priority of the realm assignment. The lower the number, the higher the priority. This helps resolve conflicts between realm assignments. &gt; **Note:** When you create realm assignments in bulk, realm assignment priorities must be unique. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewRealmAssignment

`func NewRealmAssignment() *RealmAssignment`

NewRealmAssignment instantiates a new RealmAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmAssignmentWithDefaults

`func NewRealmAssignmentWithDefaults() *RealmAssignment`

NewRealmAssignmentWithDefaults instantiates a new RealmAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *RealmAssignment) GetActions() Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RealmAssignment) GetActionsOk() (*Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RealmAssignment) SetActions(v Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RealmAssignment) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *RealmAssignment) GetConditions() Conditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RealmAssignment) GetConditionsOk() (*Conditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RealmAssignment) SetConditions(v Conditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *RealmAssignment) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreated

`func (o *RealmAssignment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RealmAssignment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RealmAssignment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RealmAssignment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDomains

`func (o *RealmAssignment) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *RealmAssignment) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *RealmAssignment) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *RealmAssignment) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetId

`func (o *RealmAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RealmAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *RealmAssignment) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RealmAssignment) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RealmAssignment) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *RealmAssignment) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RealmAssignment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RealmAssignment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RealmAssignment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RealmAssignment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *RealmAssignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealmAssignment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealmAssignment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RealmAssignment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *RealmAssignment) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RealmAssignment) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RealmAssignment) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RealmAssignment) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *RealmAssignment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RealmAssignment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RealmAssignment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RealmAssignment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *RealmAssignment) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RealmAssignment) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RealmAssignment) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RealmAssignment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


