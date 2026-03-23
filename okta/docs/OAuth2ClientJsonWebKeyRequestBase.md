# OAuth2ClientJsonWebKeyRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]

## Methods

### NewOAuth2ClientJsonWebKeyRequestBase

`func NewOAuth2ClientJsonWebKeyRequestBase() *OAuth2ClientJsonWebKeyRequestBase`

NewOAuth2ClientJsonWebKeyRequestBase instantiates a new OAuth2ClientJsonWebKeyRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults

`func NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults() *OAuth2ClientJsonWebKeyRequestBase`

NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults instantiates a new OAuth2ClientJsonWebKeyRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKid

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ClientJsonWebKeyRequestBase) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ClientJsonWebKeyRequestBase) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientJsonWebKeyRequestBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


