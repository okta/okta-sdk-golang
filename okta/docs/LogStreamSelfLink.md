# LogStreamSelfLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | The URI of the resource | 
**Method** | Pointer to **string** | HTTP method allowed for the resource | [optional] 

## Methods

### NewLogStreamSelfLink

`func NewLogStreamSelfLink(href string, ) *LogStreamSelfLink`

NewLogStreamSelfLink instantiates a new LogStreamSelfLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamSelfLinkWithDefaults

`func NewLogStreamSelfLinkWithDefaults() *LogStreamSelfLink`

NewLogStreamSelfLinkWithDefaults instantiates a new LogStreamSelfLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LogStreamSelfLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LogStreamSelfLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LogStreamSelfLink) SetHref(v string)`

SetHref sets Href field to given value.


### GetMethod

`func (o *LogStreamSelfLink) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogStreamSelfLink) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogStreamSelfLink) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogStreamSelfLink) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


