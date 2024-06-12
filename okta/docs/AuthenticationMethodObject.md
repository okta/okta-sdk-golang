# AuthenticationMethodObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A label that identifies the authenticator | [optional] 
**Method** | Pointer to **string** | Specifies the method used for the authenticator | [optional] 

## Methods

### NewAuthenticationMethodObject

`func NewAuthenticationMethodObject() *AuthenticationMethodObject`

NewAuthenticationMethodObject instantiates a new AuthenticationMethodObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodObjectWithDefaults

`func NewAuthenticationMethodObjectWithDefaults() *AuthenticationMethodObject`

NewAuthenticationMethodObjectWithDefaults instantiates a new AuthenticationMethodObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AuthenticationMethodObject) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthenticationMethodObject) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthenticationMethodObject) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuthenticationMethodObject) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMethod

`func (o *AuthenticationMethodObject) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthenticationMethodObject) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthenticationMethodObject) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuthenticationMethodObject) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


