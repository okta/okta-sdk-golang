# AndroidDeviceTrust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIntegrityLevel** | Pointer to **string** | Indicates how well a device can enforce app integrity | [optional] 
**NetworkProxyDisabled** | Pointer to **bool** | Indicates whether a device has a network proxy disabled | [optional] 
**PlayProtectVerdict** | Pointer to **string** | Indicates if Google Play Protect is enabled on the device and whether it has found known malware | [optional] 
**RequireMajorVersionUpdate** | Pointer to **bool** | Indicates whether the device needs to be on the latest major version available to the device  **Note:** This option requires an &#x60;osVersion.dynamicVersionRequirement&#x60; value to be supplied with the &#x60;osVersion.dynamicVersionRequirement.type&#x60; as either &#x60;MINIMUM&#x60; or &#x60;EXACT&#x60;.  | [optional] 
**ScreenLockComplexity** | Pointer to **string** | Indicates whether a device has a screen lock set, and the type or complexity of the screen lock  **Note:** This option requires a &#x60;screenLockType.include&#x60; value.  | [optional] 
**UsbDebuggingDisabled** | Pointer to **bool** | Indicates whether Android Debug Bridge (adb) over USB is disabled | [optional] 
**WifiSecured** | Pointer to **bool** | Indicates whether a device is on a password-protected Wi-Fi network | [optional] 

## Methods

### NewAndroidDeviceTrust

`func NewAndroidDeviceTrust() *AndroidDeviceTrust`

NewAndroidDeviceTrust instantiates a new AndroidDeviceTrust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndroidDeviceTrustWithDefaults

`func NewAndroidDeviceTrustWithDefaults() *AndroidDeviceTrust`

NewAndroidDeviceTrustWithDefaults instantiates a new AndroidDeviceTrust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIntegrityLevel

`func (o *AndroidDeviceTrust) GetDeviceIntegrityLevel() string`

GetDeviceIntegrityLevel returns the DeviceIntegrityLevel field if non-nil, zero value otherwise.

### GetDeviceIntegrityLevelOk

`func (o *AndroidDeviceTrust) GetDeviceIntegrityLevelOk() (*string, bool)`

GetDeviceIntegrityLevelOk returns a tuple with the DeviceIntegrityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIntegrityLevel

`func (o *AndroidDeviceTrust) SetDeviceIntegrityLevel(v string)`

SetDeviceIntegrityLevel sets DeviceIntegrityLevel field to given value.

### HasDeviceIntegrityLevel

`func (o *AndroidDeviceTrust) HasDeviceIntegrityLevel() bool`

HasDeviceIntegrityLevel returns a boolean if a field has been set.

### GetNetworkProxyDisabled

`func (o *AndroidDeviceTrust) GetNetworkProxyDisabled() bool`

GetNetworkProxyDisabled returns the NetworkProxyDisabled field if non-nil, zero value otherwise.

### GetNetworkProxyDisabledOk

`func (o *AndroidDeviceTrust) GetNetworkProxyDisabledOk() (*bool, bool)`

GetNetworkProxyDisabledOk returns a tuple with the NetworkProxyDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProxyDisabled

`func (o *AndroidDeviceTrust) SetNetworkProxyDisabled(v bool)`

SetNetworkProxyDisabled sets NetworkProxyDisabled field to given value.

### HasNetworkProxyDisabled

`func (o *AndroidDeviceTrust) HasNetworkProxyDisabled() bool`

HasNetworkProxyDisabled returns a boolean if a field has been set.

### GetPlayProtectVerdict

`func (o *AndroidDeviceTrust) GetPlayProtectVerdict() string`

GetPlayProtectVerdict returns the PlayProtectVerdict field if non-nil, zero value otherwise.

### GetPlayProtectVerdictOk

`func (o *AndroidDeviceTrust) GetPlayProtectVerdictOk() (*string, bool)`

GetPlayProtectVerdictOk returns a tuple with the PlayProtectVerdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayProtectVerdict

`func (o *AndroidDeviceTrust) SetPlayProtectVerdict(v string)`

SetPlayProtectVerdict sets PlayProtectVerdict field to given value.

### HasPlayProtectVerdict

`func (o *AndroidDeviceTrust) HasPlayProtectVerdict() bool`

HasPlayProtectVerdict returns a boolean if a field has been set.

### GetRequireMajorVersionUpdate

`func (o *AndroidDeviceTrust) GetRequireMajorVersionUpdate() bool`

GetRequireMajorVersionUpdate returns the RequireMajorVersionUpdate field if non-nil, zero value otherwise.

### GetRequireMajorVersionUpdateOk

`func (o *AndroidDeviceTrust) GetRequireMajorVersionUpdateOk() (*bool, bool)`

GetRequireMajorVersionUpdateOk returns a tuple with the RequireMajorVersionUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMajorVersionUpdate

`func (o *AndroidDeviceTrust) SetRequireMajorVersionUpdate(v bool)`

SetRequireMajorVersionUpdate sets RequireMajorVersionUpdate field to given value.

### HasRequireMajorVersionUpdate

`func (o *AndroidDeviceTrust) HasRequireMajorVersionUpdate() bool`

HasRequireMajorVersionUpdate returns a boolean if a field has been set.

### GetScreenLockComplexity

`func (o *AndroidDeviceTrust) GetScreenLockComplexity() string`

GetScreenLockComplexity returns the ScreenLockComplexity field if non-nil, zero value otherwise.

### GetScreenLockComplexityOk

`func (o *AndroidDeviceTrust) GetScreenLockComplexityOk() (*string, bool)`

GetScreenLockComplexityOk returns a tuple with the ScreenLockComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockComplexity

`func (o *AndroidDeviceTrust) SetScreenLockComplexity(v string)`

SetScreenLockComplexity sets ScreenLockComplexity field to given value.

### HasScreenLockComplexity

`func (o *AndroidDeviceTrust) HasScreenLockComplexity() bool`

HasScreenLockComplexity returns a boolean if a field has been set.

### GetUsbDebuggingDisabled

`func (o *AndroidDeviceTrust) GetUsbDebuggingDisabled() bool`

GetUsbDebuggingDisabled returns the UsbDebuggingDisabled field if non-nil, zero value otherwise.

### GetUsbDebuggingDisabledOk

`func (o *AndroidDeviceTrust) GetUsbDebuggingDisabledOk() (*bool, bool)`

GetUsbDebuggingDisabledOk returns a tuple with the UsbDebuggingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbDebuggingDisabled

`func (o *AndroidDeviceTrust) SetUsbDebuggingDisabled(v bool)`

SetUsbDebuggingDisabled sets UsbDebuggingDisabled field to given value.

### HasUsbDebuggingDisabled

`func (o *AndroidDeviceTrust) HasUsbDebuggingDisabled() bool`

HasUsbDebuggingDisabled returns a boolean if a field has been set.

### GetWifiSecured

`func (o *AndroidDeviceTrust) GetWifiSecured() bool`

GetWifiSecured returns the WifiSecured field if non-nil, zero value otherwise.

### GetWifiSecuredOk

`func (o *AndroidDeviceTrust) GetWifiSecuredOk() (*bool, bool)`

GetWifiSecuredOk returns a tuple with the WifiSecured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiSecured

`func (o *AndroidDeviceTrust) SetWifiSecured(v bool)`

SetWifiSecured sets WifiSecured field to given value.

### HasWifiSecured

`func (o *AndroidDeviceTrust) HasWifiSecured() bool`

HasWifiSecured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


