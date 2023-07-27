# CatalogApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SignOnModes** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**CatalogApplicationStatus**](CatalogApplicationStatus.md) |  | [optional] 
**VerificationStatus** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewCatalogApplication

`func NewCatalogApplication() *CatalogApplication`

NewCatalogApplication instantiates a new CatalogApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogApplicationWithDefaults

`func NewCatalogApplicationWithDefaults() *CatalogApplication`

NewCatalogApplicationWithDefaults instantiates a new CatalogApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CatalogApplication) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CatalogApplication) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CatalogApplication) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CatalogApplication) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogApplication) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogApplication) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogApplication) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogApplication) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFeatures

`func (o *CatalogApplication) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CatalogApplication) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CatalogApplication) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CatalogApplication) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *CatalogApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CatalogApplication) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CatalogApplication) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CatalogApplication) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CatalogApplication) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *CatalogApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSignOnModes

`func (o *CatalogApplication) GetSignOnModes() []string`

GetSignOnModes returns the SignOnModes field if non-nil, zero value otherwise.

### GetSignOnModesOk

`func (o *CatalogApplication) GetSignOnModesOk() (*[]string, bool)`

GetSignOnModesOk returns a tuple with the SignOnModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnModes

`func (o *CatalogApplication) SetSignOnModes(v []string)`

SetSignOnModes sets SignOnModes field to given value.

### HasSignOnModes

`func (o *CatalogApplication) HasSignOnModes() bool`

HasSignOnModes returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogApplication) GetStatus() CatalogApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogApplication) GetStatusOk() (*CatalogApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogApplication) SetStatus(v CatalogApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *CatalogApplication) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *CatalogApplication) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *CatalogApplication) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *CatalogApplication) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetWebsite

`func (o *CatalogApplication) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CatalogApplication) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CatalogApplication) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CatalogApplication) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetLinks

`func (o *CatalogApplication) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CatalogApplication) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CatalogApplication) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CatalogApplication) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


