# SamlAlgorithms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**SamlRequestAlgorithm**](SamlRequestAlgorithm.md) |  | [optional] 
**Response** | Pointer to [**SamlResponseAlgorithm**](SamlResponseAlgorithm.md) |  | [optional] 

## Methods

### NewSamlAlgorithms

`func NewSamlAlgorithms() *SamlAlgorithms`

NewSamlAlgorithms instantiates a new SamlAlgorithms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAlgorithmsWithDefaults

`func NewSamlAlgorithmsWithDefaults() *SamlAlgorithms`

NewSamlAlgorithmsWithDefaults instantiates a new SamlAlgorithms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *SamlAlgorithms) GetRequest() SamlRequestAlgorithm`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SamlAlgorithms) GetRequestOk() (*SamlRequestAlgorithm, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SamlAlgorithms) SetRequest(v SamlRequestAlgorithm)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *SamlAlgorithms) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *SamlAlgorithms) GetResponse() SamlResponseAlgorithm`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SamlAlgorithms) GetResponseOk() (*SamlResponseAlgorithm, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SamlAlgorithms) SetResponse(v SamlResponseAlgorithm)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SamlAlgorithms) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


