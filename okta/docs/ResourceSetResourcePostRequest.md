# ResourceSetResourcePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**ResourceConditions**](ResourceConditions.md) |  | 
**ResourceOrnOrUrl** | **string** | Resource in ORN or REST API URL format | 

## Methods

### NewResourceSetResourcePostRequest

`func NewResourceSetResourcePostRequest(conditions ResourceConditions, resourceOrnOrUrl string, ) *ResourceSetResourcePostRequest`

NewResourceSetResourcePostRequest instantiates a new ResourceSetResourcePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetResourcePostRequestWithDefaults

`func NewResourceSetResourcePostRequestWithDefaults() *ResourceSetResourcePostRequest`

NewResourceSetResourcePostRequestWithDefaults instantiates a new ResourceSetResourcePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *ResourceSetResourcePostRequest) GetConditions() ResourceConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResourceSetResourcePostRequest) GetConditionsOk() (*ResourceConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResourceSetResourcePostRequest) SetConditions(v ResourceConditions)`

SetConditions sets Conditions field to given value.


### GetResourceOrnOrUrl

`func (o *ResourceSetResourcePostRequest) GetResourceOrnOrUrl() string`

GetResourceOrnOrUrl returns the ResourceOrnOrUrl field if non-nil, zero value otherwise.

### GetResourceOrnOrUrlOk

`func (o *ResourceSetResourcePostRequest) GetResourceOrnOrUrlOk() (*string, bool)`

GetResourceOrnOrUrlOk returns a tuple with the ResourceOrnOrUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrnOrUrl

`func (o *ResourceSetResourcePostRequest) SetResourceOrnOrUrl(v string)`

SetResourceOrnOrUrl sets ResourceOrnOrUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


