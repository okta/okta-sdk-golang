# TokenPayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TokenPayLoadData**](TokenPayLoadData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The token inline hook type is &#x60;com.okta.oauth2.tokens.transform&#x60;. | [optional] 
**Source** | Pointer to **string** | The URL of the token inline hook | [optional] 

## Methods

### NewTokenPayLoad

`func NewTokenPayLoad() *TokenPayLoad`

NewTokenPayLoad instantiates a new TokenPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPayLoadWithDefaults

`func NewTokenPayLoadWithDefaults() *TokenPayLoad`

NewTokenPayLoadWithDefaults instantiates a new TokenPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenPayLoad) GetData() TokenPayLoadData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenPayLoad) GetDataOk() (*TokenPayLoadData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenPayLoad) SetData(v TokenPayLoadData)`

SetData sets Data field to given value.

### HasData

`func (o *TokenPayLoad) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *TokenPayLoad) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TokenPayLoad) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TokenPayLoad) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TokenPayLoad) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *TokenPayLoad) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenPayLoad) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenPayLoad) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TokenPayLoad) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


