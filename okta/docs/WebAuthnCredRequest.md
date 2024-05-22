# WebAuthnCredRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorEnrollmentId** | Pointer to **string** | ID for a WebAuthn Preregistration Factor in Okta | [optional] 
**CredRequestJwe** | Pointer to **string** | Encrypted JWE of credential request for the fulfillment provider | [optional] 
**KeyId** | Pointer to **string** | ID for the Okta response key-pair used to encrypt and decrypt credential requests and responses | [optional] 

## Methods

### NewWebAuthnCredRequest

`func NewWebAuthnCredRequest() *WebAuthnCredRequest`

NewWebAuthnCredRequest instantiates a new WebAuthnCredRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnCredRequestWithDefaults

`func NewWebAuthnCredRequestWithDefaults() *WebAuthnCredRequest`

NewWebAuthnCredRequestWithDefaults instantiates a new WebAuthnCredRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorEnrollmentId

`func (o *WebAuthnCredRequest) GetAuthenticatorEnrollmentId() string`

GetAuthenticatorEnrollmentId returns the AuthenticatorEnrollmentId field if non-nil, zero value otherwise.

### GetAuthenticatorEnrollmentIdOk

`func (o *WebAuthnCredRequest) GetAuthenticatorEnrollmentIdOk() (*string, bool)`

GetAuthenticatorEnrollmentIdOk returns a tuple with the AuthenticatorEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorEnrollmentId

`func (o *WebAuthnCredRequest) SetAuthenticatorEnrollmentId(v string)`

SetAuthenticatorEnrollmentId sets AuthenticatorEnrollmentId field to given value.

### HasAuthenticatorEnrollmentId

`func (o *WebAuthnCredRequest) HasAuthenticatorEnrollmentId() bool`

HasAuthenticatorEnrollmentId returns a boolean if a field has been set.

### GetCredRequestJwe

`func (o *WebAuthnCredRequest) GetCredRequestJwe() string`

GetCredRequestJwe returns the CredRequestJwe field if non-nil, zero value otherwise.

### GetCredRequestJweOk

`func (o *WebAuthnCredRequest) GetCredRequestJweOk() (*string, bool)`

GetCredRequestJweOk returns a tuple with the CredRequestJwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredRequestJwe

`func (o *WebAuthnCredRequest) SetCredRequestJwe(v string)`

SetCredRequestJwe sets CredRequestJwe field to given value.

### HasCredRequestJwe

`func (o *WebAuthnCredRequest) HasCredRequestJwe() bool`

HasCredRequestJwe returns a boolean if a field has been set.

### GetKeyId

`func (o *WebAuthnCredRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *WebAuthnCredRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *WebAuthnCredRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *WebAuthnCredRequest) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


