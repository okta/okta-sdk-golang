# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to [**ApplicationAccessibility**](ApplicationAccessibility.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Features** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Label** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Licensing** | Pointer to [**ApplicationLicensing**](ApplicationLicensing.md) |  | [optional] 
**Profile** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**SignOnMode** | Pointer to [**ApplicationSignOnMode**](ApplicationSignOnMode.md) |  | [optional] 
**Status** | Pointer to [**ApplicationLifecycleStatus**](ApplicationLifecycleStatus.md) |  | [optional] 
**Visibility** | Pointer to [**ApplicationVisibility**](ApplicationVisibility.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**ApplicationLinks**](ApplicationLinks.md) |  | [optional] 

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *Application) GetAccessibility() ApplicationAccessibility`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *Application) GetAccessibilityOk() (*ApplicationAccessibility, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *Application) SetAccessibility(v ApplicationAccessibility)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *Application) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCreated

`func (o *Application) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Application) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Application) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Application) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFeatures

`func (o *Application) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Application) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Application) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Application) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Application) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Application) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Application) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Application) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Application) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Application) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Application) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Application) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLicensing

`func (o *Application) GetLicensing() ApplicationLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *Application) GetLicensingOk() (*ApplicationLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *Application) SetLicensing(v ApplicationLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *Application) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetProfile

`func (o *Application) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Application) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Application) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Application) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSignOnMode

`func (o *Application) GetSignOnMode() ApplicationSignOnMode`

GetSignOnMode returns the SignOnMode field if non-nil, zero value otherwise.

### GetSignOnModeOk

`func (o *Application) GetSignOnModeOk() (*ApplicationSignOnMode, bool)`

GetSignOnModeOk returns a tuple with the SignOnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnMode

`func (o *Application) SetSignOnMode(v ApplicationSignOnMode)`

SetSignOnMode sets SignOnMode field to given value.

### HasSignOnMode

`func (o *Application) HasSignOnMode() bool`

HasSignOnMode returns a boolean if a field has been set.

### GetStatus

`func (o *Application) GetStatus() ApplicationLifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Application) GetStatusOk() (*ApplicationLifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Application) SetStatus(v ApplicationLifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Application) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisibility

`func (o *Application) GetVisibility() ApplicationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Application) GetVisibilityOk() (*ApplicationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Application) SetVisibility(v ApplicationVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Application) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEmbedded

`func (o *Application) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Application) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Application) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *Application) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *Application) GetLinks() ApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Application) GetLinksOk() (*ApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Application) SetLinks(v ApplicationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Application) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


