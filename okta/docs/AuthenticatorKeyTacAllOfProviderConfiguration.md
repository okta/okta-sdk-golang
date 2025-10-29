# AuthenticatorKeyTacAllOfProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinTtl** | **float32** | Minimum time-to-live (TTL) of the TAC in minutes. The &#x60;minTtl&#x60; indicates the minimum amount of time that a TAC is valid. The &#x60;minTtl&#x60; must be less than the &#x60;maxTtl&#x60;. | 
**MaxTtl** | **float32** | Maximum TTL of the TAC in minutes. The &#x60;maxTtl&#x60; indicates the maximum amount of time that a TAC is valid. The &#x60;maxTtl&#x60; must be greater than the &#x60;minTtl&#x60;. | 
**DefaultTtl** | **float32** | The default TTL in minutes when you create a TAC. The &#x60;defaultTtl&#x60; indicates the actual amount of time that a TAC is valid before it expires. The &#x60;defaultTtl&#x60; must be greater than the &#x60;minTtl&#x60; and less than the &#x60;maxTtl&#x60;. | [default to 120]
**Length** | **float32** | Defines the number of characters in a TAC. For example, a &#x60;length&#x60; of &#x60;16&#x60; means that the TAC is 16 characters. | 
**Complexity** | [**AuthenticatorKeyTacAllOfProviderConfigurationComplexity**](AuthenticatorKeyTacAllOfProviderConfigurationComplexity.md) |  | 
**MultiUseAllowed** | Pointer to **bool** | Indicates whether a TAC can be used multiple times. If set to &#x60;true&#x60;, the TAC can be used multiple times until it expires. | [optional] 

## Methods

### NewAuthenticatorKeyTacAllOfProviderConfiguration

`func NewAuthenticatorKeyTacAllOfProviderConfiguration(minTtl float32, maxTtl float32, defaultTtl float32, length float32, complexity AuthenticatorKeyTacAllOfProviderConfigurationComplexity, ) *AuthenticatorKeyTacAllOfProviderConfiguration`

NewAuthenticatorKeyTacAllOfProviderConfiguration instantiates a new AuthenticatorKeyTacAllOfProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorKeyTacAllOfProviderConfigurationWithDefaults

`func NewAuthenticatorKeyTacAllOfProviderConfigurationWithDefaults() *AuthenticatorKeyTacAllOfProviderConfiguration`

NewAuthenticatorKeyTacAllOfProviderConfigurationWithDefaults instantiates a new AuthenticatorKeyTacAllOfProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinTtl

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMinTtl() float32`

GetMinTtl returns the MinTtl field if non-nil, zero value otherwise.

### GetMinTtlOk

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMinTtlOk() (*float32, bool)`

GetMinTtlOk returns a tuple with the MinTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTtl

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetMinTtl(v float32)`

SetMinTtl sets MinTtl field to given value.


### GetMaxTtl

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMaxTtl() float32`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMaxTtlOk() (*float32, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetMaxTtl(v float32)`

SetMaxTtl sets MaxTtl field to given value.


### GetDefaultTtl

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetDefaultTtl() float32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetDefaultTtlOk() (*float32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetDefaultTtl(v float32)`

SetDefaultTtl sets DefaultTtl field to given value.


### GetLength

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetLength(v float32)`

SetLength sets Length field to given value.


### GetComplexity

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetComplexity() AuthenticatorKeyTacAllOfProviderConfigurationComplexity`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetComplexityOk() (*AuthenticatorKeyTacAllOfProviderConfigurationComplexity, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetComplexity(v AuthenticatorKeyTacAllOfProviderConfigurationComplexity)`

SetComplexity sets Complexity field to given value.


### GetMultiUseAllowed

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMultiUseAllowed() bool`

GetMultiUseAllowed returns the MultiUseAllowed field if non-nil, zero value otherwise.

### GetMultiUseAllowedOk

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) GetMultiUseAllowedOk() (*bool, bool)`

GetMultiUseAllowedOk returns a tuple with the MultiUseAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiUseAllowed

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) SetMultiUseAllowed(v bool)`

SetMultiUseAllowed sets MultiUseAllowed field to given value.

### HasMultiUseAllowed

`func (o *AuthenticatorKeyTacAllOfProviderConfiguration) HasMultiUseAllowed() bool`

HasMultiUseAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


