# TestInfoTestAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The sign-in URL to a test instance of your app | 
**Username** | **string** | The username for your app admin account | 
**Password** | **string** | The password for your app admin account | 
**Instructions** | Pointer to **string** | Additional instructions to test the app integration, including instructions for obtaining test accounts | [optional] 

## Methods

### NewTestInfoTestAccount

`func NewTestInfoTestAccount(url string, username string, password string, ) *TestInfoTestAccount`

NewTestInfoTestAccount instantiates a new TestInfoTestAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestInfoTestAccountWithDefaults

`func NewTestInfoTestAccountWithDefaults() *TestInfoTestAccount`

NewTestInfoTestAccountWithDefaults instantiates a new TestInfoTestAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *TestInfoTestAccount) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TestInfoTestAccount) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TestInfoTestAccount) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *TestInfoTestAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TestInfoTestAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TestInfoTestAccount) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *TestInfoTestAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TestInfoTestAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TestInfoTestAccount) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetInstructions

`func (o *TestInfoTestAccount) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *TestInfoTestAccount) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *TestInfoTestAccount) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *TestInfoTestAccount) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


