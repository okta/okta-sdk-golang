# LogStreamLinkObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | The URI of the resource | 
**Method** | Pointer to **string** | HTTP method allowed for the resource | [optional] 

## Methods

### NewLogStreamLinkObject

`func NewLogStreamLinkObject(href string, ) *LogStreamLinkObject`

NewLogStreamLinkObject instantiates a new LogStreamLinkObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamLinkObjectWithDefaults

`func NewLogStreamLinkObjectWithDefaults() *LogStreamLinkObject`

NewLogStreamLinkObjectWithDefaults instantiates a new LogStreamLinkObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LogStreamLinkObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LogStreamLinkObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LogStreamLinkObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetMethod

`func (o *LogStreamLinkObject) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogStreamLinkObject) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogStreamLinkObject) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogStreamLinkObject) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


