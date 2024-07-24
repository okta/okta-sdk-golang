# UserFactorVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the verification expires | [optional] [readonly] 
**FactorMessage** | Pointer to **NullableString** | Optional display message for Factor verification | [optional] [readonly] 
**FactorResult** | Pointer to **string** | Result of a Factor verification | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**UserFactorLinks**](UserFactorLinks.md) |  | [optional] 

## Methods

### NewUserFactorVerifyResponse

`func NewUserFactorVerifyResponse() *UserFactorVerifyResponse`

NewUserFactorVerifyResponse instantiates a new UserFactorVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorVerifyResponseWithDefaults

`func NewUserFactorVerifyResponseWithDefaults() *UserFactorVerifyResponse`

NewUserFactorVerifyResponseWithDefaults instantiates a new UserFactorVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *UserFactorVerifyResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserFactorVerifyResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserFactorVerifyResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UserFactorVerifyResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorMessage

`func (o *UserFactorVerifyResponse) GetFactorMessage() string`

GetFactorMessage returns the FactorMessage field if non-nil, zero value otherwise.

### GetFactorMessageOk

`func (o *UserFactorVerifyResponse) GetFactorMessageOk() (*string, bool)`

GetFactorMessageOk returns a tuple with the FactorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorMessage

`func (o *UserFactorVerifyResponse) SetFactorMessage(v string)`

SetFactorMessage sets FactorMessage field to given value.

### HasFactorMessage

`func (o *UserFactorVerifyResponse) HasFactorMessage() bool`

HasFactorMessage returns a boolean if a field has been set.

### SetFactorMessageNil

`func (o *UserFactorVerifyResponse) SetFactorMessageNil(b bool)`

 SetFactorMessageNil sets the value for FactorMessage to be an explicit nil

### UnsetFactorMessage
`func (o *UserFactorVerifyResponse) UnsetFactorMessage()`

UnsetFactorMessage ensures that no value is present for FactorMessage, not even an explicit nil
### GetFactorResult

`func (o *UserFactorVerifyResponse) GetFactorResult() string`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *UserFactorVerifyResponse) GetFactorResultOk() (*string, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *UserFactorVerifyResponse) SetFactorResult(v string)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *UserFactorVerifyResponse) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetEmbedded

`func (o *UserFactorVerifyResponse) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UserFactorVerifyResponse) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UserFactorVerifyResponse) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UserFactorVerifyResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *UserFactorVerifyResponse) GetLinks() UserFactorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFactorVerifyResponse) GetLinksOk() (*UserFactorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFactorVerifyResponse) SetLinks(v UserFactorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserFactorVerifyResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


