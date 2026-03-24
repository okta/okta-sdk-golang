# BotProtectionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforcementType** | Pointer to **string** | The type of enforcement to trigger when a bot is detected | [optional] 
**Level** | **string** | The sensitivity level of bot detection | 
**Mode** | **string** | The enforcement mode for bot protection | 
**SupportedFlows** | Pointer to **[]string** | An array of authentication flows that have bot protection enabled | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewBotProtectionConfiguration

`func NewBotProtectionConfiguration(level string, mode string, ) *BotProtectionConfiguration`

NewBotProtectionConfiguration instantiates a new BotProtectionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotProtectionConfigurationWithDefaults

`func NewBotProtectionConfigurationWithDefaults() *BotProtectionConfiguration`

NewBotProtectionConfigurationWithDefaults instantiates a new BotProtectionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforcementType

`func (o *BotProtectionConfiguration) GetEnforcementType() string`

GetEnforcementType returns the EnforcementType field if non-nil, zero value otherwise.

### GetEnforcementTypeOk

`func (o *BotProtectionConfiguration) GetEnforcementTypeOk() (*string, bool)`

GetEnforcementTypeOk returns a tuple with the EnforcementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementType

`func (o *BotProtectionConfiguration) SetEnforcementType(v string)`

SetEnforcementType sets EnforcementType field to given value.

### HasEnforcementType

`func (o *BotProtectionConfiguration) HasEnforcementType() bool`

HasEnforcementType returns a boolean if a field has been set.

### GetLevel

`func (o *BotProtectionConfiguration) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *BotProtectionConfiguration) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *BotProtectionConfiguration) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMode

`func (o *BotProtectionConfiguration) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BotProtectionConfiguration) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BotProtectionConfiguration) SetMode(v string)`

SetMode sets Mode field to given value.


### GetSupportedFlows

`func (o *BotProtectionConfiguration) GetSupportedFlows() []string`

GetSupportedFlows returns the SupportedFlows field if non-nil, zero value otherwise.

### GetSupportedFlowsOk

`func (o *BotProtectionConfiguration) GetSupportedFlowsOk() (*[]string, bool)`

GetSupportedFlowsOk returns a tuple with the SupportedFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFlows

`func (o *BotProtectionConfiguration) SetSupportedFlows(v []string)`

SetSupportedFlows sets SupportedFlows field to given value.

### HasSupportedFlows

`func (o *BotProtectionConfiguration) HasSupportedFlows() bool`

HasSupportedFlows returns a boolean if a field has been set.

### GetLinks

`func (o *BotProtectionConfiguration) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BotProtectionConfiguration) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BotProtectionConfiguration) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BotProtectionConfiguration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


