# SupportedMethodsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyProtection** | Pointer to **string** | Indicates whether you must use a hardware key store | [optional] 
**Algorithms** | Pointer to **[]string** | The encryption algorithm for this authenticator method | [optional] 
**TransactionTypes** | Pointer to **[]string** | The transaction type for this authenticator method | [optional] 

## Methods

### NewSupportedMethodsSettings

`func NewSupportedMethodsSettings() *SupportedMethodsSettings`

NewSupportedMethodsSettings instantiates a new SupportedMethodsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedMethodsSettingsWithDefaults

`func NewSupportedMethodsSettingsWithDefaults() *SupportedMethodsSettings`

NewSupportedMethodsSettingsWithDefaults instantiates a new SupportedMethodsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyProtection

`func (o *SupportedMethodsSettings) GetKeyProtection() string`

GetKeyProtection returns the KeyProtection field if non-nil, zero value otherwise.

### GetKeyProtectionOk

`func (o *SupportedMethodsSettings) GetKeyProtectionOk() (*string, bool)`

GetKeyProtectionOk returns a tuple with the KeyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProtection

`func (o *SupportedMethodsSettings) SetKeyProtection(v string)`

SetKeyProtection sets KeyProtection field to given value.

### HasKeyProtection

`func (o *SupportedMethodsSettings) HasKeyProtection() bool`

HasKeyProtection returns a boolean if a field has been set.

### GetAlgorithms

`func (o *SupportedMethodsSettings) GetAlgorithms() []string`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *SupportedMethodsSettings) GetAlgorithmsOk() (*[]string, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *SupportedMethodsSettings) SetAlgorithms(v []string)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *SupportedMethodsSettings) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetTransactionTypes

`func (o *SupportedMethodsSettings) GetTransactionTypes() []string`

GetTransactionTypes returns the TransactionTypes field if non-nil, zero value otherwise.

### GetTransactionTypesOk

`func (o *SupportedMethodsSettings) GetTransactionTypesOk() (*[]string, bool)`

GetTransactionTypesOk returns a tuple with the TransactionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTypes

`func (o *SupportedMethodsSettings) SetTransactionTypes(v []string)`

SetTransactionTypes sets TransactionTypes field to given value.

### HasTransactionTypes

`func (o *SupportedMethodsSettings) HasTransactionTypes() bool`

HasTransactionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


