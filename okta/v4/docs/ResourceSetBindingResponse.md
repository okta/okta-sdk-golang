# ResourceSetBindingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | &#x60;id&#x60; of the role | [optional] 
**Links** | Pointer to [**ResourceSetBindingResponseLinks**](ResourceSetBindingResponseLinks.md) |  | [optional] 

## Methods

### NewResourceSetBindingResponse

`func NewResourceSetBindingResponse() *ResourceSetBindingResponse`

NewResourceSetBindingResponse instantiates a new ResourceSetBindingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingResponseWithDefaults

`func NewResourceSetBindingResponseWithDefaults() *ResourceSetBindingResponse`

NewResourceSetBindingResponseWithDefaults instantiates a new ResourceSetBindingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceSetBindingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceSetBindingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceSetBindingResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceSetBindingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSetBindingResponse) GetLinks() ResourceSetBindingResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSetBindingResponse) GetLinksOk() (*ResourceSetBindingResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSetBindingResponse) SetLinks(v ResourceSetBindingResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSetBindingResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


