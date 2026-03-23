# UnconfirmedUserResponseSchemaUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The ID of the user in the target app that&#39;s linked to the Okta application user object. This value is the native app-specific identifier or primary key for the user in the target app.  The &#x60;externalId&#x60; is set during import when the user is confirmed (reconciled) or during provisioning when the user is created in the target app. This value isn&#39;t populated for SSO app assignments (for example, SAML or SWA) because it isn&#39;t synchronized with a target app. | [optional] [readonly] 
**Credentials** | Pointer to [**AppUserCredentials**](AppUserCredentials.md) |  | [optional] 
**Profile** | Pointer to **map[string]interface{}** | Specifies the default and custom profile properties for a user. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can&#39;t be configured. See [profile](/openapi/okta-management/management/user/getuser#user/getuser/t&#x3D;response&amp;c&#x3D;200&amp;path&#x3D;profile).  | [optional] 

## Methods

### NewUnconfirmedUserResponseSchemaUsersInner

`func NewUnconfirmedUserResponseSchemaUsersInner() *UnconfirmedUserResponseSchemaUsersInner`

NewUnconfirmedUserResponseSchemaUsersInner instantiates a new UnconfirmedUserResponseSchemaUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnconfirmedUserResponseSchemaUsersInnerWithDefaults

`func NewUnconfirmedUserResponseSchemaUsersInnerWithDefaults() *UnconfirmedUserResponseSchemaUsersInner`

NewUnconfirmedUserResponseSchemaUsersInnerWithDefaults instantiates a new UnconfirmedUserResponseSchemaUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *UnconfirmedUserResponseSchemaUsersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UnconfirmedUserResponseSchemaUsersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UnconfirmedUserResponseSchemaUsersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UnconfirmedUserResponseSchemaUsersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCredentials

`func (o *UnconfirmedUserResponseSchemaUsersInner) GetCredentials() AppUserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UnconfirmedUserResponseSchemaUsersInner) GetCredentialsOk() (*AppUserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UnconfirmedUserResponseSchemaUsersInner) SetCredentials(v AppUserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UnconfirmedUserResponseSchemaUsersInner) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetProfile

`func (o *UnconfirmedUserResponseSchemaUsersInner) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UnconfirmedUserResponseSchemaUsersInner) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UnconfirmedUserResponseSchemaUsersInner) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UnconfirmedUserResponseSchemaUsersInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


