# Realm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Realm was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the Realm | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Conveys whether the Realm is the default | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Realm was last updated | [optional] [readonly] 
**Profile** | Pointer to [**RealmProfile**](RealmProfile.md) |  | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewRealm

`func NewRealm() *Realm`

NewRealm instantiates a new Realm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmWithDefaults

`func NewRealmWithDefaults() *Realm`

NewRealmWithDefaults instantiates a new Realm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Realm) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Realm) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Realm) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Realm) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Realm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Realm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Realm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Realm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *Realm) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Realm) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Realm) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Realm) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Realm) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Realm) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Realm) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Realm) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *Realm) GetProfile() RealmProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Realm) GetProfileOk() (*RealmProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Realm) SetProfile(v RealmProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Realm) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetLinks

`func (o *Realm) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Realm) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Realm) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Realm) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


