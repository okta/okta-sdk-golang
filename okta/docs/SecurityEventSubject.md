# SecurityEventSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **map[string]interface{}** | The device involved with the event | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | The tenant involved with the event | [optional] 
**User** | Pointer to **map[string]interface{}** | The user involved with the event | [optional] 

## Methods

### NewSecurityEventSubject

`func NewSecurityEventSubject() *SecurityEventSubject`

NewSecurityEventSubject instantiates a new SecurityEventSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventSubjectWithDefaults

`func NewSecurityEventSubjectWithDefaults() *SecurityEventSubject`

NewSecurityEventSubjectWithDefaults instantiates a new SecurityEventSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *SecurityEventSubject) GetDevice() map[string]interface{}`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SecurityEventSubject) GetDeviceOk() (*map[string]interface{}, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SecurityEventSubject) SetDevice(v map[string]interface{})`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SecurityEventSubject) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetTenant

`func (o *SecurityEventSubject) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *SecurityEventSubject) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *SecurityEventSubject) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *SecurityEventSubject) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUser

`func (o *SecurityEventSubject) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SecurityEventSubject) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SecurityEventSubject) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *SecurityEventSubject) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


