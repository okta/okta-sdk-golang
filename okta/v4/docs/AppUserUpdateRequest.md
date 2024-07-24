# AppUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**AppUserCredentials**](AppUserCredentials.md) |  | [optional] 
**Profile** | Pointer to **map[string]interface{}** | Specifies the default and custom profile properties for a user. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can&#39;t be configured. See [profile](/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c&#x3D;200&amp;path&#x3D;profile&amp;t&#x3D;response).  | [optional] 

## Methods

### NewAppUserUpdateRequest

`func NewAppUserUpdateRequest() *AppUserUpdateRequest`

NewAppUserUpdateRequest instantiates a new AppUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserUpdateRequestWithDefaults

`func NewAppUserUpdateRequestWithDefaults() *AppUserUpdateRequest`

NewAppUserUpdateRequestWithDefaults instantiates a new AppUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *AppUserUpdateRequest) GetCredentials() AppUserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AppUserUpdateRequest) GetCredentialsOk() (*AppUserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AppUserUpdateRequest) SetCredentials(v AppUserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AppUserUpdateRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetProfile

`func (o *AppUserUpdateRequest) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AppUserUpdateRequest) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AppUserUpdateRequest) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AppUserUpdateRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


