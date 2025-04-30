# ApplicationGroupAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the [Group](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/) | [optional] [readonly] 
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Priority** | Pointer to **int32** | Priority assigned to the group. If an app has more than one group assigned to the same user, then the group with the higher priority has its profile applied to the [Application User](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/ApplicationUsers/). If a priority value isn&#39;t specified, then the next highest priority is assigned by default. See [Assign attribute group priority](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;ext-usgp-app-group-priority) and the [sample priority use case](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;ext-usgp-combine-values-use). | [optional] 
**Profile** | Pointer to **map[string]interface{}** | Specifies the profile properties applied to [Application Users](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/ApplicationUsers/) that are assigned to the app through group membership. Some reference properties are imported from the target app and can&#39;t be configured. See [profile](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c&#x3D;200&amp;path&#x3D;profile&amp;t&#x3D;response). | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | Embedded resource related to the Application Group using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. If the &#x60;expand&#x3D;group&#x60; query parameter is specified, then the [Group](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/) object is embedded. If the &#x60;expand&#x3D;metadata&#x60; query parameter is specified, then the group assignment metadata is embedded. | [optional] [readonly] 
**Links** | Pointer to [**ApplicationGroupAssignmentLinks**](ApplicationGroupAssignmentLinks.md) |  | [optional] 

## Methods

### NewApplicationGroupAssignment

`func NewApplicationGroupAssignment() *ApplicationGroupAssignment`

NewApplicationGroupAssignment instantiates a new ApplicationGroupAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGroupAssignmentWithDefaults

`func NewApplicationGroupAssignmentWithDefaults() *ApplicationGroupAssignment`

NewApplicationGroupAssignmentWithDefaults instantiates a new ApplicationGroupAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationGroupAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationGroupAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationGroupAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationGroupAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ApplicationGroupAssignment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApplicationGroupAssignment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApplicationGroupAssignment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ApplicationGroupAssignment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPriority

`func (o *ApplicationGroupAssignment) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplicationGroupAssignment) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplicationGroupAssignment) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApplicationGroupAssignment) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProfile

`func (o *ApplicationGroupAssignment) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ApplicationGroupAssignment) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ApplicationGroupAssignment) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ApplicationGroupAssignment) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetEmbedded

`func (o *ApplicationGroupAssignment) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApplicationGroupAssignment) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApplicationGroupAssignment) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ApplicationGroupAssignment) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationGroupAssignment) GetLinks() ApplicationGroupAssignmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationGroupAssignment) GetLinksOk() (*ApplicationGroupAssignmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationGroupAssignment) SetLinks(v ApplicationGroupAssignmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationGroupAssignment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


