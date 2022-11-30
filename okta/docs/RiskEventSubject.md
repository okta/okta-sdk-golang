# RiskEventSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | Either an IpV4 or IpV6 address. | 
**Message** | Pointer to **string** | Any additional message that the provider can send specifying the reason for the risk level of the IP. | [optional] 
**RiskLevel** | Pointer to [**RiskEventSubjectRiskLevel**](RiskEventSubjectRiskLevel.md) |  | [optional] 

## Methods

### NewRiskEventSubject

`func NewRiskEventSubject(ip string, ) *RiskEventSubject`

NewRiskEventSubject instantiates a new RiskEventSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEventSubjectWithDefaults

`func NewRiskEventSubjectWithDefaults() *RiskEventSubject`

NewRiskEventSubjectWithDefaults instantiates a new RiskEventSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *RiskEventSubject) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RiskEventSubject) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RiskEventSubject) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMessage

`func (o *RiskEventSubject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RiskEventSubject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RiskEventSubject) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RiskEventSubject) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRiskLevel

`func (o *RiskEventSubject) GetRiskLevel() RiskEventSubjectRiskLevel`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *RiskEventSubject) GetRiskLevelOk() (*RiskEventSubjectRiskLevel, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *RiskEventSubject) SetRiskLevel(v RiskEventSubjectRiskLevel)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *RiskEventSubject) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


