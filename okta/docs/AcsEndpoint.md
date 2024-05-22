# AcsEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Index of the URL in the array of ACS endpoints | 
**Url** | **string** | URL of the ACS | 

## Methods

### NewAcsEndpoint

`func NewAcsEndpoint(index int32, url string, ) *AcsEndpoint`

NewAcsEndpoint instantiates a new AcsEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcsEndpointWithDefaults

`func NewAcsEndpointWithDefaults() *AcsEndpoint`

NewAcsEndpointWithDefaults instantiates a new AcsEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *AcsEndpoint) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AcsEndpoint) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AcsEndpoint) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetUrl

`func (o *AcsEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AcsEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AcsEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


