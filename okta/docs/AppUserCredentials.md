# AppUserCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**AppUserPasswordCredential**](AppUserPasswordCredential.md) |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewAppUserCredentials

`func NewAppUserCredentials() *AppUserCredentials`

NewAppUserCredentials instantiates a new AppUserCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserCredentialsWithDefaults

`func NewAppUserCredentialsWithDefaults() *AppUserCredentials`

NewAppUserCredentialsWithDefaults instantiates a new AppUserCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *AppUserCredentials) GetPassword() AppUserPasswordCredential`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AppUserCredentials) GetPasswordOk() (*AppUserPasswordCredential, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AppUserCredentials) SetPassword(v AppUserPasswordCredential)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AppUserCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserName

`func (o *AppUserCredentials) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AppUserCredentials) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AppUserCredentials) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AppUserCredentials) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


