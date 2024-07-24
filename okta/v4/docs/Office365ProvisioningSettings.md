# Office365ProvisioningSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPassword** | **string** | Microsoft Office 365 global administrator password | 
**AdminUsername** | **string** | Microsoft Office 365 global administrator username | 

## Methods

### NewOffice365ProvisioningSettings

`func NewOffice365ProvisioningSettings(adminPassword string, adminUsername string, ) *Office365ProvisioningSettings`

NewOffice365ProvisioningSettings instantiates a new Office365ProvisioningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ProvisioningSettingsWithDefaults

`func NewOffice365ProvisioningSettingsWithDefaults() *Office365ProvisioningSettings`

NewOffice365ProvisioningSettingsWithDefaults instantiates a new Office365ProvisioningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPassword

`func (o *Office365ProvisioningSettings) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *Office365ProvisioningSettings) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *Office365ProvisioningSettings) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.


### GetAdminUsername

`func (o *Office365ProvisioningSettings) GetAdminUsername() string`

GetAdminUsername returns the AdminUsername field if non-nil, zero value otherwise.

### GetAdminUsernameOk

`func (o *Office365ProvisioningSettings) GetAdminUsernameOk() (*string, bool)`

GetAdminUsernameOk returns a tuple with the AdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsername

`func (o *Office365ProvisioningSettings) SetAdminUsername(v string)`

SetAdminUsername sets AdminUsername field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


