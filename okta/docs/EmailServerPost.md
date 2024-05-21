# EmailServerPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Human-readable name for your SMTP server | 
**Enabled** | Pointer to **bool** | If &#x60;true&#x60;, routes all email traffic through your SMTP server | [optional] 
**Host** | **string** | Hostname or IP address of your SMTP server | 
**Port** | **int32** | Port number of your SMTP server | 
**Username** | **string** | Username used to access your SMTP server | 
**Password** | **string** | Password used to access your SMTP server | 

## Methods

### NewEmailServerPost

`func NewEmailServerPost(alias string, host string, port int32, username string, password string, ) *EmailServerPost`

NewEmailServerPost instantiates a new EmailServerPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailServerPostWithDefaults

`func NewEmailServerPostWithDefaults() *EmailServerPost`

NewEmailServerPostWithDefaults instantiates a new EmailServerPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *EmailServerPost) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *EmailServerPost) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *EmailServerPost) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetEnabled

`func (o *EmailServerPost) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EmailServerPost) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EmailServerPost) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EmailServerPost) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *EmailServerPost) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EmailServerPost) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EmailServerPost) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *EmailServerPost) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EmailServerPost) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EmailServerPost) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *EmailServerPost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmailServerPost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmailServerPost) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *EmailServerPost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EmailServerPost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EmailServerPost) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


