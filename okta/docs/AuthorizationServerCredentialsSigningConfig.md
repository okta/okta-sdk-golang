# AuthorizationServerCredentialsSigningConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kid** | Pointer to **string** |  | [optional] 
**LastRotated** | Pointer to **time.Time** |  | [optional] [readonly] 
**NextRotation** | Pointer to **time.Time** |  | [optional] [readonly] 
**RotationMode** | Pointer to [**AuthorizationServerCredentialsRotationMode**](AuthorizationServerCredentialsRotationMode.md) |  | [optional] 
**Use** | Pointer to [**AuthorizationServerCredentialsUse**](AuthorizationServerCredentialsUse.md) |  | [optional] 

## Methods

### NewAuthorizationServerCredentialsSigningConfig

`func NewAuthorizationServerCredentialsSigningConfig() *AuthorizationServerCredentialsSigningConfig`

NewAuthorizationServerCredentialsSigningConfig instantiates a new AuthorizationServerCredentialsSigningConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerCredentialsSigningConfigWithDefaults

`func NewAuthorizationServerCredentialsSigningConfigWithDefaults() *AuthorizationServerCredentialsSigningConfig`

NewAuthorizationServerCredentialsSigningConfigWithDefaults instantiates a new AuthorizationServerCredentialsSigningConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKid

`func (o *AuthorizationServerCredentialsSigningConfig) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AuthorizationServerCredentialsSigningConfig) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AuthorizationServerCredentialsSigningConfig) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AuthorizationServerCredentialsSigningConfig) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetLastRotated

`func (o *AuthorizationServerCredentialsSigningConfig) GetLastRotated() time.Time`

GetLastRotated returns the LastRotated field if non-nil, zero value otherwise.

### GetLastRotatedOk

`func (o *AuthorizationServerCredentialsSigningConfig) GetLastRotatedOk() (*time.Time, bool)`

GetLastRotatedOk returns a tuple with the LastRotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotated

`func (o *AuthorizationServerCredentialsSigningConfig) SetLastRotated(v time.Time)`

SetLastRotated sets LastRotated field to given value.

### HasLastRotated

`func (o *AuthorizationServerCredentialsSigningConfig) HasLastRotated() bool`

HasLastRotated returns a boolean if a field has been set.

### GetNextRotation

`func (o *AuthorizationServerCredentialsSigningConfig) GetNextRotation() time.Time`

GetNextRotation returns the NextRotation field if non-nil, zero value otherwise.

### GetNextRotationOk

`func (o *AuthorizationServerCredentialsSigningConfig) GetNextRotationOk() (*time.Time, bool)`

GetNextRotationOk returns a tuple with the NextRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRotation

`func (o *AuthorizationServerCredentialsSigningConfig) SetNextRotation(v time.Time)`

SetNextRotation sets NextRotation field to given value.

### HasNextRotation

`func (o *AuthorizationServerCredentialsSigningConfig) HasNextRotation() bool`

HasNextRotation returns a boolean if a field has been set.

### GetRotationMode

`func (o *AuthorizationServerCredentialsSigningConfig) GetRotationMode() AuthorizationServerCredentialsRotationMode`

GetRotationMode returns the RotationMode field if non-nil, zero value otherwise.

### GetRotationModeOk

`func (o *AuthorizationServerCredentialsSigningConfig) GetRotationModeOk() (*AuthorizationServerCredentialsRotationMode, bool)`

GetRotationModeOk returns a tuple with the RotationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationMode

`func (o *AuthorizationServerCredentialsSigningConfig) SetRotationMode(v AuthorizationServerCredentialsRotationMode)`

SetRotationMode sets RotationMode field to given value.

### HasRotationMode

`func (o *AuthorizationServerCredentialsSigningConfig) HasRotationMode() bool`

HasRotationMode returns a boolean if a field has been set.

### GetUse

`func (o *AuthorizationServerCredentialsSigningConfig) GetUse() AuthorizationServerCredentialsUse`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AuthorizationServerCredentialsSigningConfig) GetUseOk() (*AuthorizationServerCredentialsUse, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AuthorizationServerCredentialsSigningConfig) SetUse(v AuthorizationServerCredentialsUse)`

SetUse sets Use field to given value.

### HasUse

`func (o *AuthorizationServerCredentialsSigningConfig) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


