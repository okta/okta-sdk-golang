# OAuth2ClientJsonSigningKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Algorithm used in the key | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]

## Methods

### NewOAuth2ClientJsonSigningKeyRequest

`func NewOAuth2ClientJsonSigningKeyRequest() *OAuth2ClientJsonSigningKeyRequest`

NewOAuth2ClientJsonSigningKeyRequest instantiates a new OAuth2ClientJsonSigningKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonSigningKeyRequestWithDefaults

`func NewOAuth2ClientJsonSigningKeyRequestWithDefaults() *OAuth2ClientJsonSigningKeyRequest`

NewOAuth2ClientJsonSigningKeyRequestWithDefaults instantiates a new OAuth2ClientJsonSigningKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *OAuth2ClientJsonSigningKeyRequest) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *OAuth2ClientJsonSigningKeyRequest) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *OAuth2ClientJsonSigningKeyRequest) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *OAuth2ClientJsonSigningKeyRequest) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *OAuth2ClientJsonSigningKeyRequest) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *OAuth2ClientJsonSigningKeyRequest) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *OAuth2ClientJsonSigningKeyRequest) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *OAuth2ClientJsonSigningKeyRequest) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2ClientJsonSigningKeyRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ClientJsonSigningKeyRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ClientJsonSigningKeyRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ClientJsonSigningKeyRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ClientJsonSigningKeyRequest) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ClientJsonSigningKeyRequest) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OAuth2ClientJsonSigningKeyRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientJsonSigningKeyRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientJsonSigningKeyRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientJsonSigningKeyRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


