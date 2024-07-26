# AuthenticatorMethodPushAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithms** | Pointer to **[]string** |  | [optional] 
**KeyProtection** | Pointer to **string** | Indicates whether you must use a hardware key store | [optional] 
**TransactionTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAuthenticatorMethodPushAllOfSettings

`func NewAuthenticatorMethodPushAllOfSettings() *AuthenticatorMethodPushAllOfSettings`

NewAuthenticatorMethodPushAllOfSettings instantiates a new AuthenticatorMethodPushAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodPushAllOfSettingsWithDefaults

`func NewAuthenticatorMethodPushAllOfSettingsWithDefaults() *AuthenticatorMethodPushAllOfSettings`

NewAuthenticatorMethodPushAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodPushAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithms

`func (o *AuthenticatorMethodPushAllOfSettings) GetAlgorithms() []string`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *AuthenticatorMethodPushAllOfSettings) GetAlgorithmsOk() (*[]string, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *AuthenticatorMethodPushAllOfSettings) SetAlgorithms(v []string)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *AuthenticatorMethodPushAllOfSettings) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetKeyProtection

`func (o *AuthenticatorMethodPushAllOfSettings) GetKeyProtection() string`

GetKeyProtection returns the KeyProtection field if non-nil, zero value otherwise.

### GetKeyProtectionOk

`func (o *AuthenticatorMethodPushAllOfSettings) GetKeyProtectionOk() (*string, bool)`

GetKeyProtectionOk returns a tuple with the KeyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProtection

`func (o *AuthenticatorMethodPushAllOfSettings) SetKeyProtection(v string)`

SetKeyProtection sets KeyProtection field to given value.

### HasKeyProtection

`func (o *AuthenticatorMethodPushAllOfSettings) HasKeyProtection() bool`

HasKeyProtection returns a boolean if a field has been set.

### GetTransactionTypes

`func (o *AuthenticatorMethodPushAllOfSettings) GetTransactionTypes() []string`

GetTransactionTypes returns the TransactionTypes field if non-nil, zero value otherwise.

### GetTransactionTypesOk

`func (o *AuthenticatorMethodPushAllOfSettings) GetTransactionTypesOk() (*[]string, bool)`

GetTransactionTypesOk returns a tuple with the TransactionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTypes

`func (o *AuthenticatorMethodPushAllOfSettings) SetTransactionTypes(v []string)`

SetTransactionTypes sets TransactionTypes field to given value.

### HasTransactionTypes

`func (o *AuthenticatorMethodPushAllOfSettings) HasTransactionTypes() bool`

HasTransactionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


