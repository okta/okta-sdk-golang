# AuthenticatorProfileTacRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiUse** | Pointer to **bool** | Determines whether the enrollment can be used more than once. To enable multi-use, the org-level authenticator’s configuration must allow multi-use. | [optional] 
**Ttl** | Pointer to **string** | Time-to-live (TTL) in minutes.  Specifies how long the TAC enrollment is valid after it&#39;s created and activated. The configured value must be between 10 minutes (&#x60;10&#x60;) and 10 days (&#x60;14400&#x60;), inclusive. The actual allowed range depends on the org-level authenticator configuration. | [optional] 

## Methods

### NewAuthenticatorProfileTacRequest

`func NewAuthenticatorProfileTacRequest() *AuthenticatorProfileTacRequest`

NewAuthenticatorProfileTacRequest instantiates a new AuthenticatorProfileTacRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorProfileTacRequestWithDefaults

`func NewAuthenticatorProfileTacRequestWithDefaults() *AuthenticatorProfileTacRequest`

NewAuthenticatorProfileTacRequestWithDefaults instantiates a new AuthenticatorProfileTacRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultiUse

`func (o *AuthenticatorProfileTacRequest) GetMultiUse() bool`

GetMultiUse returns the MultiUse field if non-nil, zero value otherwise.

### GetMultiUseOk

`func (o *AuthenticatorProfileTacRequest) GetMultiUseOk() (*bool, bool)`

GetMultiUseOk returns a tuple with the MultiUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiUse

`func (o *AuthenticatorProfileTacRequest) SetMultiUse(v bool)`

SetMultiUse sets MultiUse field to given value.

### HasMultiUse

`func (o *AuthenticatorProfileTacRequest) HasMultiUse() bool`

HasMultiUse returns a boolean if a field has been set.

### GetTtl

`func (o *AuthenticatorProfileTacRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *AuthenticatorProfileTacRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *AuthenticatorProfileTacRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *AuthenticatorProfileTacRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


