# ResourceSetResourcePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**ResourceConditions**](ResourceConditions.md) |  | [optional] 

## Methods

### NewResourceSetResourcePutRequest

`func NewResourceSetResourcePutRequest() *ResourceSetResourcePutRequest`

NewResourceSetResourcePutRequest instantiates a new ResourceSetResourcePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetResourcePutRequestWithDefaults

`func NewResourceSetResourcePutRequestWithDefaults() *ResourceSetResourcePutRequest`

NewResourceSetResourcePutRequestWithDefaults instantiates a new ResourceSetResourcePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *ResourceSetResourcePutRequest) GetConditions() ResourceConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResourceSetResourcePutRequest) GetConditionsOk() (*ResourceConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResourceSetResourcePutRequest) SetConditions(v ResourceConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ResourceSetResourcePutRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


