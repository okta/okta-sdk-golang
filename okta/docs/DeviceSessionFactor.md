# DeviceSessionFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The human-readable authenticator name that&#39;s displayed.  There are three places where the display name can be retrieved from, which are evaluated in a priority order. 1. The display name is retrieved from the org&#39;s configured authenticator name. 2. The display name is retrieved from the authenticator catalog default. 3. For legacy verifications with no authenticator key, the display name is retrieved from the enrollment&#39;s own display name.  The display name is omitted only when all three are absent. | [optional] 
**Key** | Pointer to **string** | The authenticator key, &#x60;webauthn&#x60; or &#x60;okta_verify&#x60;, for example. This property is omitted for legacy verifications that carry no authenticator key. | [optional] 
**LastUsedOn** | **time.Time** | Timestamp when this factor was last verified in the session | 
**Type** | **string** | The authenticator method name, &#x60;password&#x60;, &#x60;signed_nonce&#x60;, &#x60;push&#x60;, or &#x60;otp&#x60;, for example | 

## Methods

### NewDeviceSessionFactor

`func NewDeviceSessionFactor(lastUsedOn time.Time, type_ string, ) *DeviceSessionFactor`

NewDeviceSessionFactor instantiates a new DeviceSessionFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSessionFactorWithDefaults

`func NewDeviceSessionFactorWithDefaults() *DeviceSessionFactor`

NewDeviceSessionFactorWithDefaults instantiates a new DeviceSessionFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DeviceSessionFactor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceSessionFactor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceSessionFactor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeviceSessionFactor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKey

`func (o *DeviceSessionFactor) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeviceSessionFactor) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeviceSessionFactor) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DeviceSessionFactor) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUsedOn

`func (o *DeviceSessionFactor) GetLastUsedOn() time.Time`

GetLastUsedOn returns the LastUsedOn field if non-nil, zero value otherwise.

### GetLastUsedOnOk

`func (o *DeviceSessionFactor) GetLastUsedOnOk() (*time.Time, bool)`

GetLastUsedOnOk returns a tuple with the LastUsedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedOn

`func (o *DeviceSessionFactor) SetLastUsedOn(v time.Time)`

SetLastUsedOn sets LastUsedOn field to given value.


### GetType

`func (o *DeviceSessionFactor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceSessionFactor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceSessionFactor) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


