# OpenIdConnectApplicationSettingsRefreshToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leeway** | Pointer to **int32** |  | [optional] 
**RotationType** | Pointer to [**OpenIdConnectRefreshTokenRotationType**](OpenIdConnectRefreshTokenRotationType.md) |  | [optional] 

## Methods

### NewOpenIdConnectApplicationSettingsRefreshToken

`func NewOpenIdConnectApplicationSettingsRefreshToken() *OpenIdConnectApplicationSettingsRefreshToken`

NewOpenIdConnectApplicationSettingsRefreshToken instantiates a new OpenIdConnectApplicationSettingsRefreshToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationSettingsRefreshTokenWithDefaults

`func NewOpenIdConnectApplicationSettingsRefreshTokenWithDefaults() *OpenIdConnectApplicationSettingsRefreshToken`

NewOpenIdConnectApplicationSettingsRefreshTokenWithDefaults instantiates a new OpenIdConnectApplicationSettingsRefreshToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeeway

`func (o *OpenIdConnectApplicationSettingsRefreshToken) GetLeeway() int32`

GetLeeway returns the Leeway field if non-nil, zero value otherwise.

### GetLeewayOk

`func (o *OpenIdConnectApplicationSettingsRefreshToken) GetLeewayOk() (*int32, bool)`

GetLeewayOk returns a tuple with the Leeway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeeway

`func (o *OpenIdConnectApplicationSettingsRefreshToken) SetLeeway(v int32)`

SetLeeway sets Leeway field to given value.

### HasLeeway

`func (o *OpenIdConnectApplicationSettingsRefreshToken) HasLeeway() bool`

HasLeeway returns a boolean if a field has been set.

### GetRotationType

`func (o *OpenIdConnectApplicationSettingsRefreshToken) GetRotationType() OpenIdConnectRefreshTokenRotationType`

GetRotationType returns the RotationType field if non-nil, zero value otherwise.

### GetRotationTypeOk

`func (o *OpenIdConnectApplicationSettingsRefreshToken) GetRotationTypeOk() (*OpenIdConnectRefreshTokenRotationType, bool)`

GetRotationTypeOk returns a tuple with the RotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationType

`func (o *OpenIdConnectApplicationSettingsRefreshToken) SetRotationType(v OpenIdConnectRefreshTokenRotationType)`

SetRotationType sets RotationType field to given value.

### HasRotationType

`func (o *OpenIdConnectApplicationSettingsRefreshToken) HasRotationType() bool`

HasRotationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


