# EnrollmentInitializationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnrollmentRpIds** | Pointer to **[]string** | List of Relying Party hostnames to register on the YubiKey. | [optional] 
**FulfillmentProvider** | Pointer to **string** | Name of the fulfillment provider for the WebAuthn Preregistration Factor | [optional] 
**UserId** | Pointer to **string** | ID of an existing Okta user | [optional] 
**YubicoTransportKeyJWK** | Pointer to [**ECKeyJWK**](ECKeyJWK.md) |  | [optional] 

## Methods

### NewEnrollmentInitializationRequest

`func NewEnrollmentInitializationRequest() *EnrollmentInitializationRequest`

NewEnrollmentInitializationRequest instantiates a new EnrollmentInitializationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentInitializationRequestWithDefaults

`func NewEnrollmentInitializationRequestWithDefaults() *EnrollmentInitializationRequest`

NewEnrollmentInitializationRequestWithDefaults instantiates a new EnrollmentInitializationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnrollmentRpIds

`func (o *EnrollmentInitializationRequest) GetEnrollmentRpIds() []string`

GetEnrollmentRpIds returns the EnrollmentRpIds field if non-nil, zero value otherwise.

### GetEnrollmentRpIdsOk

`func (o *EnrollmentInitializationRequest) GetEnrollmentRpIdsOk() (*[]string, bool)`

GetEnrollmentRpIdsOk returns a tuple with the EnrollmentRpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentRpIds

`func (o *EnrollmentInitializationRequest) SetEnrollmentRpIds(v []string)`

SetEnrollmentRpIds sets EnrollmentRpIds field to given value.

### HasEnrollmentRpIds

`func (o *EnrollmentInitializationRequest) HasEnrollmentRpIds() bool`

HasEnrollmentRpIds returns a boolean if a field has been set.

### GetFulfillmentProvider

`func (o *EnrollmentInitializationRequest) GetFulfillmentProvider() string`

GetFulfillmentProvider returns the FulfillmentProvider field if non-nil, zero value otherwise.

### GetFulfillmentProviderOk

`func (o *EnrollmentInitializationRequest) GetFulfillmentProviderOk() (*string, bool)`

GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentProvider

`func (o *EnrollmentInitializationRequest) SetFulfillmentProvider(v string)`

SetFulfillmentProvider sets FulfillmentProvider field to given value.

### HasFulfillmentProvider

`func (o *EnrollmentInitializationRequest) HasFulfillmentProvider() bool`

HasFulfillmentProvider returns a boolean if a field has been set.

### GetUserId

`func (o *EnrollmentInitializationRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EnrollmentInitializationRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EnrollmentInitializationRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EnrollmentInitializationRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetYubicoTransportKeyJWK

`func (o *EnrollmentInitializationRequest) GetYubicoTransportKeyJWK() ECKeyJWK`

GetYubicoTransportKeyJWK returns the YubicoTransportKeyJWK field if non-nil, zero value otherwise.

### GetYubicoTransportKeyJWKOk

`func (o *EnrollmentInitializationRequest) GetYubicoTransportKeyJWKOk() (*ECKeyJWK, bool)`

GetYubicoTransportKeyJWKOk returns a tuple with the YubicoTransportKeyJWK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubicoTransportKeyJWK

`func (o *EnrollmentInitializationRequest) SetYubicoTransportKeyJWK(v ECKeyJWK)`

SetYubicoTransportKeyJWK sets YubicoTransportKeyJWK field to given value.

### HasYubicoTransportKeyJWK

`func (o *EnrollmentInitializationRequest) HasYubicoTransportKeyJWK() bool`

HasYubicoTransportKeyJWK returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


