# CimdClientEntityBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundEntityId** | **string** | Unique identifier of the entity to bind to the CIMD client entity | 
**BoundEntityType** | **string** | The type of resource that&#39;s bound to the CIMD client entity | 

## Methods

### NewCimdClientEntityBindingRequest

`func NewCimdClientEntityBindingRequest(boundEntityId string, boundEntityType string, ) *CimdClientEntityBindingRequest`

NewCimdClientEntityBindingRequest instantiates a new CimdClientEntityBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdClientEntityBindingRequestWithDefaults

`func NewCimdClientEntityBindingRequestWithDefaults() *CimdClientEntityBindingRequest`

NewCimdClientEntityBindingRequestWithDefaults instantiates a new CimdClientEntityBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundEntityId

`func (o *CimdClientEntityBindingRequest) GetBoundEntityId() string`

GetBoundEntityId returns the BoundEntityId field if non-nil, zero value otherwise.

### GetBoundEntityIdOk

`func (o *CimdClientEntityBindingRequest) GetBoundEntityIdOk() (*string, bool)`

GetBoundEntityIdOk returns a tuple with the BoundEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEntityId

`func (o *CimdClientEntityBindingRequest) SetBoundEntityId(v string)`

SetBoundEntityId sets BoundEntityId field to given value.


### GetBoundEntityType

`func (o *CimdClientEntityBindingRequest) GetBoundEntityType() string`

GetBoundEntityType returns the BoundEntityType field if non-nil, zero value otherwise.

### GetBoundEntityTypeOk

`func (o *CimdClientEntityBindingRequest) GetBoundEntityTypeOk() (*string, bool)`

GetBoundEntityTypeOk returns a tuple with the BoundEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEntityType

`func (o *CimdClientEntityBindingRequest) SetBoundEntityType(v string)`

SetBoundEntityType sets BoundEntityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


