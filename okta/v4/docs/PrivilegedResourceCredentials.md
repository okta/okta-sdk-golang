# PrivilegedResourceCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password associated with the privileged resource | [optional] 
**UserName** | Pointer to **string** | The username associated with the privileged resource | [optional] 

## Methods

### NewPrivilegedResourceCredentials

`func NewPrivilegedResourceCredentials() *PrivilegedResourceCredentials`

NewPrivilegedResourceCredentials instantiates a new PrivilegedResourceCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedResourceCredentialsWithDefaults

`func NewPrivilegedResourceCredentialsWithDefaults() *PrivilegedResourceCredentials`

NewPrivilegedResourceCredentialsWithDefaults instantiates a new PrivilegedResourceCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PrivilegedResourceCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PrivilegedResourceCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PrivilegedResourceCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PrivilegedResourceCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserName

`func (o *PrivilegedResourceCredentials) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PrivilegedResourceCredentials) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PrivilegedResourceCredentials) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *PrivilegedResourceCredentials) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


