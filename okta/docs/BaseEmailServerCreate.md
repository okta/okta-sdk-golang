# BaseEmailServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Human-readable name for your SMTP server | 
**Enabled** | **bool** | If &#x60;true&#x60;, all email traffic is routed through your SMTP server | 
**Host** | **string** | Hostname or IP address of your SMTP server | 
**Id** | Pointer to **string** | ID of your SMTP server | [optional] [readonly] 
**Port** | **int32** | Port number of your SMTP server | 
**Username** | **string** | Username that&#39;s used to access your SMTP server | 

## Methods

### NewBaseEmailServerCreate

`func NewBaseEmailServerCreate(alias string, enabled bool, host string, port int32, username string, ) *BaseEmailServerCreate`

NewBaseEmailServerCreate instantiates a new BaseEmailServerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEmailServerCreateWithDefaults

`func NewBaseEmailServerCreateWithDefaults() *BaseEmailServerCreate`

NewBaseEmailServerCreateWithDefaults instantiates a new BaseEmailServerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *BaseEmailServerCreate) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *BaseEmailServerCreate) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *BaseEmailServerCreate) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetEnabled

`func (o *BaseEmailServerCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseEmailServerCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseEmailServerCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHost

`func (o *BaseEmailServerCreate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BaseEmailServerCreate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BaseEmailServerCreate) SetHost(v string)`

SetHost sets Host field to given value.


### GetId

`func (o *BaseEmailServerCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEmailServerCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEmailServerCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseEmailServerCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *BaseEmailServerCreate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BaseEmailServerCreate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BaseEmailServerCreate) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *BaseEmailServerCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BaseEmailServerCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BaseEmailServerCreate) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


