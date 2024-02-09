# BookmarkApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**ApplicationCredentials**](ApplicationCredentials.md) |  | [optional] 
**Name** | Pointer to **string** | Unique key for the app definition | [optional] [default to "bookmark"]
**Settings** | Pointer to [**BookmarkApplicationSettings**](BookmarkApplicationSettings.md) |  | [optional] 

## Methods

### NewBookmarkApplication

`func NewBookmarkApplication() *BookmarkApplication`

NewBookmarkApplication instantiates a new BookmarkApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkApplicationWithDefaults

`func NewBookmarkApplicationWithDefaults() *BookmarkApplication`

NewBookmarkApplicationWithDefaults instantiates a new BookmarkApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *BookmarkApplication) GetCredentials() ApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BookmarkApplication) GetCredentialsOk() (*ApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BookmarkApplication) SetCredentials(v ApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BookmarkApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *BookmarkApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BookmarkApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BookmarkApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BookmarkApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *BookmarkApplication) GetSettings() BookmarkApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *BookmarkApplication) GetSettingsOk() (*BookmarkApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *BookmarkApplication) SetSettings(v BookmarkApplicationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *BookmarkApplication) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


