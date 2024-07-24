# StreamConfigurationDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationHeader** | Pointer to **NullableString** | The HTTP Authorization header that is included for each HTTP POST request | [optional] 
**EndpointUrl** | **string** | The target endpoint URL where the transmitter delivers the SET using HTTP POST requests | 
**Method** | **string** | The delivery method that the transmitter uses for delivering a SET | 

## Methods

### NewStreamConfigurationDelivery

`func NewStreamConfigurationDelivery(endpointUrl string, method string, ) *StreamConfigurationDelivery`

NewStreamConfigurationDelivery instantiates a new StreamConfigurationDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConfigurationDeliveryWithDefaults

`func NewStreamConfigurationDeliveryWithDefaults() *StreamConfigurationDelivery`

NewStreamConfigurationDeliveryWithDefaults instantiates a new StreamConfigurationDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationHeader

`func (o *StreamConfigurationDelivery) GetAuthorizationHeader() string`

GetAuthorizationHeader returns the AuthorizationHeader field if non-nil, zero value otherwise.

### GetAuthorizationHeaderOk

`func (o *StreamConfigurationDelivery) GetAuthorizationHeaderOk() (*string, bool)`

GetAuthorizationHeaderOk returns a tuple with the AuthorizationHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationHeader

`func (o *StreamConfigurationDelivery) SetAuthorizationHeader(v string)`

SetAuthorizationHeader sets AuthorizationHeader field to given value.

### HasAuthorizationHeader

`func (o *StreamConfigurationDelivery) HasAuthorizationHeader() bool`

HasAuthorizationHeader returns a boolean if a field has been set.

### SetAuthorizationHeaderNil

`func (o *StreamConfigurationDelivery) SetAuthorizationHeaderNil(b bool)`

 SetAuthorizationHeaderNil sets the value for AuthorizationHeader to be an explicit nil

### UnsetAuthorizationHeader
`func (o *StreamConfigurationDelivery) UnsetAuthorizationHeader()`

UnsetAuthorizationHeader ensures that no value is present for AuthorizationHeader, not even an explicit nil
### GetEndpointUrl

`func (o *StreamConfigurationDelivery) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *StreamConfigurationDelivery) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *StreamConfigurationDelivery) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.


### GetMethod

`func (o *StreamConfigurationDelivery) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *StreamConfigurationDelivery) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *StreamConfigurationDelivery) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


