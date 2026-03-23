# LogRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionName** | Pointer to **NullableString** | The name of the detection mechanism that identified the risk | [optional] [readonly] 
**Issuer** | Pointer to **NullableString** | The entity that issued the associated risk | [optional] [readonly] 
**Level** | Pointer to **NullableString** | The risk level associated with the request | [optional] [readonly] 
**PreviousLevel** | Pointer to **NullableString** | The previous risk level (if any) associated with the user | [optional] [readonly] 
**Reasons** | Pointer to **[]string** | Reasons for the associated risk | [optional] [readonly] 

## Methods

### NewLogRisk

`func NewLogRisk() *LogRisk`

NewLogRisk instantiates a new LogRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogRiskWithDefaults

`func NewLogRiskWithDefaults() *LogRisk`

NewLogRiskWithDefaults instantiates a new LogRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionName

`func (o *LogRisk) GetDetectionName() string`

GetDetectionName returns the DetectionName field if non-nil, zero value otherwise.

### GetDetectionNameOk

`func (o *LogRisk) GetDetectionNameOk() (*string, bool)`

GetDetectionNameOk returns a tuple with the DetectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionName

`func (o *LogRisk) SetDetectionName(v string)`

SetDetectionName sets DetectionName field to given value.

### HasDetectionName

`func (o *LogRisk) HasDetectionName() bool`

HasDetectionName returns a boolean if a field has been set.

### SetDetectionNameNil

`func (o *LogRisk) SetDetectionNameNil(b bool)`

 SetDetectionNameNil sets the value for DetectionName to be an explicit nil

### UnsetDetectionName
`func (o *LogRisk) UnsetDetectionName()`

UnsetDetectionName ensures that no value is present for DetectionName, not even an explicit nil
### GetIssuer

`func (o *LogRisk) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *LogRisk) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *LogRisk) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *LogRisk) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *LogRisk) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *LogRisk) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetLevel

`func (o *LogRisk) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogRisk) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogRisk) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LogRisk) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *LogRisk) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *LogRisk) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetPreviousLevel

`func (o *LogRisk) GetPreviousLevel() string`

GetPreviousLevel returns the PreviousLevel field if non-nil, zero value otherwise.

### GetPreviousLevelOk

`func (o *LogRisk) GetPreviousLevelOk() (*string, bool)`

GetPreviousLevelOk returns a tuple with the PreviousLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLevel

`func (o *LogRisk) SetPreviousLevel(v string)`

SetPreviousLevel sets PreviousLevel field to given value.

### HasPreviousLevel

`func (o *LogRisk) HasPreviousLevel() bool`

HasPreviousLevel returns a boolean if a field has been set.

### SetPreviousLevelNil

`func (o *LogRisk) SetPreviousLevelNil(b bool)`

 SetPreviousLevelNil sets the value for PreviousLevel to be an explicit nil

### UnsetPreviousLevel
`func (o *LogRisk) UnsetPreviousLevel()`

UnsetPreviousLevel ensures that no value is present for PreviousLevel, not even an explicit nil
### GetReasons

`func (o *LogRisk) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *LogRisk) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *LogRisk) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *LogRisk) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### SetReasonsNil

`func (o *LogRisk) SetReasonsNil(b bool)`

 SetReasonsNil sets the value for Reasons to be an explicit nil

### UnsetReasons
`func (o *LogRisk) UnsetReasons()`

UnsetReasons ensures that no value is present for Reasons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


