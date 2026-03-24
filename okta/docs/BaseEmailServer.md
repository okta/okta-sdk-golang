# BaseEmailServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Human-readable name for your SMTP server | 
**AuthType** | **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;The authentication type that&#39;s used by your SMTP server | 
**Enabled** | **bool** | If &#x60;true&#x60;, all email traffic is routed through your SMTP server | 
**Host** | **string** | Hostname or IP address of your SMTP server | 
**Id** | Pointer to **string** | ID of your SMTP server | [optional] [readonly] 
**Port** | **int32** | Port number of your SMTP server | 
**Username** | **string** | Username that&#39;s used to access your SMTP server | 

## Methods

### NewBaseEmailServer

`func NewBaseEmailServer(alias string, authType string, enabled bool, host string, port int32, username string, ) *BaseEmailServer`

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


### GetAuthType

`func (o *BaseEmailServer) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *BaseEmailServer) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *BaseEmailServer) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


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


### GetId

`func (o *BaseEmailServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEmailServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEmailServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseEmailServer) HasId() bool`

HasId returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


