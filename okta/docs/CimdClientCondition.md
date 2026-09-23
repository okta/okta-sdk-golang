# CimdClientCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchType** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;The type of CIMD client to match | [optional] 
**Value** | Pointer to **string** | The ID of the app or agent to match, based on &#x60;matchType&#x60;. For &#x60;APP&#x60;, this is the app instance ID. For &#x60;AGENT&#x60;, this is the AI agent ID. | [optional] 

## Methods

### NewCimdClientCondition

`func NewCimdClientCondition() *CimdClientCondition`

NewCimdClientCondition instantiates a new CimdClientCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdClientConditionWithDefaults

`func NewCimdClientConditionWithDefaults() *CimdClientCondition`

NewCimdClientConditionWithDefaults instantiates a new CimdClientCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchType

`func (o *CimdClientCondition) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *CimdClientCondition) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *CimdClientCondition) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *CimdClientCondition) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetValue

`func (o *CimdClientCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CimdClientCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CimdClientCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CimdClientCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


