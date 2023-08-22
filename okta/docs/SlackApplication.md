# SlackApplication

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
**Name** | **string** |  | [default to "slack"]
**Profile** | Pointer to **map[string]map[string]interface{}** | Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps) | [optional] 
**Settings** | [**SlackApplicationSettings**](SlackApplicationSettings.md) |  | 
**SignOnMode** | Pointer to **interface{}** |  | [optional] [default to SAML_2_0]
**Status** | Pointer to [**ApplicationLifecycleStatus**](ApplicationLifecycleStatus.md) |  | [optional] 
**Visibility** | Pointer to [**ApplicationVisibility**](ApplicationVisibility.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**ApplicationLinks**](ApplicationLinks.md) |  | [optional] 

## Methods

### NewSlackApplication

`func NewSlackApplication(label string, name string, settings SlackApplicationSettings, ) *SlackApplication`

NewSlackApplication instantiates a new SlackApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackApplicationWithDefaults

`func NewSlackApplicationWithDefaults() *SlackApplication`

NewSlackApplicationWithDefaults instantiates a new SlackApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *SlackApplication) GetAccessibility() ApplicationAccessibility`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *SlackApplication) GetAccessibilityOk() (*ApplicationAccessibility, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *SlackApplication) SetAccessibility(v ApplicationAccessibility)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *SlackApplication) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCreated

`func (o *SlackApplication) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SlackApplication) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SlackApplication) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SlackApplication) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *SlackApplication) GetCredentials() SchemeApplicationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SlackApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SlackApplication) SetCredentials(v SchemeApplicationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SlackApplication) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetFeatures

`func (o *SlackApplication) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *SlackApplication) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *SlackApplication) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *SlackApplication) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *SlackApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlackApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlackApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SlackApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *SlackApplication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SlackApplication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SlackApplication) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUpdated

`func (o *SlackApplication) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SlackApplication) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SlackApplication) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SlackApplication) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLicensing

`func (o *SlackApplication) GetLicensing() ApplicationLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *SlackApplication) GetLicensingOk() (*ApplicationLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *SlackApplication) SetLicensing(v ApplicationLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *SlackApplication) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetName

`func (o *SlackApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlackApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlackApplication) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *SlackApplication) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SlackApplication) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SlackApplication) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SlackApplication) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSettings

`func (o *SlackApplication) GetSettings() SlackApplicationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SlackApplication) GetSettingsOk() (*SlackApplicationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SlackApplication) SetSettings(v SlackApplicationSettings)`

SetSettings sets Settings field to given value.


### GetSignOnMode

`func (o *SlackApplication) GetSignOnMode() interface{}`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *SlackApplication) GetSignOnModeOk() (*interface{}, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *SlackApplication) SetSignOnMode(v interface{})`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *SlackApplication) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### SetSignOnModeNil

`func (o *SlackApplication) SetSignOnModeNil(b bool)`

 SetSignOnModeNil sets the value for SignOnMode to be an explicit nil

### UnsetSignOnMode
`func (o *SlackApplication) UnsetSignOnMode()`

UnsetSignOnMode ensures that no value is present for SignOnMode, not even an explicit nil
### GetStatus

`func (o *SlackApplication) GetStatus() ApplicationLifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SlackApplication) GetStatusOk() (*ApplicationLifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SlackApplication) SetStatus(v ApplicationLifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SlackApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *SlackApplication) GetVisibility() ApplicationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SlackApplication) GetVisibilityOk() (*ApplicationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SlackApplication) SetVisibility(v ApplicationVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SlackApplication) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEmbedded

`func (o *SlackApplication) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SlackApplication) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SlackApplication) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SlackApplication) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *SlackApplication) GetLinks() ApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SlackApplication) GetLinksOk() (*ApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SlackApplication) SetLinks(v ApplicationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SlackApplication) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


