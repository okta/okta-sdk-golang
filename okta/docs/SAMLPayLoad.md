# SAMLPayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**SAMLPayLoadData**](SAMLPayLoadData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The SAML assertion inline hook type is &#x60;com.okta.saml.tokens.transform&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID and URL of the SAML assertion inline hook | [optional] 

## Methods

### NewSAMLPayLoad

`func NewSAMLPayLoad() *SAMLPayLoad`

NewSAMLPayLoad instantiates a new SAMLPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadWithDefaults

`func NewSAMLPayLoadWithDefaults() *SAMLPayLoad`

NewSAMLPayLoadWithDefaults instantiates a new SAMLPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SAMLPayLoad) GetData() SAMLPayLoadData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SAMLPayLoad) GetDataOk() (*SAMLPayLoadData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SAMLPayLoad) SetData(v SAMLPayLoadData)`

SetData sets Data field to given value.

### HasData

`func (o *SAMLPayLoad) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *SAMLPayLoad) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SAMLPayLoad) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SAMLPayLoad) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *SAMLPayLoad) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *SAMLPayLoad) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SAMLPayLoad) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SAMLPayLoad) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SAMLPayLoad) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


