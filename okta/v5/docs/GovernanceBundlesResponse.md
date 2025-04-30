# GovernanceBundlesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundles** | Pointer to [**[]GovernanceBundle**](GovernanceBundle.md) |  | [optional] 
**Links** | Pointer to [**GovernanceBundlesResponseLinks**](GovernanceBundlesResponseLinks.md) |  | [optional] 

## Methods

### NewGovernanceBundlesResponse

`func NewGovernanceBundlesResponse() *GovernanceBundlesResponse`

NewGovernanceBundlesResponse instantiates a new GovernanceBundlesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGovernanceBundlesResponseWithDefaults

`func NewGovernanceBundlesResponseWithDefaults() *GovernanceBundlesResponse`

NewGovernanceBundlesResponseWithDefaults instantiates a new GovernanceBundlesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundles

`func (o *GovernanceBundlesResponse) GetBundles() []GovernanceBundle`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *GovernanceBundlesResponse) GetBundlesOk() (*[]GovernanceBundle, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *GovernanceBundlesResponse) SetBundles(v []GovernanceBundle)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *GovernanceBundlesResponse) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetLinks

`func (o *GovernanceBundlesResponse) GetLinks() GovernanceBundlesResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GovernanceBundlesResponse) GetLinksOk() (*GovernanceBundlesResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GovernanceBundlesResponse) SetLinks(v GovernanceBundlesResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GovernanceBundlesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


