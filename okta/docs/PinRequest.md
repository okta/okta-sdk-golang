# PinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorEnrollmentId** | Pointer to **string** | ID for a WebAuthn Preregistration Factor in Okta | [optional] 
**FulfillmentProvider** | Pointer to **string** | Name of the fulfillment provider for the WebAuthn Preregistration Factor | [optional] 
**UserId** | Pointer to **string** | ID of an existing Okta user | [optional] 

## Methods

### NewPinRequest

`func NewPinRequest() *PinRequest`

NewPinRequest instantiates a new PinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinRequestWithDefaults

`func NewPinRequestWithDefaults() *PinRequest`

NewPinRequestWithDefaults instantiates a new PinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorEnrollmentId

`func (o *PinRequest) GetAuthenticatorEnrollmentId() string`

GetAuthenticatorEnrollmentId returns the AuthenticatorEnrollmentId field if non-nil, zero value otherwise.

### GetAuthenticatorEnrollmentIdOk

`func (o *PinRequest) GetAuthenticatorEnrollmentIdOk() (*string, bool)`

GetAuthenticatorEnrollmentIdOk returns a tuple with the AuthenticatorEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorEnrollmentId

`func (o *PinRequest) SetAuthenticatorEnrollmentId(v string)`

SetAuthenticatorEnrollmentId sets AuthenticatorEnrollmentId field to given value.

### HasAuthenticatorEnrollmentId

`func (o *PinRequest) HasAuthenticatorEnrollmentId() bool`

HasAuthenticatorEnrollmentId returns a boolean if a field has been set.

### GetFulfillmentProvider

`func (o *PinRequest) GetFulfillmentProvider() string`

GetFulfillmentProvider returns the FulfillmentProvider field if non-nil, zero value otherwise.

### GetFulfillmentProviderOk

`func (o *PinRequest) GetFulfillmentProviderOk() (*string, bool)`

GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentProvider

`func (o *PinRequest) SetFulfillmentProvider(v string)`

SetFulfillmentProvider sets FulfillmentProvider field to given value.

### HasFulfillmentProvider

`func (o *PinRequest) HasFulfillmentProvider() bool`

HasFulfillmentProvider returns a boolean if a field has been set.

### GetUserId

`func (o *PinRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PinRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PinRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PinRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


