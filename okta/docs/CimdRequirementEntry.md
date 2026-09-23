# CimdRequirementEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementType** | **string** | The type of the CIMD requirement | 
**RequirementValue** | **string** | An [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language/) expression that Okta evaluates against the client&#39;s fetched CIMD. The expression must evaluate to a boolean. When the expression returns &#x60;false&#x60;, Okta rejects the client&#39;s authorization request. | 

## Methods

### NewCimdRequirementEntry

`func NewCimdRequirementEntry(requirementType string, requirementValue string, ) *CimdRequirementEntry`

NewCimdRequirementEntry instantiates a new CimdRequirementEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdRequirementEntryWithDefaults

`func NewCimdRequirementEntryWithDefaults() *CimdRequirementEntry`

NewCimdRequirementEntryWithDefaults instantiates a new CimdRequirementEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementType

`func (o *CimdRequirementEntry) GetRequirementType() string`

GetRequirementType returns the RequirementType field if non-nil, zero value otherwise.

### GetRequirementTypeOk

`func (o *CimdRequirementEntry) GetRequirementTypeOk() (*string, bool)`

GetRequirementTypeOk returns a tuple with the RequirementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementType

`func (o *CimdRequirementEntry) SetRequirementType(v string)`

SetRequirementType sets RequirementType field to given value.


### GetRequirementValue

`func (o *CimdRequirementEntry) GetRequirementValue() string`

GetRequirementValue returns the RequirementValue field if non-nil, zero value otherwise.

### GetRequirementValueOk

`func (o *CimdRequirementEntry) GetRequirementValueOk() (*string, bool)`

GetRequirementValueOk returns a tuple with the RequirementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementValue

`func (o *CimdRequirementEntry) SetRequirementValue(v string)`

SetRequirementValue sets RequirementValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


