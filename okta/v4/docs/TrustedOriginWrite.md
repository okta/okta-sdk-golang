# TrustedOriginWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name for the Trusted Origin | [optional] 
**Origin** | Pointer to **string** | Unique origin URL for the Trusted Origin. The supported schemes for this attribute are HTTP, HTTPS, FTP, Ionic 2, and Capacitor. | [optional] 
**Scopes** | Pointer to [**[]TrustedOriginScope**](TrustedOriginScope.md) | Array of Scope types that this Trusted Origin is used for | [optional] 

## Methods

### NewTrustedOriginWrite

`func NewTrustedOriginWrite() *TrustedOriginWrite`

NewTrustedOriginWrite instantiates a new TrustedOriginWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedOriginWriteWithDefaults

`func NewTrustedOriginWriteWithDefaults() *TrustedOriginWrite`

NewTrustedOriginWriteWithDefaults instantiates a new TrustedOriginWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrustedOriginWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustedOriginWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustedOriginWrite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustedOriginWrite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *TrustedOriginWrite) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *TrustedOriginWrite) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *TrustedOriginWrite) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *TrustedOriginWrite) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetScopes

`func (o *TrustedOriginWrite) GetScopes() []TrustedOriginScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TrustedOriginWrite) GetScopesOk() (*[]TrustedOriginScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TrustedOriginWrite) SetScopes(v []TrustedOriginScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TrustedOriginWrite) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


