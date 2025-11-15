# ValidationDetailProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The unique identifier of the action flow in the provider system | 
**Type** | **string** | Type of action provider | 

## Methods

### NewValidationDetailProvider

`func NewValidationDetailProvider(externalId string, type_ string, ) *ValidationDetailProvider`

NewValidationDetailProvider instantiates a new ValidationDetailProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationDetailProviderWithDefaults

`func NewValidationDetailProviderWithDefaults() *ValidationDetailProvider`

NewValidationDetailProviderWithDefaults instantiates a new ValidationDetailProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ValidationDetailProvider) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ValidationDetailProvider) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ValidationDetailProvider) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *ValidationDetailProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidationDetailProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidationDetailProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


