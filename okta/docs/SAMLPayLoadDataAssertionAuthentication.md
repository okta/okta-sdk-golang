# SAMLPayLoadDataAssertionAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionIndex** | Pointer to **string** | The unique identifier describing the assertion statement | [optional] 
**AuthnContext** | Pointer to [**SAMLPayLoadDataAssertionAuthenticationAuthnContext**](SAMLPayLoadDataAssertionAuthenticationAuthnContext.md) |  | [optional] 

## Methods

### NewSAMLPayLoadDataAssertionAuthentication

`func NewSAMLPayLoadDataAssertionAuthentication() *SAMLPayLoadDataAssertionAuthentication`

NewSAMLPayLoadDataAssertionAuthentication instantiates a new SAMLPayLoadDataAssertionAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadDataAssertionAuthenticationWithDefaults

`func NewSAMLPayLoadDataAssertionAuthenticationWithDefaults() *SAMLPayLoadDataAssertionAuthentication`

NewSAMLPayLoadDataAssertionAuthenticationWithDefaults instantiates a new SAMLPayLoadDataAssertionAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionIndex

`func (o *SAMLPayLoadDataAssertionAuthentication) GetSessionIndex() string`

GetSessionIndex returns the SessionIndex field if non-nil, zero value otherwise.

### GetSessionIndexOk

`func (o *SAMLPayLoadDataAssertionAuthentication) GetSessionIndexOk() (*string, bool)`

GetSessionIndexOk returns a tuple with the SessionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIndex

`func (o *SAMLPayLoadDataAssertionAuthentication) SetSessionIndex(v string)`

SetSessionIndex sets SessionIndex field to given value.

### HasSessionIndex

`func (o *SAMLPayLoadDataAssertionAuthentication) HasSessionIndex() bool`

HasSessionIndex returns a boolean if a field has been set.

### GetAuthnContext

`func (o *SAMLPayLoadDataAssertionAuthentication) GetAuthnContext() SAMLPayLoadDataAssertionAuthenticationAuthnContext`

GetAuthnContext returns the AuthnContext field if non-nil, zero value otherwise.

### GetAuthnContextOk

`func (o *SAMLPayLoadDataAssertionAuthentication) GetAuthnContextOk() (*SAMLPayLoadDataAssertionAuthenticationAuthnContext, bool)`

GetAuthnContextOk returns a tuple with the AuthnContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContext

`func (o *SAMLPayLoadDataAssertionAuthentication) SetAuthnContext(v SAMLPayLoadDataAssertionAuthenticationAuthnContext)`

SetAuthnContext sets AuthnContext field to given value.

### HasAuthnContext

`func (o *SAMLPayLoadDataAssertionAuthentication) HasAuthnContext() bool`

HasAuthnContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


