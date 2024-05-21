# SalesforceApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | **string** | Salesforce instance that you want to connect to | 
**IntegrationType** | **string** | Salesforce integration type | 
**LoginUrl** | Pointer to **string** | The Login URL specified in your Salesforce Single Sign-On settings | [optional] 
**LogoutUrl** | Pointer to **string** | Salesforce Logout URL | [optional] 

## Methods

### NewSalesforceApplicationSettingsApplication

`func NewSalesforceApplicationSettingsApplication(instanceType string, integrationType string, ) *SalesforceApplicationSettingsApplication`

NewSalesforceApplicationSettingsApplication instantiates a new SalesforceApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceApplicationSettingsApplicationWithDefaults

`func NewSalesforceApplicationSettingsApplicationWithDefaults() *SalesforceApplicationSettingsApplication`

NewSalesforceApplicationSettingsApplicationWithDefaults instantiates a new SalesforceApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *SalesforceApplicationSettingsApplication) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *SalesforceApplicationSettingsApplication) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *SalesforceApplicationSettingsApplication) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetIntegrationType

`func (o *SalesforceApplicationSettingsApplication) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *SalesforceApplicationSettingsApplication) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *SalesforceApplicationSettingsApplication) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.


### GetLoginUrl

`func (o *SalesforceApplicationSettingsApplication) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *SalesforceApplicationSettingsApplication) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *SalesforceApplicationSettingsApplication) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *SalesforceApplicationSettingsApplication) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *SalesforceApplicationSettingsApplication) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SalesforceApplicationSettingsApplication) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SalesforceApplicationSettingsApplication) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *SalesforceApplicationSettingsApplication) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


