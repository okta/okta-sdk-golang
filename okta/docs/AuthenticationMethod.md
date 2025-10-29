# AuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareProtection** | Pointer to **string** | Indicates if any secrets or private keys used during authentication must be hardware protected and not exportable. This property is only set for &#x60;POSSESSION&#x60; constraints. | [optional] [default to "OPTIONAL"]
**Id** | Pointer to **string** | An ID that identifies the authenticator | [optional] 
**Key** | **string** | A label that identifies the authenticator | 
**Method** | **string** | Specifies the method used for the authenticator | 
**PhishingResistant** | Pointer to **string** | Indicates if phishing-resistant Factors are required. This property is only set for &#x60;POSSESSION&#x60; constraints | [optional] [default to "OPTIONAL"]
**UserVerification** | Pointer to **string** | Indicates if a user is required to be verified with a verification method. | [optional] [default to "OPTIONAL"]
**UserVerificationMethods** | Pointer to **[]string** | Indicates which methods can be used for user verification. &#x60;userVerificationMethods&#x60; can only be used when &#x60;userVerification&#x60; is &#x60;REQUIRED&#x60;. &#x60;BIOMETRICS&#x60; is currently the only supported method. | [optional] 

## Methods

### NewAuthenticationMethod

`func NewAuthenticationMethod(key string, method string, ) *AuthenticationMethod`

NewAuthenticationMethod instantiates a new AuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodWithDefaults

`func NewAuthenticationMethodWithDefaults() *AuthenticationMethod`

NewAuthenticationMethodWithDefaults instantiates a new AuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareProtection

`func (o *AuthenticationMethod) GetHardwareProtection() string`

GetHardwareProtection returns the HardwareProtection field if non-nil, zero value otherwise.

### GetHardwareProtectionOk

`func (o *AuthenticationMethod) GetHardwareProtectionOk() (*string, bool)`

GetHardwareProtectionOk returns a tuple with the HardwareProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareProtection

`func (o *AuthenticationMethod) SetHardwareProtection(v string)`

SetHardwareProtection sets HardwareProtection field to given value.

### HasHardwareProtection

`func (o *AuthenticationMethod) HasHardwareProtection() bool`

HasHardwareProtection returns a boolean if a field has been set.

### GetId

`func (o *AuthenticationMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AuthenticationMethod) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthenticationMethod) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthenticationMethod) SetKey(v string)`

SetKey sets Key field to given value.


### GetMethod

`func (o *AuthenticationMethod) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthenticationMethod) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthenticationMethod) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPhishingResistant

`func (o *AuthenticationMethod) GetPhishingResistant() string`

GetPhishingResistant returns the PhishingResistant field if non-nil, zero value otherwise.

### GetPhishingResistantOk

`func (o *AuthenticationMethod) GetPhishingResistantOk() (*string, bool)`

GetPhishingResistantOk returns a tuple with the PhishingResistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhishingResistant

`func (o *AuthenticationMethod) SetPhishingResistant(v string)`

SetPhishingResistant sets PhishingResistant field to given value.

### HasPhishingResistant

`func (o *AuthenticationMethod) HasPhishingResistant() bool`

HasPhishingResistant returns a boolean if a field has been set.

### GetUserVerification

`func (o *AuthenticationMethod) GetUserVerification() string`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticationMethod) GetUserVerificationOk() (*string, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticationMethod) SetUserVerification(v string)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticationMethod) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetUserVerificationMethods

`func (o *AuthenticationMethod) GetUserVerificationMethods() []string`

GetUserVerificationMethods returns the UserVerificationMethods field if non-nil, zero value otherwise.

### GetUserVerificationMethodsOk

`func (o *AuthenticationMethod) GetUserVerificationMethodsOk() (*[]string, bool)`

GetUserVerificationMethodsOk returns a tuple with the UserVerificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerificationMethods

`func (o *AuthenticationMethod) SetUserVerificationMethods(v []string)`

SetUserVerificationMethods sets UserVerificationMethods field to given value.

### HasUserVerificationMethods

`func (o *AuthenticationMethod) HasUserVerificationMethods() bool`

HasUserVerificationMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


