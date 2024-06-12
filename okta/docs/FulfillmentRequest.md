# FulfillmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentData** | Pointer to [**FulfillmentData**](FulfillmentData.md) |  | [optional] 
**FulfillmentProvider** | Pointer to **string** | Name of the fulfillment provider for the WebAuthn Preregistration Factor | [optional] 
**UserId** | Pointer to **string** | ID of an existing Okta user | [optional] 

## Methods

### NewFulfillmentRequest

`func NewFulfillmentRequest() *FulfillmentRequest`

NewFulfillmentRequest instantiates a new FulfillmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentRequestWithDefaults

`func NewFulfillmentRequestWithDefaults() *FulfillmentRequest`

NewFulfillmentRequestWithDefaults instantiates a new FulfillmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentData

`func (o *FulfillmentRequest) GetFulfillmentData() FulfillmentData`

GetFulfillmentData returns the FulfillmentData field if non-nil, zero value otherwise.

### GetFulfillmentDataOk

`func (o *FulfillmentRequest) GetFulfillmentDataOk() (*FulfillmentData, bool)`

GetFulfillmentDataOk returns a tuple with the FulfillmentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentData

`func (o *FulfillmentRequest) SetFulfillmentData(v FulfillmentData)`

SetFulfillmentData sets FulfillmentData field to given value.

### HasFulfillmentData

`func (o *FulfillmentRequest) HasFulfillmentData() bool`

HasFulfillmentData returns a boolean if a field has been set.

### GetFulfillmentProvider

`func (o *FulfillmentRequest) GetFulfillmentProvider() string`

GetFulfillmentProvider returns the FulfillmentProvider field if non-nil, zero value otherwise.

### GetFulfillmentProviderOk

`func (o *FulfillmentRequest) GetFulfillmentProviderOk() (*string, bool)`

GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentProvider

`func (o *FulfillmentRequest) SetFulfillmentProvider(v string)`

SetFulfillmentProvider sets FulfillmentProvider field to given value.

### HasFulfillmentProvider

`func (o *FulfillmentRequest) HasFulfillmentProvider() bool`

HasFulfillmentProvider returns a boolean if a field has been set.

### GetUserId

`func (o *FulfillmentRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FulfillmentRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FulfillmentRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FulfillmentRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


