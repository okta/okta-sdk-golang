# WebAuthnCredResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorEnrollmentId** | Pointer to **string** | ID for a WebAuthn preregistration factor in Okta | [optional] 
**CredResponseJwe** | Pointer to **string** | Encrypted JSON Web Encryption (JWE) of the credential response from the fulfillment provider | [optional] 

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

### GetCredResponseJwe

`func (o *WebAuthnCredResponse) GetCredResponseJwe() string`

GetCredResponseJwe returns the CredResponseJwe field if non-nil, zero value otherwise.

### GetCredResponseJweOk

`func (o *WebAuthnCredResponse) GetCredResponseJweOk() (*string, bool)`

GetCredResponseJweOk returns a tuple with the CredResponseJwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredResponseJwe

`func (o *WebAuthnCredResponse) SetCredResponseJwe(v string)`

SetCredResponseJwe sets CredResponseJwe field to given value.

### HasCredResponseJwe

`func (o *WebAuthnCredResponse) HasCredResponseJwe() bool`

HasCredResponseJwe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


