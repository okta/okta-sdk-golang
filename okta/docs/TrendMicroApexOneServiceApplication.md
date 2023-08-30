# TrendMicroApexOneServiceApplication

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
**Name** | **string** |  | [default to "trendmicroapexoneservice"]
**Profile** | Pointer to **map[string]map[string]interface{}** | Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps) | [optional] 
**Settings** | [**TrendMicroApexOneServiceApplicationSettings**](TrendMicroApexOneServiceApplicationSettings.md) |  | 
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to SAML_2_0]
**Status** | Pointer to [**ApplicationLifecycleStatus**](ApplicationLifecycleStatus.md) |  | [optional] 
**Visibility** | Pointer to [**ApplicationVisibility**](ApplicationVisibility.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**ApplicationLinks**](ApplicationLinks.md) |  | [optional] 

## Methods

### NewTrendMicroApexOneServiceApplication

`func NewTrendMicroApexOneServiceApplication(label string, name string, settings TrendMicroApexOneServiceApplicationSettings, ) *TrendMicroApexOneServiceApplication`

NewTrendMicroApexOneServiceApplication instantiates a new TrendMicroApexOneServiceApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendMicroApexOneServiceApplicationWithDefaults

`func NewTrendMicroApexOneServiceApplicationWithDefaults() *TrendMicroApexOneServiceApplication`

NewTrendMicroApexOneServiceApplicationWithDefaults instantiates a new TrendMicroApexOneServiceApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *TrendMicroApexOneServiceApplication) GetAccessibility() ApplicationAccessibility`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *TrendMicroApexOneServiceApplication) GetAccessibilityOk() (*ApplicationAccessibility, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *TrendMicroApexOneServiceApplication) SetAccessibility(v ApplicationAccessibility)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *TrendMicroApexOneServiceApplication) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCreated

`func (o *TrendMicroApexOneServiceApplication) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrendMicroApexOneServiceApplication) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrendMicroApexOneServiceApplication) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrendMicroApexOneServiceApplication) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *TrendMicroApexOneServiceApplication) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *TrendMicroApexOneServiceApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *TrendMicroApexOneServiceApplication) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *TrendMicroApexOneServiceApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetFeatures

`func (o *TrendMicroApexOneServiceApplication) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TrendMicroApexOneServiceApplication) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TrendMicroApexOneServiceApplication) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *TrendMicroApexOneServiceApplication) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *TrendMicroApexOneServiceApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrendMicroApexOneServiceApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrendMicroApexOneServiceApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrendMicroApexOneServiceApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *TrendMicroApexOneServiceApplication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TrendMicroApexOneServiceApplication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TrendMicroApexOneServiceApplication) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUpdated

`func (o *TrendMicroApexOneServiceApplication) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TrendMicroApexOneServiceApplication) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TrendMicroApexOneServiceApplication) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TrendMicroApexOneServiceApplication) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLicensing

`func (o *TrendMicroApexOneServiceApplication) GetLicensing() ApplicationLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *TrendMicroApexOneServiceApplication) GetLicensingOk() (*ApplicationLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *TrendMicroApexOneServiceApplication) SetLicensing(v ApplicationLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *TrendMicroApexOneServiceApplication) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetName

`func (o *TrendMicroApexOneServiceApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrendMicroApexOneServiceApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrendMicroApexOneServiceApplication) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *TrendMicroApexOneServiceApplication) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *TrendMicroApexOneServiceApplication) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *TrendMicroApexOneServiceApplication) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *TrendMicroApexOneServiceApplication) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSettings

`func (o *TrendMicroApexOneServiceApplication) GetSettings() TrendMicroApexOneServiceApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TrendMicroApexOneServiceApplication) GetSettingsOk() (*TrendMicroApexOneServiceApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TrendMicroApexOneServiceApplication) SetSettings(v TrendMicroApexOneServiceApplicationSettings)`

SetSettings sets Settings field to given value.


### GetSignOnMode

`func (o *TrendMicroApexOneServiceApplication) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *TrendMicroApexOneServiceApplication) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *TrendMicroApexOneServiceApplication) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *TrendMicroApexOneServiceApplication) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *TrendMicroApexOneServiceApplication) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *TrendMicroApexOneServiceApplication) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetStatus

`func (o *TrendMicroApexOneServiceApplication) GetStatus() ApplicationLifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrendMicroApexOneServiceApplication) GetStatusOk() (*ApplicationLifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrendMicroApexOneServiceApplication) SetStatus(v ApplicationLifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrendMicroApexOneServiceApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *TrendMicroApexOneServiceApplication) GetVisibility() ApplicationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TrendMicroApexOneServiceApplication) GetVisibilityOk() (*ApplicationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TrendMicroApexOneServiceApplication) SetVisibility(v ApplicationVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TrendMicroApexOneServiceApplication) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEmbedded

`func (o *TrendMicroApexOneServiceApplication) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *TrendMicroApexOneServiceApplication) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *TrendMicroApexOneServiceApplication) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *TrendMicroApexOneServiceApplication) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *TrendMicroApexOneServiceApplication) GetLinks() ApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TrendMicroApexOneServiceApplication) GetLinksOk() (*ApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TrendMicroApexOneServiceApplication) SetLinks(v ApplicationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TrendMicroApexOneServiceApplication) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


