# IdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IssuerMode** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**IdentityProviderPolicy**](IdentityProviderPolicy.md) |  | [optional] 
**Properties** | Pointer to [**NullableIdentityProviderProperties**](IdentityProviderProperties.md) |  | [optional] 
**Protocol** | Pointer to [**Protocol**](Protocol.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**IdentityProviderLinks**](IdentityProviderLinks.md) |  | [optional] 

## Methods

### NewIdentityProvider

`func NewIdentityProvider() *IdentityProvider`

NewIdentityProvider instantiates a new IdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderWithDefaults

`func NewIdentityProviderWithDefaults() *IdentityProvider`

NewIdentityProviderWithDefaults instantiates a new IdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IdentityProvider) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdentityProvider) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdentityProvider) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdentityProvider) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *IdentityProvider) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *IdentityProvider) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetId

`func (o *IdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuerMode

`func (o *IdentityProvider) GetIssuerMode() string`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *IdentityProvider) GetIssuerModeOk() (*string, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *IdentityProvider) SetIssuerMode(v string)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *IdentityProvider) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetLastUpdated

`func (o *IdentityProvider) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IdentityProvider) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IdentityProvider) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IdentityProvider) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *IdentityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *IdentityProvider) GetPolicy() IdentityProviderPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IdentityProvider) GetPolicyOk() (*IdentityProviderPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IdentityProvider) SetPolicy(v IdentityProviderPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IdentityProvider) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProperties

`func (o *IdentityProvider) GetProperties() IdentityProviderProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IdentityProvider) GetPropertiesOk() (*IdentityProviderProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IdentityProvider) SetProperties(v IdentityProviderProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IdentityProvider) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *IdentityProvider) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *IdentityProvider) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetProtocol

`func (o *IdentityProvider) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IdentityProvider) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IdentityProvider) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IdentityProvider) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityProvider) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityProvider) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityProvider) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityProvider) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *IdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *IdentityProvider) GetLinks() IdentityProviderLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityProvider) GetLinksOk() (*IdentityProviderLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityProvider) SetLinks(v IdentityProviderLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityProvider) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


