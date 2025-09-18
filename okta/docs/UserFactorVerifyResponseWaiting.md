# UserFactorVerifyResponseWaiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the verification expires | [optional] [readonly] 
**FactorMessage** | Pointer to **NullableString** | Optional display message for factor verification | [optional] [readonly] 
**FactorResult** | Pointer to **string** | Result of a factor verification | [optional] 
**Profile** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Embedded** | Pointer to [**UserFactorVerifyResponseWaitingEmbedded**](UserFactorVerifyResponseWaitingEmbedded.md) |  | [optional] 
**Links** | Pointer to [**UserFactorLinks**](UserFactorLinks.md) |  | [optional] 

## Methods

### NewUserFactorVerifyResponseWaiting

`func NewUserFactorVerifyResponseWaiting() *UserFactorVerifyResponseWaiting`

NewUserFactorVerifyResponseWaiting instantiates a new UserFactorVerifyResponseWaiting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorVerifyResponseWaitingWithDefaults

`func NewUserFactorVerifyResponseWaitingWithDefaults() *UserFactorVerifyResponseWaiting`

NewUserFactorVerifyResponseWaitingWithDefaults instantiates a new UserFactorVerifyResponseWaiting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *UserFactorVerifyResponseWaiting) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserFactorVerifyResponseWaiting) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserFactorVerifyResponseWaiting) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UserFactorVerifyResponseWaiting) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorMessage

`func (o *UserFactorVerifyResponseWaiting) GetFactorMessage() string`

GetFactorMessage returns the FactorMessage field if non-nil, zero value otherwise.

### GetFactorMessageOk

`func (o *UserFactorVerifyResponseWaiting) GetFactorMessageOk() (*string, bool)`

GetFactorMessageOk returns a tuple with the FactorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorMessage

`func (o *UserFactorVerifyResponseWaiting) SetFactorMessage(v string)`

SetFactorMessage sets FactorMessage field to given value.

### HasFactorMessage

`func (o *UserFactorVerifyResponseWaiting) HasFactorMessage() bool`

HasFactorMessage returns a boolean if a field has been set.

### SetFactorMessageNil

`func (o *UserFactorVerifyResponseWaiting) SetFactorMessageNil(b bool)`

 SetFactorMessageNil sets the value for FactorMessage to be an explicit nil

### UnsetFactorMessage
`func (o *UserFactorVerifyResponseWaiting) UnsetFactorMessage()`

UnsetFactorMessage ensures that no value is present for FactorMessage, not even an explicit nil
### GetFactorResult

`func (o *UserFactorVerifyResponseWaiting) GetFactorResult() string`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *UserFactorVerifyResponseWaiting) GetFactorResultOk() (*string, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *UserFactorVerifyResponseWaiting) SetFactorResult(v string)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *UserFactorVerifyResponseWaiting) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetProfile

`func (o *UserFactorVerifyResponseWaiting) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorVerifyResponseWaiting) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorVerifyResponseWaiting) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorVerifyResponseWaiting) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetEmbedded

`func (o *UserFactorVerifyResponseWaiting) GetEmbedded() UserFactorVerifyResponseWaitingEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UserFactorVerifyResponseWaiting) GetEmbeddedOk() (*UserFactorVerifyResponseWaitingEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UserFactorVerifyResponseWaiting) SetEmbedded(v UserFactorVerifyResponseWaitingEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UserFactorVerifyResponseWaiting) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *UserFactorVerifyResponseWaiting) GetLinks() UserFactorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFactorVerifyResponseWaiting) GetLinksOk() (*UserFactorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFactorVerifyResponseWaiting) SetLinks(v UserFactorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserFactorVerifyResponseWaiting) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


