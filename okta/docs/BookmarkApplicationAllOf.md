# BookmarkApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ApplicationCredentials**](ApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [default to "bookmark"]
**Settings** | Pointer to [**BookmarkApplicationSettings**](BookmarkApplicationSettings.md) |  | [optional] 

## Methods

### NewBookmarkApplicationAllOf

`func NewBookmarkApplicationAllOf() *BookmarkApplicationAllOf`

NewBookmarkApplicationAllOf instantiates a new BookmarkApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkApplicationAllOfWithDefaults

`func NewBookmarkApplicationAllOfWithDefaults() *BookmarkApplicationAllOf`

NewBookmarkApplicationAllOfWithDefaults instantiates a new BookmarkApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *BookmarkApplicationAllOf) GetCredentials() ApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BookmarkApplicationAllOf) GetCredentialsOk() (*ApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BookmarkApplicationAllOf) SetCredentials(v ApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BookmarkApplicationAllOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *BookmarkApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BookmarkApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BookmarkApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BookmarkApplicationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *BookmarkApplicationAllOf) GetSettings() BookmarkApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BookmarkApplicationAllOf) GetSettingsOk() (*BookmarkApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BookmarkApplicationAllOf) SetSettings(v BookmarkApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *BookmarkApplicationAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


