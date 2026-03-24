# DbscScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeSite** | **bool** | Whether to include subdomains in the binding scope (&#x60;false&#x60; &#x3D; exact origin only, &#x60;true&#x60; &#x3D; includes subdomains) | [readonly] 
**Origin** | **string** | The origin URL for which the DBSC binding is valid | [readonly] 

## Methods

### NewDbscScope

`func NewDbscScope(includeSite bool, origin string, ) *DbscScope`

NewDbscScope instantiates a new DbscScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbscScopeWithDefaults

`func NewDbscScopeWithDefaults() *DbscScope`

NewDbscScopeWithDefaults instantiates a new DbscScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeSite

`func (o *DbscScope) GetIncludeSite() bool`

GetIncludeSite returns the IncludeSite field if non-nil, zero value otherwise.

### GetIncludeSiteOk

`func (o *DbscScope) GetIncludeSiteOk() (*bool, bool)`

GetIncludeSiteOk returns a tuple with the IncludeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSite

`func (o *DbscScope) SetIncludeSite(v bool)`

SetIncludeSite sets IncludeSite field to given value.


### GetOrigin

`func (o *DbscScope) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DbscScope) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DbscScope) SetOrigin(v string)`

SetOrigin sets Origin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


