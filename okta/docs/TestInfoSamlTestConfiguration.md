# TestInfoSamlTestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Idp** | Pointer to **bool** | Indicates if your integration supports IdP-initiated sign-in | [optional] 
**Sp** | Pointer to **bool** | Indicates if your integration supports SP-initiated sign-in | [optional] 
**Jit** | Pointer to **bool** | Indicates if your integration supports Just-In-Time (JIT) provisioning | [optional] 
**SpInitiateUrl** | **string** | URL for SP-initiated sign-in flows (required if &#x60;sp &#x3D; true&#x60;) | 
**SpInitiateDescription** | Pointer to **string** | Instructions on how to sign in to your app using the SP-initiated flow (required if &#x60;sp &#x3D; true&#x60;) | [optional] 

## Methods

### NewTestInfoSamlTestConfiguration

`func NewTestInfoSamlTestConfiguration(spInitiateUrl string, ) *TestInfoSamlTestConfiguration`

NewTestInfoSamlTestConfiguration instantiates a new TestInfoSamlTestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestInfoSamlTestConfigurationWithDefaults

`func NewTestInfoSamlTestConfigurationWithDefaults() *TestInfoSamlTestConfiguration`

NewTestInfoSamlTestConfigurationWithDefaults instantiates a new TestInfoSamlTestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdp

`func (o *TestInfoSamlTestConfiguration) GetIdp() bool`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *TestInfoSamlTestConfiguration) GetIdpOk() (*bool, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *TestInfoSamlTestConfiguration) SetIdp(v bool)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *TestInfoSamlTestConfiguration) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetSp

`func (o *TestInfoSamlTestConfiguration) GetSp() bool`

GetSp returns the Sp field if non-nil, zero value otherwise.

### GetSpOk

`func (o *TestInfoSamlTestConfiguration) GetSpOk() (*bool, bool)`

GetSpOk returns a tuple with the Sp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSp

`func (o *TestInfoSamlTestConfiguration) SetSp(v bool)`

SetSp sets Sp field to given value.

### HasSp

`func (o *TestInfoSamlTestConfiguration) HasSp() bool`

HasSp returns a boolean if a field has been set.

### GetJit

`func (o *TestInfoSamlTestConfiguration) GetJit() bool`

GetJit returns the Jit field if non-nil, zero value otherwise.

### GetJitOk

`func (o *TestInfoSamlTestConfiguration) GetJitOk() (*bool, bool)`

GetJitOk returns a tuple with the Jit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJit

`func (o *TestInfoSamlTestConfiguration) SetJit(v bool)`

SetJit sets Jit field to given value.

### HasJit

`func (o *TestInfoSamlTestConfiguration) HasJit() bool`

HasJit returns a boolean if a field has been set.

### GetSpInitiateUrl

`func (o *TestInfoSamlTestConfiguration) GetSpInitiateUrl() string`

GetSpInitiateUrl returns the SpInitiateUrl field if non-nil, zero value otherwise.

### GetSpInitiateUrlOk

`func (o *TestInfoSamlTestConfiguration) GetSpInitiateUrlOk() (*string, bool)`

GetSpInitiateUrlOk returns a tuple with the SpInitiateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiateUrl

`func (o *TestInfoSamlTestConfiguration) SetSpInitiateUrl(v string)`

SetSpInitiateUrl sets SpInitiateUrl field to given value.


### GetSpInitiateDescription

`func (o *TestInfoSamlTestConfiguration) GetSpInitiateDescription() string`

GetSpInitiateDescription returns the SpInitiateDescription field if non-nil, zero value otherwise.

### GetSpInitiateDescriptionOk

`func (o *TestInfoSamlTestConfiguration) GetSpInitiateDescriptionOk() (*string, bool)`

GetSpInitiateDescriptionOk returns a tuple with the SpInitiateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiateDescription

`func (o *TestInfoSamlTestConfiguration) SetSpInitiateDescription(v string)`

SetSpInitiateDescription sets SpInitiateDescription field to given value.

### HasSpInitiateDescription

`func (o *TestInfoSamlTestConfiguration) HasSpInitiateDescription() bool`

HasSpInitiateDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


