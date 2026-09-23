# CimdDocumentRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the CIMD requirement | [optional] 
**Id** | Pointer to **string** | Unique identifier for the CIMD requirement | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Name** | **string** | Human-readable name for the CIMD requirement | 
**Requirements** | [**[]CimdRequirementEntry**](CimdRequirementEntry.md) | List of requirement entries. Currently, you can only have one entry. | 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewCimdDocumentRequirements

`func NewCimdDocumentRequirements(name string, requirements []CimdRequirementEntry, ) *CimdDocumentRequirements`

NewCimdDocumentRequirements instantiates a new CimdDocumentRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdDocumentRequirementsWithDefaults

`func NewCimdDocumentRequirementsWithDefaults() *CimdDocumentRequirements`

NewCimdDocumentRequirementsWithDefaults instantiates a new CimdDocumentRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CimdDocumentRequirements) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CimdDocumentRequirements) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CimdDocumentRequirements) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CimdDocumentRequirements) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *CimdDocumentRequirements) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CimdDocumentRequirements) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CimdDocumentRequirements) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CimdDocumentRequirements) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CimdDocumentRequirements) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CimdDocumentRequirements) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CimdDocumentRequirements) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CimdDocumentRequirements) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CimdDocumentRequirements) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CimdDocumentRequirements) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CimdDocumentRequirements) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CimdDocumentRequirements) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *CimdDocumentRequirements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CimdDocumentRequirements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CimdDocumentRequirements) SetName(v string)`

SetName sets Name field to given value.


### GetRequirements

`func (o *CimdDocumentRequirements) GetRequirements() []CimdRequirementEntry`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *CimdDocumentRequirements) GetRequirementsOk() (*[]CimdRequirementEntry, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *CimdDocumentRequirements) SetRequirements(v []CimdRequirementEntry)`

SetRequirements sets Requirements field to given value.


### GetLinks

`func (o *CimdDocumentRequirements) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CimdDocumentRequirements) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CimdDocumentRequirements) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CimdDocumentRequirements) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


