# GracePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to [**GracePeriodExpiry**](GracePeriodExpiry.md) |  | [optional] 
**Type** | Pointer to **string** | Represents the type of Grace Period configured for the device assurance policy | [optional] 

## Methods

### NewGracePeriod

`func NewGracePeriod() *GracePeriod`

NewGracePeriod instantiates a new GracePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGracePeriodWithDefaults

`func NewGracePeriodWithDefaults() *GracePeriod`

NewGracePeriodWithDefaults instantiates a new GracePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *GracePeriod) GetExpiry() GracePeriodExpiry`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GracePeriod) GetExpiryOk() (*GracePeriodExpiry, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GracePeriod) SetExpiry(v GracePeriodExpiry)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GracePeriod) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetType

`func (o *GracePeriod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GracePeriod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GracePeriod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GracePeriod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


