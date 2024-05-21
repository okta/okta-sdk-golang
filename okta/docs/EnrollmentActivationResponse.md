# EnrollmentActivationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorEnrollmentIds** | Pointer to **[]string** | List of IDs for preregistered WebAuthn Factors in Okta | [optional] 
**FulfillmentProvider** | Pointer to **string** | Name of the fulfillment provider for the WebAuthn Preregistration Factor | [optional] 
**UserId** | Pointer to **string** | ID of an existing Okta user | [optional] 

## Methods

### NewEnrollmentActivationResponse

`func NewEnrollmentActivationResponse() *EnrollmentActivationResponse`

NewEnrollmentActivationResponse instantiates a new EnrollmentActivationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentActivationResponseWithDefaults

`func NewEnrollmentActivationResponseWithDefaults() *EnrollmentActivationResponse`

NewEnrollmentActivationResponseWithDefaults instantiates a new EnrollmentActivationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorEnrollmentIds

`func (o *EnrollmentActivationResponse) GetAuthenticatorEnrollmentIds() []string`

GetAuthenticatorEnrollmentIds returns the AuthenticatorEnrollmentIds field if non-nil, zero value otherwise.

### GetAuthenticatorEnrollmentIdsOk

`func (o *EnrollmentActivationResponse) GetAuthenticatorEnrollmentIdsOk() (*[]string, bool)`

GetAuthenticatorEnrollmentIdsOk returns a tuple with the AuthenticatorEnrollmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorEnrollmentIds

`func (o *EnrollmentActivationResponse) SetAuthenticatorEnrollmentIds(v []string)`

SetAuthenticatorEnrollmentIds sets AuthenticatorEnrollmentIds field to given value.

### HasAuthenticatorEnrollmentIds

`func (o *EnrollmentActivationResponse) HasAuthenticatorEnrollmentIds() bool`

HasAuthenticatorEnrollmentIds returns a boolean if a field has been set.

### GetFulfillmentProvider

`func (o *EnrollmentActivationResponse) GetFulfillmentProvider() string`

GetFulfillmentProvider returns the FulfillmentProvider field if non-nil, zero value otherwise.

### GetFulfillmentProviderOk

`func (o *EnrollmentActivationResponse) GetFulfillmentProviderOk() (*string, bool)`

GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentProvider

`func (o *EnrollmentActivationResponse) SetFulfillmentProvider(v string)`

SetFulfillmentProvider sets FulfillmentProvider field to given value.

### HasFulfillmentProvider

`func (o *EnrollmentActivationResponse) HasFulfillmentProvider() bool`

HasFulfillmentProvider returns a boolean if a field has been set.

### GetUserId

`func (o *EnrollmentActivationResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EnrollmentActivationResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EnrollmentActivationResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EnrollmentActivationResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


