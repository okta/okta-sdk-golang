# OAuth2SettingsPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kty** | Pointer to **string** | Key type (e.g., &#x60;RSA&#x60;, &#x60;EC&#x60;) | [optional] 
**Kid** | Pointer to **string** | Key ID | [optional] 
**E** | Pointer to **string** | RSA public exponent | [optional] 
**N** | Pointer to **string** | RSA modulus | [optional] 

## Methods

### NewOAuth2SettingsPublicKey

`func NewOAuth2SettingsPublicKey() *OAuth2SettingsPublicKey`

NewOAuth2SettingsPublicKey instantiates a new OAuth2SettingsPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2SettingsPublicKeyWithDefaults

`func NewOAuth2SettingsPublicKeyWithDefaults() *OAuth2SettingsPublicKey`

NewOAuth2SettingsPublicKeyWithDefaults instantiates a new OAuth2SettingsPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKty

`func (o *OAuth2SettingsPublicKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OAuth2SettingsPublicKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OAuth2SettingsPublicKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OAuth2SettingsPublicKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2SettingsPublicKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2SettingsPublicKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2SettingsPublicKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2SettingsPublicKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetE

`func (o *OAuth2SettingsPublicKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OAuth2SettingsPublicKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OAuth2SettingsPublicKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *OAuth2SettingsPublicKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetN

`func (o *OAuth2SettingsPublicKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OAuth2SettingsPublicKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OAuth2SettingsPublicKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *OAuth2SettingsPublicKey) HasN() bool`

HasN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


