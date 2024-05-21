# SessionIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identity Provider ID. If the &#x60;type&#x60; is &#x60;OKTA&#x60;, then the &#x60;id&#x60; is the org ID. | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionIdentityProvider

`func NewSessionIdentityProvider() *SessionIdentityProvider`

NewSessionIdentityProvider instantiates a new SessionIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionIdentityProviderWithDefaults

`func NewSessionIdentityProviderWithDefaults() *SessionIdentityProvider`

NewSessionIdentityProviderWithDefaults instantiates a new SessionIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionIdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionIdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionIdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionIdentityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SessionIdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionIdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionIdentityProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SessionIdentityProvider) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


