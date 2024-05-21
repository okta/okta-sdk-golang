# EmailServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Human-readable name for your SMTP server | [optional] 
**Enabled** | Pointer to **bool** | If &#x60;true&#x60;, routes all email traffic through your SMTP server | [optional] 
**Host** | Pointer to **string** | Hostname or IP address of your SMTP server | [optional] 
**Port** | Pointer to **int32** | Port number of your SMTP server | [optional] 
**Username** | Pointer to **string** | Username used to access your SMTP server | [optional] 
**Id** | Pointer to **string** | ID of your SMTP server | [optional] 

## Methods

### NewEmailServerResponse

`func NewEmailServerResponse() *EmailServerResponse`

NewEmailServerResponse instantiates a new EmailServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailServerResponseWithDefaults

`func NewEmailServerResponseWithDefaults() *EmailServerResponse`

NewEmailServerResponseWithDefaults instantiates a new EmailServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *EmailServerResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *EmailServerResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *EmailServerResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *EmailServerResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEnabled

`func (o *EmailServerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailServerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailServerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EmailServerResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *EmailServerResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EmailServerResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EmailServerResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EmailServerResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *EmailServerResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailServerResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailServerResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EmailServerResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *EmailServerResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmailServerResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmailServerResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EmailServerResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetId

`func (o *EmailServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailServerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailServerResponse) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


