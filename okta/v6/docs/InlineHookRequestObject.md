# InlineHookRequestObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier that Okta assigned to the API request | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the client that made the API request | [optional] 
**Method** | Pointer to **string** | The HTTP request method of the API request | [optional] 
**Url** | Pointer to [**InlineHookRequestObjectUrl**](InlineHookRequestObjectUrl.md) |  | [optional] 

## Methods

### NewInlineHookRequestObject

`func NewInlineHookRequestObject() *InlineHookRequestObject`

NewInlineHookRequestObject instantiates a new InlineHookRequestObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookRequestObjectWithDefaults

`func NewInlineHookRequestObjectWithDefaults() *InlineHookRequestObject`

NewInlineHookRequestObjectWithDefaults instantiates a new InlineHookRequestObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineHookRequestObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineHookRequestObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineHookRequestObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineHookRequestObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *InlineHookRequestObject) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineHookRequestObject) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineHookRequestObject) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineHookRequestObject) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookRequestObject) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookRequestObject) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookRequestObject) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookRequestObject) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *InlineHookRequestObject) GetUrl() InlineHookRequestObjectUrl`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineHookRequestObject) GetUrlOk() (*InlineHookRequestObjectUrl, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineHookRequestObject) SetUrl(v InlineHookRequestObjectUrl)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineHookRequestObject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


