# InlineHookRequestObjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier that Okta assigned to the API request | [optional] 
**Method** | Pointer to **string** | The HTTP request method of the API request | [optional] 
**Url** | Pointer to [**InlineHookRequestObjectRequestUrl**](InlineHookRequestObjectRequestUrl.md) |  | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the client that made the API request | [optional] 

## Methods

### NewInlineHookRequestObjectRequest

`func NewInlineHookRequestObjectRequest() *InlineHookRequestObjectRequest`

NewInlineHookRequestObjectRequest instantiates a new InlineHookRequestObjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookRequestObjectRequestWithDefaults

`func NewInlineHookRequestObjectRequestWithDefaults() *InlineHookRequestObjectRequest`

NewInlineHookRequestObjectRequestWithDefaults instantiates a new InlineHookRequestObjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineHookRequestObjectRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineHookRequestObjectRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineHookRequestObjectRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineHookRequestObjectRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookRequestObjectRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookRequestObjectRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookRequestObjectRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookRequestObjectRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUrl

`func (o *InlineHookRequestObjectRequest) GetUrl() InlineHookRequestObjectRequestUrl`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineHookRequestObjectRequest) GetUrlOk() (*InlineHookRequestObjectRequestUrl, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineHookRequestObjectRequest) SetUrl(v InlineHookRequestObjectRequestUrl)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineHookRequestObjectRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIpAddress

`func (o *InlineHookRequestObjectRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineHookRequestObjectRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineHookRequestObjectRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineHookRequestObjectRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


