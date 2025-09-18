# AppServiceAccountCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password associated with the service account | [optional] 
**Username** | **string** | The username associated with the service account | 

## Methods

### NewAppServiceAccountCredentials

`func NewAppServiceAccountCredentials(username string, ) *AppServiceAccountCredentials`

NewAppServiceAccountCredentials instantiates a new AppServiceAccountCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAccountCredentialsWithDefaults

`func NewAppServiceAccountCredentialsWithDefaults() *AppServiceAccountCredentials`

NewAppServiceAccountCredentialsWithDefaults instantiates a new AppServiceAccountCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *AppServiceAccountCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AppServiceAccountCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AppServiceAccountCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AppServiceAccountCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *AppServiceAccountCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AppServiceAccountCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AppServiceAccountCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


