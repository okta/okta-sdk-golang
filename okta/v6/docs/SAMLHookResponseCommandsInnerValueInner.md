# SAMLHookResponseCommandsInnerValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | The name of one of the supported ops: &#x60;add&#x60;:  Add a new claim to the assertion &#x60;replace&#x60;: Modify any element of the assertion  &gt; **Note:** If a response to the SAML assertion inline hook request isn&#39;t received from your external service within three seconds, a timeout occurs. In this scenario, the Okta process flow continues with the original SAML assertion returned. | [optional] 
**Path** | Pointer to **string** | Location, within the assertion, to apply the operation | [optional] 
**Value** | Pointer to [**SAMLHookResponseCommandsInnerValueInnerValue**](SAMLHookResponseCommandsInnerValueInnerValue.md) |  | [optional] 

## Methods

### NewSAMLHookResponseCommandsInnerValueInner

`func NewSAMLHookResponseCommandsInnerValueInner() *SAMLHookResponseCommandsInnerValueInner`

NewSAMLHookResponseCommandsInnerValueInner instantiates a new SAMLHookResponseCommandsInnerValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLHookResponseCommandsInnerValueInnerWithDefaults

`func NewSAMLHookResponseCommandsInnerValueInnerWithDefaults() *SAMLHookResponseCommandsInnerValueInner`

NewSAMLHookResponseCommandsInnerValueInnerWithDefaults instantiates a new SAMLHookResponseCommandsInnerValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *SAMLHookResponseCommandsInnerValueInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SAMLHookResponseCommandsInnerValueInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SAMLHookResponseCommandsInnerValueInner) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SAMLHookResponseCommandsInnerValueInner) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *SAMLHookResponseCommandsInnerValueInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SAMLHookResponseCommandsInnerValueInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SAMLHookResponseCommandsInnerValueInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SAMLHookResponseCommandsInnerValueInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *SAMLHookResponseCommandsInnerValueInner) GetValue() SAMLHookResponseCommandsInnerValueInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SAMLHookResponseCommandsInnerValueInner) GetValueOk() (*SAMLHookResponseCommandsInnerValueInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SAMLHookResponseCommandsInnerValueInner) SetValue(v SAMLHookResponseCommandsInnerValueInnerValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *SAMLHookResponseCommandsInnerValueInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


