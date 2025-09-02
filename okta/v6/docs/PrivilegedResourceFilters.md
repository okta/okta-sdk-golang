# PrivilegedResourceFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppGroups** | Pointer to [**[]AppGroup**](AppGroup.md) | Array of app groups whose members might be privileged app users | [optional] 
**OrganizationalUnits** | Pointer to [**[]OrganizationalUnit**](OrganizationalUnit.md) | Array of organizational units where privileged app users are present | [optional] 

## Methods

### NewPrivilegedResourceFilters

`func NewPrivilegedResourceFilters() *PrivilegedResourceFilters`

NewPrivilegedResourceFilters instantiates a new PrivilegedResourceFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedResourceFiltersWithDefaults

`func NewPrivilegedResourceFiltersWithDefaults() *PrivilegedResourceFilters`

NewPrivilegedResourceFiltersWithDefaults instantiates a new PrivilegedResourceFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppGroups

`func (o *PrivilegedResourceFilters) GetAppGroups() []AppGroup`

GetAppGroups returns the AppGroups field if non-nil, zero value otherwise.

### GetAppGroupsOk

`func (o *PrivilegedResourceFilters) GetAppGroupsOk() (*[]AppGroup, bool)`

GetAppGroupsOk returns a tuple with the AppGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGroups

`func (o *PrivilegedResourceFilters) SetAppGroups(v []AppGroup)`

SetAppGroups sets AppGroups field to given value.

### HasAppGroups

`func (o *PrivilegedResourceFilters) HasAppGroups() bool`

HasAppGroups returns a boolean if a field has been set.

### GetOrganizationalUnits

`func (o *PrivilegedResourceFilters) GetOrganizationalUnits() []OrganizationalUnit`

GetOrganizationalUnits returns the OrganizationalUnits field if non-nil, zero value otherwise.

### GetOrganizationalUnitsOk

`func (o *PrivilegedResourceFilters) GetOrganizationalUnitsOk() (*[]OrganizationalUnit, bool)`

GetOrganizationalUnitsOk returns a tuple with the OrganizationalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnits

`func (o *PrivilegedResourceFilters) SetOrganizationalUnits(v []OrganizationalUnit)`

SetOrganizationalUnits sets OrganizationalUnits field to given value.

### HasOrganizationalUnits

`func (o *PrivilegedResourceFilters) HasOrganizationalUnits() bool`

HasOrganizationalUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


