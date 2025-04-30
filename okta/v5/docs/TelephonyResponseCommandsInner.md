# TelephonyResponseCommandsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The location where you specify the command. For the Telephony inline hook, there&#39;s only one command, &#x60;com.okta.telephony.action&#x60;. | [optional] 
**Value** | Pointer to [**[]TelephonyResponseCommandsInnerValueInner**](TelephonyResponseCommandsInnerValueInner.md) | The status of the telephony operation along with optional additional information about the provider, transaction ID and any other transaction metadata. | [optional] 

## Methods

### NewTelephonyResponseCommandsInner

`func NewTelephonyResponseCommandsInner() *TelephonyResponseCommandsInner`

NewTelephonyResponseCommandsInner instantiates a new TelephonyResponseCommandsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyResponseCommandsInnerWithDefaults

`func NewTelephonyResponseCommandsInnerWithDefaults() *TelephonyResponseCommandsInner`

NewTelephonyResponseCommandsInnerWithDefaults instantiates a new TelephonyResponseCommandsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelephonyResponseCommandsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelephonyResponseCommandsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelephonyResponseCommandsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TelephonyResponseCommandsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *TelephonyResponseCommandsInner) GetValue() []TelephonyResponseCommandsInnerValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TelephonyResponseCommandsInner) GetValueOk() (*[]TelephonyResponseCommandsInnerValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TelephonyResponseCommandsInner) SetValue(v []TelephonyResponseCommandsInnerValueInner)`

SetValue sets Value field to given value.

### HasValue

`func (o *TelephonyResponseCommandsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


