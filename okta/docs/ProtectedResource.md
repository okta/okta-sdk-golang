# ProtectedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the protected resource | [readonly] 
**ResourceType** | **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;The type of a resource protected by an authorization server | 
**ResourceUrl** | Pointer to **string** | The URL of the protected resource. Present only for resource types that expose a resource URL. | [optional] 
**Status** | **string** | Current status of the resource in its lifecycle | [readonly] 
**Links** | [**LinksSelf**](LinksSelf.md) |  | 

## Methods

### NewProtectedResource

`func NewProtectedResource(id string, resourceType string, status string, links LinksSelf, ) *ProtectedResource`

NewProtectedResource instantiates a new ProtectedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedResourceWithDefaults

`func NewProtectedResourceWithDefaults() *ProtectedResource`

NewProtectedResourceWithDefaults instantiates a new ProtectedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProtectedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectedResource) SetId(v string)`

SetId sets Id field to given value.


### GetResourceType

`func (o *ProtectedResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ProtectedResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ProtectedResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceUrl

`func (o *ProtectedResource) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *ProtectedResource) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *ProtectedResource) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *ProtectedResource) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.

### GetStatus

`func (o *ProtectedResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProtectedResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProtectedResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLinks

`func (o *ProtectedResource) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProtectedResource) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProtectedResource) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


