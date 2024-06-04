# CreateRealmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**RealmProfile**](RealmProfile.md) |  | [optional] 

## Methods

### NewCreateRealmRequest

`func NewCreateRealmRequest() *CreateRealmRequest`

NewCreateRealmRequest instantiates a new CreateRealmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRealmRequestWithDefaults

`func NewCreateRealmRequestWithDefaults() *CreateRealmRequest`

NewCreateRealmRequestWithDefaults instantiates a new CreateRealmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *CreateRealmRequest) GetProfile() RealmProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateRealmRequest) GetProfileOk() (*RealmProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateRealmRequest) SetProfile(v RealmProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateRealmRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


