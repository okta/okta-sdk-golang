# ApplicationCredentialsUsernameTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushStatus** | Pointer to **string** | Determines if the username is pushed to the app on updates for CUSTOM &#x60;type&#x60; | [optional] 
**Template** | Pointer to **string** | Mapping expression used to generate usernames.  The following are supported mapping expressions that are used with the &#x60;BUILT_IN&#x60; template type:  | Name                            | Template Expression                            | | ------------------------------- | ---------------------------------------------- | | AD Employee ID                  | &#x60;${source.employeeID}&#x60;                         | | AD SAM Account Name             | &#x60;${source.samAccountName}&#x60;                     | | AD SAM Account Name (lowercase) | &#x60;${fn:toLowerCase(source.samAccountName)}&#x60;     | | AD User Principal Name          | &#x60;${source.userName}&#x60;                           | | AD User Principal Name prefix   | &#x60;${fn:substringBefore(source.userName, \&quot;@\&quot;)}&#x60;  | | Email                           | &#x60;${source.email}&#x60;                              | | Email (lowercase)               | &#x60;${fn:toLowerCase(source.email)}&#x60;              | | Email prefix                    | &#x60;${fn:substringBefore(source.email, \&quot;@\&quot;)}&#x60;     | | LDAP UID + custom suffix        | &#x60;${source.userName}${instance.userSuffix}&#x60;     | | Okta username                   | &#x60;${source.login}&#x60;                              | | Okta username prefix            | &#x60;${fn:substringBefore(source.login, \&quot;@\&quot;)}&#x60;     | | [optional] [default to "${source.login}"]
**Type** | Pointer to **string** | Type of mapping expression. Empty string is allowed. | [optional] [default to "BUILT_IN"]
**UserSuffix** | Pointer to **string** | An optional suffix appended to usernames for &#x60;BUILT_IN&#x60; mapping expressions | [optional] 

## Methods

### NewApplicationCredentialsUsernameTemplate

`func NewApplicationCredentialsUsernameTemplate() *ApplicationCredentialsUsernameTemplate`

NewApplicationCredentialsUsernameTemplate instantiates a new ApplicationCredentialsUsernameTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCredentialsUsernameTemplateWithDefaults

`func NewApplicationCredentialsUsernameTemplateWithDefaults() *ApplicationCredentialsUsernameTemplate`

NewApplicationCredentialsUsernameTemplateWithDefaults instantiates a new ApplicationCredentialsUsernameTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushStatus

`func (o *ApplicationCredentialsUsernameTemplate) GetPushStatus() string`

GetPushStatus returns the PushStatus field if non-nil, zero value otherwise.

### GetPushStatusOk

`func (o *ApplicationCredentialsUsernameTemplate) GetPushStatusOk() (*string, bool)`

GetPushStatusOk returns a tuple with the PushStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushStatus

`func (o *ApplicationCredentialsUsernameTemplate) SetPushStatus(v string)`

SetPushStatus sets PushStatus field to given value.

### HasPushStatus

`func (o *ApplicationCredentialsUsernameTemplate) HasPushStatus() bool`

HasPushStatus returns a boolean if a field has been set.

### GetTemplate

`func (o *ApplicationCredentialsUsernameTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ApplicationCredentialsUsernameTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ApplicationCredentialsUsernameTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ApplicationCredentialsUsernameTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *ApplicationCredentialsUsernameTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCredentialsUsernameTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCredentialsUsernameTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationCredentialsUsernameTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserSuffix

`func (o *ApplicationCredentialsUsernameTemplate) GetUserSuffix() string`

GetUserSuffix returns the UserSuffix field if non-nil, zero value otherwise.

### GetUserSuffixOk

`func (o *ApplicationCredentialsUsernameTemplate) GetUserSuffixOk() (*string, bool)`

GetUserSuffixOk returns a tuple with the UserSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSuffix

`func (o *ApplicationCredentialsUsernameTemplate) SetUserSuffix(v string)`

SetUserSuffix sets UserSuffix field to given value.

### HasUserSuffix

`func (o *ApplicationCredentialsUsernameTemplate) HasUserSuffix() bool`

HasUserSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


