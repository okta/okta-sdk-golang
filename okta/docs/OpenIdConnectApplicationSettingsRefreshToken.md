# OpenIdConnectApplicationSettingsRefreshToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leeway** | Pointer to **int32** | The leeway, in seconds, allowed for the OAuth 2.0 client. After the refresh token is rotated, the previous token remains valid for the specified period of time so clients can get the new token.  &gt; **Note:** A leeway of 0 doesn&#39;t necessarily mean that the previous token is immediately invalidated. The previous token is invalidated after the new token is generated and returned in the response.  | [optional] [default to 30]
**RotationType** | **string** | The refresh token rotation mode for the OAuth 2.0 client | 

## Methods

### NewOpenIdConnectApplicationSettingsRefreshToken

`func NewOpenIdConnectApplicationSettingsRefreshToken(rotationType string, ) *OpenIdConnectApplicationSettingsRefreshToken`

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

`func (o *OpenIdConnectApplicationSettingsRefreshToken) GetRotationType() string`

GetRotationType returns the RotationType field if non-nil, zero value otherwise.

### GetRotationTypeOk

`func (o *OpenIdConnectApplicationSettingsRefreshToken) GetRotationTypeOk() (*string, bool)`

GetRotationTypeOk returns a tuple with the RotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationType

`func (o *OpenIdConnectApplicationSettingsRefreshToken) SetRotationType(v string)`

SetRotationType sets RotationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


