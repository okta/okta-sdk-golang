# OAuth2ClientJsonWebKeySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]AddJwk201Response**](AddJwk201Response.md) | The JSON Web Keys in this key set | [optional] 

## Methods

### NewOAuth2ClientJsonWebKeySet

`func NewOAuth2ClientJsonWebKeySet() *OAuth2ClientJsonWebKeySet`

NewOAuth2ClientJsonWebKeySet instantiates a new OAuth2ClientJsonWebKeySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeySetWithDefaults

`func NewOAuth2ClientJsonWebKeySetWithDefaults() *OAuth2ClientJsonWebKeySet`

NewOAuth2ClientJsonWebKeySetWithDefaults instantiates a new OAuth2ClientJsonWebKeySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *OAuth2ClientJsonWebKeySet) GetKeys() []AddJwk201Response`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *OAuth2ClientJsonWebKeySet) GetKeysOk() (*[]AddJwk201Response, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *OAuth2ClientJsonWebKeySet) SetKeys(v []AddJwk201Response)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *OAuth2ClientJsonWebKeySet) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


