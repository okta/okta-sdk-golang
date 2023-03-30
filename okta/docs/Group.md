# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastMembershipUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**ObjectClass** | Pointer to **[]string** |  | [optional] [readonly] 
**Profile** | Pointer to [**GroupProfile**](GroupProfile.md) |  | [optional] 
**Type** | Pointer to [**GroupType**](GroupType.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**GroupLinks**](GroupLinks.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Group) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Group) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Group) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Group) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastMembershipUpdated

`func (o *Group) GetLastMembershipUpdated() time.Time`

GetLastMembershipUpdated returns the LastMembershipUpdated field if non-nil, zero value otherwise.

### GetLastMembershipUpdatedOk

`func (o *Group) GetLastMembershipUpdatedOk() (*time.Time, bool)`

GetLastMembershipUpdatedOk returns a tuple with the LastMembershipUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMembershipUpdated

`func (o *Group) SetLastMembershipUpdated(v time.Time)`

SetLastMembershipUpdated sets LastMembershipUpdated field to given value.

### HasLastMembershipUpdated

`func (o *Group) HasLastMembershipUpdated() bool`

HasLastMembershipUpdated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Group) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Group) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Group) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Group) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetObjectClass

`func (o *Group) GetObjectClass() []string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *Group) GetObjectClassOk() (*[]string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *Group) SetObjectClass(v []string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *Group) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetProfile

`func (o *Group) GetProfile() GroupProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Group) GetProfileOk() (*GroupProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Group) SetProfile(v GroupProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Group) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetType

`func (o *Group) GetType() GroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Group) GetTypeOk() (*GroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Group) SetType(v GroupType)`

SetType sets Type field to given value.

### HasType

`func (o *Group) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *Group) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Group) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Group) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *Group) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *Group) GetLinks() GroupLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Group) GetLinksOk() (*GroupLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Group) SetLinks(v GroupLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Group) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


