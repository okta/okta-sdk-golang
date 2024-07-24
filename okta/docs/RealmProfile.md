# RealmProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of a Realm | 
**RealmType** | Pointer to **string** | Used to store partner users. This must be set to Partner to access Okta&#39;s external partner portal. | [optional] 

## Methods

### NewRealmProfile

`func NewRealmProfile(name string, ) *RealmProfile`

NewRealmProfile instantiates a new RealmProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmProfileWithDefaults

`func NewRealmProfileWithDefaults() *RealmProfile`

NewRealmProfileWithDefaults instantiates a new RealmProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RealmProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealmProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealmProfile) SetName(v string)`

SetName sets Name field to given value.


### GetRealmType

`func (o *RealmProfile) GetRealmType() string`

GetRealmType returns the RealmType field if non-nil, zero value otherwise.

### GetRealmTypeOk

`func (o *RealmProfile) GetRealmTypeOk() (*string, bool)`

GetRealmTypeOk returns a tuple with the RealmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmType

`func (o *RealmProfile) SetRealmType(v string)`

SetRealmType sets RealmType field to given value.

### HasRealmType

`func (o *RealmProfile) HasRealmType() bool`

HasRealmType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


