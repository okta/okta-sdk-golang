# CimdClientEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to [**[]CimdClientEntityBindingRequest**](CimdClientEntityBindingRequest.md) | Optional list of bindings to create alongside the CIMD client entity | [optional] 
**ClientIdMatchPattern** | **string** | The client ID match pattern. * For the &#x60;EXACT&#x60; strategy, this is a single CIMD client ID URL that&#39;s stored in normalized form (lowercase and with the trailing slash removed). * For the &#x60;REGEX&#x60; strategy, this is a regular expression (regex) evaluated against the client&#39;s &#x60;client_id&#x60; URL. | 
**ClientIdMatchStrategy** | **string** | The strategy that Okta uses to match a &#x60;client_id&#x60; URL to a CIMD client entity | 
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the CIMD client entity | [optional] 
**Id** | Pointer to **string** | Unique identifier of the CIMD client entity | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**MetadataRequirementsId** | Pointer to **NullableString** | The ID of a [CIMD requirement](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/CimdDocumentRequirements/) that Okta evaluates against the metadata document of any CIMD client this entity matches. If any requirement fails, Okta rejects the client&#39;s authorization request. | [optional] 
**Name** | **string** | Human-readable name of the CIMD client entity | 
**Priority** | Pointer to **NullableInt32** | Determines the evaluation order when multiple &#x60;REGEX&#x60;-strategy client entities can match the same &#x60;client_id&#x60; URL. You must have a value for &#x60;priority&#x60; if you use a &#x60;REGEX&#x60; matching strategy and the value must be unique across all &#x60;REGEX&#x60;-strategy client entities in the org. For example, only one &#x60;REGEX&#x60;-strategy client entity can have a priority of &#x60;0&#x60;, only one can have a priority of &#x60;1&#x60;, and so on.  If you use the &#x60;EXACT&#x60; strategy, &#x60;priority&#x60; must be null.  When a &#x60;client_id&#x60; URL matches more than one &#x60;REGEX&#x60;-strategy client entity, Okta evaluates client entities from lowest to highest priority value and stops at the first match. Okta applies the matching client entity&#39;s configuration, including any associated document requirements, to the authorization request. | [optional] 
**Embedded** | Pointer to [**CimdClientEntityEmbedded**](CimdClientEntityEmbedded.md) |  | [optional] 
**Links** | Pointer to [**CimdClientEntityLinks**](CimdClientEntityLinks.md) |  | [optional] 

## Methods

### NewCimdClientEntity

`func NewCimdClientEntity(clientIdMatchPattern string, clientIdMatchStrategy string, name string, ) *CimdClientEntity`

NewCimdClientEntity instantiates a new CimdClientEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdClientEntityWithDefaults

`func NewCimdClientEntityWithDefaults() *CimdClientEntity`

NewCimdClientEntityWithDefaults instantiates a new CimdClientEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *CimdClientEntity) GetBindings() []CimdClientEntityBindingRequest`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *CimdClientEntity) GetBindingsOk() (*[]CimdClientEntityBindingRequest, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *CimdClientEntity) SetBindings(v []CimdClientEntityBindingRequest)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *CimdClientEntity) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetClientIdMatchPattern

`func (o *CimdClientEntity) GetClientIdMatchPattern() string`

GetClientIdMatchPattern returns the ClientIdMatchPattern field if non-nil, zero value otherwise.

### GetClientIdMatchPatternOk

`func (o *CimdClientEntity) GetClientIdMatchPatternOk() (*string, bool)`

GetClientIdMatchPatternOk returns a tuple with the ClientIdMatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdMatchPattern

`func (o *CimdClientEntity) SetClientIdMatchPattern(v string)`

SetClientIdMatchPattern sets ClientIdMatchPattern field to given value.


### GetClientIdMatchStrategy

`func (o *CimdClientEntity) GetClientIdMatchStrategy() string`

GetClientIdMatchStrategy returns the ClientIdMatchStrategy field if non-nil, zero value otherwise.

### GetClientIdMatchStrategyOk

`func (o *CimdClientEntity) GetClientIdMatchStrategyOk() (*string, bool)`

GetClientIdMatchStrategyOk returns a tuple with the ClientIdMatchStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdMatchStrategy

`func (o *CimdClientEntity) SetClientIdMatchStrategy(v string)`

SetClientIdMatchStrategy sets ClientIdMatchStrategy field to given value.


### GetCreated

`func (o *CimdClientEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CimdClientEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CimdClientEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CimdClientEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *CimdClientEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CimdClientEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CimdClientEntity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CimdClientEntity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CimdClientEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CimdClientEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CimdClientEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CimdClientEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CimdClientEntity) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CimdClientEntity) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CimdClientEntity) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CimdClientEntity) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMetadataRequirementsId

`func (o *CimdClientEntity) GetMetadataRequirementsId() string`

GetMetadataRequirementsId returns the MetadataRequirementsId field if non-nil, zero value otherwise.

### GetMetadataRequirementsIdOk

`func (o *CimdClientEntity) GetMetadataRequirementsIdOk() (*string, bool)`

GetMetadataRequirementsIdOk returns a tuple with the MetadataRequirementsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRequirementsId

`func (o *CimdClientEntity) SetMetadataRequirementsId(v string)`

SetMetadataRequirementsId sets MetadataRequirementsId field to given value.

### HasMetadataRequirementsId

`func (o *CimdClientEntity) HasMetadataRequirementsId() bool`

HasMetadataRequirementsId returns a boolean if a field has been set.

### SetMetadataRequirementsIdNil

`func (o *CimdClientEntity) SetMetadataRequirementsIdNil(b bool)`

 SetMetadataRequirementsIdNil sets the value for MetadataRequirementsId to be an explicit nil

### UnsetMetadataRequirementsId
`func (o *CimdClientEntity) UnsetMetadataRequirementsId()`

UnsetMetadataRequirementsId ensures that no value is present for MetadataRequirementsId, not even an explicit nil
### GetName

`func (o *CimdClientEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CimdClientEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CimdClientEntity) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *CimdClientEntity) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CimdClientEntity) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CimdClientEntity) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CimdClientEntity) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CimdClientEntity) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CimdClientEntity) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetEmbedded

`func (o *CimdClientEntity) GetEmbedded() CimdClientEntityEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CimdClientEntity) GetEmbeddedOk() (*CimdClientEntityEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CimdClientEntity) SetEmbedded(v CimdClientEntityEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CimdClientEntity) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *CimdClientEntity) GetLinks() CimdClientEntityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CimdClientEntity) GetLinksOk() (*CimdClientEntityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CimdClientEntity) SetLinks(v CimdClientEntityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CimdClientEntity) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


