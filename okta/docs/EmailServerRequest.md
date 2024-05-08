# EmailServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Human-readable name for your SMTP server | [optional] 
**Enabled** | Pointer to **bool** | If &#x60;true&#x60;, routes all email traffic through your SMTP server | [optional] 
**Host** | Pointer to **string** | Hostname or IP address of your SMTP server | [optional] 
**Port** | Pointer to **int32** | Port number of your SMTP server | [optional] 
**Username** | Pointer to **string** | Username used to access your SMTP server | [optional] 
**Password** | Pointer to **string** | Password used to access your SMTP server | [optional] 

## Methods

### NewEmailServerRequest

`func NewEmailServerRequest() *EmailServerRequest`

NewEmailServerRequest instantiates a new EmailServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailServerRequestWithDefaults

`func NewEmailServerRequestWithDefaults() *EmailServerRequest`

NewEmailServerRequestWithDefaults instantiates a new EmailServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *EmailServerRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *EmailServerRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *EmailServerRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *EmailServerRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEnabled

`func (o *EmailServerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailServerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailServerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EmailServerRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *EmailServerRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EmailServerRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EmailServerRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EmailServerRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *EmailServerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailServerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailServerRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EmailServerRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *EmailServerRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmailServerRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmailServerRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EmailServerRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *EmailServerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EmailServerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EmailServerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EmailServerRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


