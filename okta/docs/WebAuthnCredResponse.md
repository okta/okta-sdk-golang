# WebAuthnCredResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorEnrollmentId** | Pointer to **string** | ID for a WebAuthn Preregistration Factor in Okta | [optional] 
**CredResponseJWE** | Pointer to **string** | Encrypted JWE of credential response from the fulfillment provider | [optional] 

## Methods

### NewWebAuthnCredResponse

`func NewWebAuthnCredResponse() *WebAuthnCredResponse`

NewWebAuthnCredResponse instantiates a new WebAuthnCredResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnCredResponseWithDefaults

`func NewWebAuthnCredResponseWithDefaults() *WebAuthnCredResponse`

NewWebAuthnCredResponseWithDefaults instantiates a new WebAuthnCredResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorEnrollmentId

`func (o *WebAuthnCredResponse) GetAuthenticatorEnrollmentId() string`

GetAuthenticatorEnrollmentId returns the AuthenticatorEnrollmentId field if non-nil, zero value otherwise.

### GetAuthenticatorEnrollmentIdOk

`func (o *WebAuthnCredResponse) GetAuthenticatorEnrollmentIdOk() (*string, bool)`

GetAuthenticatorEnrollmentIdOk returns a tuple with the AuthenticatorEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorEnrollmentId

`func (o *WebAuthnCredResponse) SetAuthenticatorEnrollmentId(v string)`

SetAuthenticatorEnrollmentId sets AuthenticatorEnrollmentId field to given value.

### HasAuthenticatorEnrollmentId

`func (o *WebAuthnCredResponse) HasAuthenticatorEnrollmentId() bool`

HasAuthenticatorEnrollmentId returns a boolean if a field has been set.

### GetCredResponseJWE

`func (o *WebAuthnCredResponse) GetCredResponseJWE() string`

GetCredResponseJWE returns the CredResponseJWE field if non-nil, zero value otherwise.

### GetCredResponseJWEOk

`func (o *WebAuthnCredResponse) GetCredResponseJWEOk() (*string, bool)`

GetCredResponseJWEOk returns a tuple with the CredResponseJWE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredResponseJWE

`func (o *WebAuthnCredResponse) SetCredResponseJWE(v string)`

SetCredResponseJWE sets CredResponseJWE field to given value.

### HasCredResponseJWE

`func (o *WebAuthnCredResponse) HasCredResponseJWE() bool`

HasCredResponseJWE returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


