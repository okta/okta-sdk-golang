# AuthenticationProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AuthenticationProviderType**](AuthenticationProviderType.md) |  | [optional] 

## Methods

### NewAuthenticationProvider

`func NewAuthenticationProvider() *AuthenticationProvider`

NewAuthenticationProvider instantiates a new AuthenticationProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProviderWithDefaults

`func NewAuthenticationProviderWithDefaults() *AuthenticationProvider`

NewAuthenticationProviderWithDefaults instantiates a new AuthenticationProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticationProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AuthenticationProvider) GetType() AuthenticationProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticationProvider) GetTypeOk() (*AuthenticationProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticationProvider) SetType(v AuthenticationProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *AuthenticationProvider) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


