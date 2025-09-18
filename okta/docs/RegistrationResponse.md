# RegistrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]RegistrationResponseCommandsInner**](RegistrationResponseCommandsInner.md) | The &#x60;commands&#x60; object lets you invoke commands to modify or add values to the attributes in the Okta user profile that are created for this user. The object also lets you control whether or not the registration attempt is allowed to proceed.  This object is an array, allowing you to send multiple commands in your response. Each array element requires a &#x60;type&#x60; property and a &#x60;value&#x60; property. The &#x60;type&#x60; property is where you specify which of the supported commands you wish to execute, and &#x60;value&#x60; is where you supply parameters for that command.  The registration inline hook supports these three commands: * &#x60;com.okta.user.profile.update&#x60;: Change attribute values in the user&#39;s Okta user profile. For SSR only. Invalid if used with a Progressive Profile response. * &#x60;com.okta.action.update&#x60;: Allow or deny the user&#39;s registration. * &#x60;com.okta.user.progressive.profile.update&#x60;: Change attribute values in the user&#39;s Okta Progressive Profile. | [optional] 
**Error** | Pointer to [**RegistrationResponseError**](RegistrationResponseError.md) |  | [optional] 

## Methods

### NewRegistrationResponse

`func NewRegistrationResponse() *RegistrationResponse`

NewRegistrationResponse instantiates a new RegistrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationResponseWithDefaults

`func NewRegistrationResponseWithDefaults() *RegistrationResponse`

NewRegistrationResponseWithDefaults instantiates a new RegistrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *RegistrationResponse) GetCommands() []RegistrationResponseCommandsInner`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *RegistrationResponse) GetCommandsOk() (*[]RegistrationResponseCommandsInner, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *RegistrationResponse) SetCommands(v []RegistrationResponseCommandsInner)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *RegistrationResponse) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetError

`func (o *RegistrationResponse) GetError() RegistrationResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RegistrationResponse) GetErrorOk() (*RegistrationResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RegistrationResponse) SetError(v RegistrationResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *RegistrationResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


