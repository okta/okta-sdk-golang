# BASICSMTPAUTH

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password of the user account that&#39;s used to sign in to your SMTP server | [optional] 

## Methods

### NewBASICSMTPAUTH

`func NewBASICSMTPAUTH() *BASICSMTPAUTH`

NewBASICSMTPAUTH instantiates a new BASICSMTPAUTH object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBASICSMTPAUTHWithDefaults

`func NewBASICSMTPAUTHWithDefaults() *BASICSMTPAUTH`

NewBASICSMTPAUTHWithDefaults instantiates a new BASICSMTPAUTH object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *BASICSMTPAUTH) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BASICSMTPAUTH) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BASICSMTPAUTH) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BASICSMTPAUTH) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


