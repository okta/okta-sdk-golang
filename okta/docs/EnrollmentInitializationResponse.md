# EnrollmentInitializationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredRequests** | Pointer to [**[]WebAuthnCredRequest**](WebAuthnCredRequest.md) | List of credential requests for the fulfillment provider | [optional] 
**FulfillmentProvider** | Pointer to **string** | Name of the fulfillment provider for the WebAuthn Preregistration Factor | [optional] 
**PinRequestJwe** | Pointer to **string** | Encrypted JWE of PIN request for the fulfillment provider | [optional] 
**UserId** | Pointer to **string** | ID of an existing Okta user | [optional] 

## Methods

### NewEnrollmentInitializationResponse

`func NewEnrollmentInitializationResponse() *EnrollmentInitializationResponse`

NewEnrollmentInitializationResponse instantiates a new EnrollmentInitializationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentInitializationResponseWithDefaults

`func NewEnrollmentInitializationResponseWithDefaults() *EnrollmentInitializationResponse`

NewEnrollmentInitializationResponseWithDefaults instantiates a new EnrollmentInitializationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredRequests

`func (o *EnrollmentInitializationResponse) GetCredRequests() []WebAuthnCredRequest`

GetCredRequests returns the CredRequests field if non-nil, zero value otherwise.

### GetCredRequestsOk

`func (o *EnrollmentInitializationResponse) GetCredRequestsOk() (*[]WebAuthnCredRequest, bool)`

GetCredRequestsOk returns a tuple with the CredRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredRequests

`func (o *EnrollmentInitializationResponse) SetCredRequests(v []WebAuthnCredRequest)`

SetCredRequests sets CredRequests field to given value.

### HasCredRequests

`func (o *EnrollmentInitializationResponse) HasCredRequests() bool`

HasCredRequests returns a boolean if a field has been set.

### GetFulfillmentProvider

`func (o *EnrollmentInitializationResponse) GetFulfillmentProvider() string`

GetFulfillmentProvider returns the FulfillmentProvider field if non-nil, zero value otherwise.

### GetFulfillmentProviderOk

`func (o *EnrollmentInitializationResponse) GetFulfillmentProviderOk() (*string, bool)`

GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentProvider

`func (o *EnrollmentInitializationResponse) SetFulfillmentProvider(v string)`

SetFulfillmentProvider sets FulfillmentProvider field to given value.

### HasFulfillmentProvider

`func (o *EnrollmentInitializationResponse) HasFulfillmentProvider() bool`

HasFulfillmentProvider returns a boolean if a field has been set.

### GetPinRequestJwe

`func (o *EnrollmentInitializationResponse) GetPinRequestJwe() string`

GetPinRequestJwe returns the PinRequestJwe field if non-nil, zero value otherwise.

### GetPinRequestJweOk

`func (o *EnrollmentInitializationResponse) GetPinRequestJweOk() (*string, bool)`

GetPinRequestJweOk returns a tuple with the PinRequestJwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinRequestJwe

`func (o *EnrollmentInitializationResponse) SetPinRequestJwe(v string)`

SetPinRequestJwe sets PinRequestJwe field to given value.

### HasPinRequestJwe

`func (o *EnrollmentInitializationResponse) HasPinRequestJwe() bool`

HasPinRequestJwe returns a boolean if a field has been set.

### GetUserId

`func (o *EnrollmentInitializationResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EnrollmentInitializationResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EnrollmentInitializationResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EnrollmentInitializationResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


