# TestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EscalationSupportContact** | **string** | An email for Okta to contact your company about your integration. This email isn&#39;t shared with customers. | 
**OidcTestConfiguration** | Pointer to [**TestInfoOidcTestConfiguration**](TestInfoOidcTestConfiguration.md) |  | [optional] 
**SamlTestConfiguration** | Pointer to [**TestInfoSamlTestConfiguration**](TestInfoSamlTestConfiguration.md) |  | [optional] 
**TestAccount** | Pointer to [**TestInfoTestAccount**](TestInfoTestAccount.md) |  | [optional] 

## Methods

### NewTestInfo

`func NewTestInfo(escalationSupportContact string, ) *TestInfo`

NewTestInfo instantiates a new TestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestInfoWithDefaults

`func NewTestInfoWithDefaults() *TestInfo`

NewTestInfoWithDefaults instantiates a new TestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEscalationSupportContact

`func (o *TestInfo) GetEscalationSupportContact() string`

GetEscalationSupportContact returns the EscalationSupportContact field if non-nil, zero value otherwise.

### GetEscalationSupportContactOk

`func (o *TestInfo) GetEscalationSupportContactOk() (*string, bool)`

GetEscalationSupportContactOk returns a tuple with the EscalationSupportContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationSupportContact

`func (o *TestInfo) SetEscalationSupportContact(v string)`

SetEscalationSupportContact sets EscalationSupportContact field to given value.


### GetOidcTestConfiguration

`func (o *TestInfo) GetOidcTestConfiguration() TestInfoOidcTestConfiguration`

GetOidcTestConfiguration returns the OidcTestConfiguration field if non-nil, zero value otherwise.

### GetOidcTestConfigurationOk

`func (o *TestInfo) GetOidcTestConfigurationOk() (*TestInfoOidcTestConfiguration, bool)`

GetOidcTestConfigurationOk returns a tuple with the OidcTestConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTestConfiguration

`func (o *TestInfo) SetOidcTestConfiguration(v TestInfoOidcTestConfiguration)`

SetOidcTestConfiguration sets OidcTestConfiguration field to given value.

### HasOidcTestConfiguration

`func (o *TestInfo) HasOidcTestConfiguration() bool`

HasOidcTestConfiguration returns a boolean if a field has been set.

### GetSamlTestConfiguration

`func (o *TestInfo) GetSamlTestConfiguration() TestInfoSamlTestConfiguration`

GetSamlTestConfiguration returns the SamlTestConfiguration field if non-nil, zero value otherwise.

### GetSamlTestConfigurationOk

`func (o *TestInfo) GetSamlTestConfigurationOk() (*TestInfoSamlTestConfiguration, bool)`

GetSamlTestConfigurationOk returns a tuple with the SamlTestConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlTestConfiguration

`func (o *TestInfo) SetSamlTestConfiguration(v TestInfoSamlTestConfiguration)`

SetSamlTestConfiguration sets SamlTestConfiguration field to given value.

### HasSamlTestConfiguration

`func (o *TestInfo) HasSamlTestConfiguration() bool`

HasSamlTestConfiguration returns a boolean if a field has been set.

### GetTestAccount

`func (o *TestInfo) GetTestAccount() TestInfoTestAccount`

GetTestAccount returns the TestAccount field if non-nil, zero value otherwise.

### GetTestAccountOk

`func (o *TestInfo) GetTestAccountOk() (*TestInfoTestAccount, bool)`

GetTestAccountOk returns a tuple with the TestAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccount

`func (o *TestInfo) SetTestAccount(v TestInfoTestAccount)`

SetTestAccount sets TestAccount field to given value.

### HasTestAccount

`func (o *TestInfo) HasTestAccount() bool`

HasTestAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


