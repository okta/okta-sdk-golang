# BaseEmailServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Human-readable name for your SMTP server | [optional] 
**Enabled** | Pointer to **bool** | If &#x60;true&#x60;, routes all email traffic through your SMTP server | [optional] 
**Host** | Pointer to **string** | Hostname or IP address of your SMTP server | [optional] 
**Port** | Pointer to **int32** | Port number of your SMTP server | [optional] 
**Username** | Pointer to **string** | Username used to access your SMTP server | [optional] 

## Methods

### NewBaseEmailServer

`func NewBaseEmailServer() *BaseEmailServer`

NewBaseEmailServer instantiates a new BaseEmailServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEmailServerWithDefaults

`func NewBaseEmailServerWithDefaults() *BaseEmailServer`

NewBaseEmailServerWithDefaults instantiates a new BaseEmailServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *BaseEmailServer) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *BaseEmailServer) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *BaseEmailServer) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *BaseEmailServer) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEnabled

`func (o *BaseEmailServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseEmailServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseEmailServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseEmailServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *BaseEmailServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BaseEmailServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BaseEmailServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *BaseEmailServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *BaseEmailServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BaseEmailServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BaseEmailServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BaseEmailServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *BaseEmailServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BaseEmailServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BaseEmailServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BaseEmailServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


