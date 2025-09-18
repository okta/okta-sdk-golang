# SAMLPayLoadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**SAMLPayLoadDataContext**](SAMLPayLoadDataContext.md) |  | [optional] 
**Assertion** | Pointer to [**SAMLPayLoadDataAssertion**](SAMLPayLoadDataAssertion.md) |  | [optional] 

## Methods

### NewSAMLPayLoadData

`func NewSAMLPayLoadData() *SAMLPayLoadData`

NewSAMLPayLoadData instantiates a new SAMLPayLoadData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadDataWithDefaults

`func NewSAMLPayLoadDataWithDefaults() *SAMLPayLoadData`

NewSAMLPayLoadDataWithDefaults instantiates a new SAMLPayLoadData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SAMLPayLoadData) GetContext() SAMLPayLoadDataContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SAMLPayLoadData) GetContextOk() (*SAMLPayLoadDataContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SAMLPayLoadData) SetContext(v SAMLPayLoadDataContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SAMLPayLoadData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAssertion

`func (o *SAMLPayLoadData) GetAssertion() SAMLPayLoadDataAssertion`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *SAMLPayLoadData) GetAssertionOk() (*SAMLPayLoadDataAssertion, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *SAMLPayLoadData) SetAssertion(v SAMLPayLoadDataAssertion)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *SAMLPayLoadData) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


