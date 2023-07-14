# VerifyUserFactorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**FactorResult** | Pointer to [**VerifyUserFactorResult**](VerifyUserFactorResult.md) |  | [optional] 
**FactorResultMessage** | Pointer to **string** |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewVerifyUserFactorResponse

`func NewVerifyUserFactorResponse() *VerifyUserFactorResponse`

NewVerifyUserFactorResponse instantiates a new VerifyUserFactorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyUserFactorResponseWithDefaults

`func NewVerifyUserFactorResponseWithDefaults() *VerifyUserFactorResponse`

NewVerifyUserFactorResponseWithDefaults instantiates a new VerifyUserFactorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *VerifyUserFactorResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VerifyUserFactorResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VerifyUserFactorResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *VerifyUserFactorResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorResult

`func (o *VerifyUserFactorResponse) GetFactorResult() VerifyUserFactorResult`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *VerifyUserFactorResponse) GetFactorResultOk() (*VerifyUserFactorResult, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *VerifyUserFactorResponse) SetFactorResult(v VerifyUserFactorResult)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *VerifyUserFactorResponse) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetFactorResultMessage

`func (o *VerifyUserFactorResponse) GetFactorResultMessage() string`

GetFactorResultMessage returns the FactorResultMessage field if non-nil, zero value otherwise.

### GetFactorResultMessageOk

`func (o *VerifyUserFactorResponse) GetFactorResultMessageOk() (*string, bool)`

GetFactorResultMessageOk returns a tuple with the FactorResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResultMessage

`func (o *VerifyUserFactorResponse) SetFactorResultMessage(v string)`

SetFactorResultMessage sets FactorResultMessage field to given value.

### HasFactorResultMessage

`func (o *VerifyUserFactorResponse) HasFactorResultMessage() bool`

HasFactorResultMessage returns a boolean if a field has been set.

### GetEmbedded

`func (o *VerifyUserFactorResponse) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *VerifyUserFactorResponse) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *VerifyUserFactorResponse) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *VerifyUserFactorResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *VerifyUserFactorResponse) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VerifyUserFactorResponse) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VerifyUserFactorResponse) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VerifyUserFactorResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


