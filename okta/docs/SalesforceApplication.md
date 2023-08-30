# SalesforceApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to [**ApplicationAccessibility**](ApplicationAccessibility.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the Application object was created | [optional] [readonly] 
**Credentials** | Pointer to [**SchemeApplicationCredentials**](SchemeApplicationCredentials.md) |  | [optional] 
**Features** | Pointer to **[]string** | Enabled app features | [optional] 
**Id** | Pointer to **string** | Unique ID for the app instance | [optional] [readonly] 
**Label** | **string** | User-defined display name for app | 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Application object was last updated | [optional] [readonly] 
**Licensing** | Pointer to [**ApplicationLicensing**](ApplicationLicensing.md) |  | [optional] 
**Name** | **string** |  | [default to "salesforce"]
**Profile** | Pointer to **map[string]map[string]interface{}** | Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps) | [optional] 
**Settings** | [**SalesforceApplicationSettings**](SalesforceApplicationSettings.md) |  | 
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to BROWSER_PLUGIN]
**Status** | Pointer to [**ApplicationLifecycleStatus**](ApplicationLifecycleStatus.md) |  | [optional] 
**Visibility** | Pointer to [**ApplicationVisibility**](ApplicationVisibility.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**ApplicationLinks**](ApplicationLinks.md) |  | [optional] 

## Methods

### NewSalesforceApplication

`func NewSalesforceApplication(label string, name string, settings SalesforceApplicationSettings, ) *SalesforceApplication`

NewSalesforceApplication instantiates a new SalesforceApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceApplicationWithDefaults

`func NewSalesforceApplicationWithDefaults() *SalesforceApplication`

NewSalesforceApplicationWithDefaults instantiates a new SalesforceApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *SalesforceApplication) GetAccessibility() ApplicationAccessibility`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *SalesforceApplication) GetAccessibilityOk() (*ApplicationAccessibility, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *SalesforceApplication) SetAccessibility(v ApplicationAccessibility)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *SalesforceApplication) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCreated

`func (o *SalesforceApplication) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SalesforceApplication) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SalesforceApplication) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SalesforceApplication) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *SalesforceApplication) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SalesforceApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SalesforceApplication) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SalesforceApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetFeatures

`func (o *SalesforceApplication) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *SalesforceApplication) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *SalesforceApplication) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *SalesforceApplication) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *SalesforceApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesforceApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesforceApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SalesforceApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *SalesforceApplication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SalesforceApplication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SalesforceApplication) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUpdated

`func (o *SalesforceApplication) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SalesforceApplication) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SalesforceApplication) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SalesforceApplication) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLicensing

`func (o *SalesforceApplication) GetLicensing() ApplicationLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *SalesforceApplication) GetLicensingOk() (*ApplicationLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *SalesforceApplication) SetLicensing(v ApplicationLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *SalesforceApplication) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetName

`func (o *SalesforceApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SalesforceApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SalesforceApplication) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *SalesforceApplication) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SalesforceApplication) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SalesforceApplication) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SalesforceApplication) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSettings

`func (o *SalesforceApplication) GetSettings() SalesforceApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SalesforceApplication) GetSettingsOk() (*SalesforceApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SalesforceApplication) SetSettings(v SalesforceApplicationSettings)`

SetSettings sets Settings field to given value.


### GetSignOnMode

`func (o *SalesforceApplication) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *SalesforceApplication) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *SalesforceApplication) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *SalesforceApplication) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *SalesforceApplication) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *SalesforceApplication) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetStatus

`func (o *SalesforceApplication) GetStatus() ApplicationLifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SalesforceApplication) GetStatusOk() (*ApplicationLifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SalesforceApplication) SetStatus(v ApplicationLifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SalesforceApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *SalesforceApplication) GetVisibility() ApplicationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SalesforceApplication) GetVisibilityOk() (*ApplicationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SalesforceApplication) SetVisibility(v ApplicationVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SalesforceApplication) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEmbedded

`func (o *SalesforceApplication) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SalesforceApplication) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SalesforceApplication) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SalesforceApplication) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *SalesforceApplication) GetLinks() ApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SalesforceApplication) GetLinksOk() (*ApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SalesforceApplication) SetLinks(v ApplicationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SalesforceApplication) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


