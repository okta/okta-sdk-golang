# Org2OrgApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to [**ApplicationAccessibility**](ApplicationAccessibility.md) |  | [optional] 
**Credentials** | Pointer to [**SchemeApplicationCredentials**](SchemeApplicationCredentials.md) |  | [optional] 
**Label** | **string** | User-defined display name for app | 
**Licensing** | Pointer to [**ApplicationLicensing**](ApplicationLicensing.md) |  | [optional] 
**Name** | **string** |  | 
**Profile** | Pointer to **map[string]map[string]interface{}** | Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps) | [optional] 
**SignOnMode** | Pointer to **string** |  | [optional] [default to "SAML_2_0"]
**Status** | Pointer to **string** | App instance status | [optional] [readonly] 
**Visibility** | Pointer to [**ApplicationVisibility**](ApplicationVisibility.md) |  | [optional] 
**Settings** | [**Org2OrgApplicationSettings**](Org2OrgApplicationSettings.md) |  | 

## Methods

### NewOrg2OrgApplication

`func NewOrg2OrgApplication(label string, name string, settings Org2OrgApplicationSettings, ) *Org2OrgApplication`

NewOrg2OrgApplication instantiates a new Org2OrgApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrg2OrgApplicationWithDefaults

`func NewOrg2OrgApplicationWithDefaults() *Org2OrgApplication`

NewOrg2OrgApplicationWithDefaults instantiates a new Org2OrgApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *Org2OrgApplication) GetAccessibility() ApplicationAccessibility`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *Org2OrgApplication) GetAccessibilityOk() (*ApplicationAccessibility, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *Org2OrgApplication) SetAccessibility(v ApplicationAccessibility)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *Org2OrgApplication) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCredentials

`func (o *Org2OrgApplication) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Org2OrgApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Org2OrgApplication) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Org2OrgApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetLabel

`func (o *Org2OrgApplication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Org2OrgApplication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Org2OrgApplication) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLicensing

`func (o *Org2OrgApplication) GetLicensing() ApplicationLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *Org2OrgApplication) GetLicensingOk() (*ApplicationLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *Org2OrgApplication) SetLicensing(v ApplicationLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *Org2OrgApplication) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetName

`func (o *Org2OrgApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Org2OrgApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Org2OrgApplication) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *Org2OrgApplication) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Org2OrgApplication) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Org2OrgApplication) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Org2OrgApplication) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSignOnMode

`func (o *Org2OrgApplication) GetSignOnMode() string`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *Org2OrgApplication) GetSignOnModeOk() (*string, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *Org2OrgApplication) SetSignOnMode(v string)`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *Org2OrgApplication) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### GetStatus

`func (o *Org2OrgApplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Org2OrgApplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Org2OrgApplication) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Org2OrgApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *Org2OrgApplication) GetVisibility() ApplicationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Org2OrgApplication) GetVisibilityOk() (*ApplicationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Org2OrgApplication) SetVisibility(v ApplicationVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Org2OrgApplication) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSettings

`func (o *Org2OrgApplication) GetSettings() Org2OrgApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Org2OrgApplication) GetSettingsOk() (*Org2OrgApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Org2OrgApplication) SetSettings(v Org2OrgApplicationSettings)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


