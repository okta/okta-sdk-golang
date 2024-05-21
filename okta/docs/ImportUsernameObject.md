# ImportUsernameObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserNameExpression** | Pointer to **string** | For &#x60;usernameFormat&#x3D;CUSTOM&#x60;, specifies the Okta Expression Language statement for a username format that imported users use to sign in to Okta | [optional] 
**UsernameFormat** | **string** | Determines the username format when users sign in to Okta | [default to "EMAIL"]

## Methods

### NewImportUsernameObject

`func NewImportUsernameObject(usernameFormat string, ) *ImportUsernameObject`

NewImportUsernameObject instantiates a new ImportUsernameObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUsernameObjectWithDefaults

`func NewImportUsernameObjectWithDefaults() *ImportUsernameObject`

NewImportUsernameObjectWithDefaults instantiates a new ImportUsernameObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserNameExpression

`func (o *ImportUsernameObject) GetUserNameExpression() string`

GetUserNameExpression returns the UserNameExpression field if non-nil, zero value otherwise.

### GetUserNameExpressionOk

`func (o *ImportUsernameObject) GetUserNameExpressionOk() (*string, bool)`

GetUserNameExpressionOk returns a tuple with the UserNameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameExpression

`func (o *ImportUsernameObject) SetUserNameExpression(v string)`

SetUserNameExpression sets UserNameExpression field to given value.

### HasUserNameExpression

`func (o *ImportUsernameObject) HasUserNameExpression() bool`

HasUserNameExpression returns a boolean if a field has been set.

### GetUsernameFormat

`func (o *ImportUsernameObject) GetUsernameFormat() string`

GetUsernameFormat returns the UsernameFormat field if non-nil, zero value otherwise.

### GetUsernameFormatOk

`func (o *ImportUsernameObject) GetUsernameFormatOk() (*string, bool)`

GetUsernameFormatOk returns a tuple with the UsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameFormat

`func (o *ImportUsernameObject) SetUsernameFormat(v string)`

SetUsernameFormat sets UsernameFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


