# U2f

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientData** | Pointer to **string** | Base64-encoded client data from the U2F token | [optional] 
**RegistrationData** | Pointer to **string** | Base64-encoded registration data from the U2F token | [optional] 

## Methods

### NewU2f

`func NewU2f() *U2f`

NewU2f instantiates a new U2f object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewU2fWithDefaults

`func NewU2fWithDefaults() *U2f`

NewU2fWithDefaults instantiates a new U2f object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientData

`func (o *U2f) GetClientData() string`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *U2f) GetClientDataOk() (*string, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *U2f) SetClientData(v string)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *U2f) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetRegistrationData

`func (o *U2f) GetRegistrationData() string`

GetRegistrationData returns the RegistrationData field if non-nil, zero value otherwise.

### GetRegistrationDataOk

`func (o *U2f) GetRegistrationDataOk() (*string, bool)`

GetRegistrationDataOk returns a tuple with the RegistrationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationData

`func (o *U2f) SetRegistrationData(v string)`

SetRegistrationData sets RegistrationData field to given value.

### HasRegistrationData

`func (o *U2f) HasRegistrationData() bool`

HasRegistrationData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


