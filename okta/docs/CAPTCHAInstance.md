# CAPTCHAInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SiteKey** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**CAPTCHAType**](CAPTCHAType.md) |  | [optional] 
**Links** | Pointer to [**ApiTokenLink**](ApiTokenLink.md) |  | [optional] 

## Methods

### NewCAPTCHAInstance

`func NewCAPTCHAInstance() *CAPTCHAInstance`

NewCAPTCHAInstance instantiates a new CAPTCHAInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCAPTCHAInstanceWithDefaults

`func NewCAPTCHAInstanceWithDefaults() *CAPTCHAInstance`

NewCAPTCHAInstanceWithDefaults instantiates a new CAPTCHAInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CAPTCHAInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CAPTCHAInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CAPTCHAInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CAPTCHAInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CAPTCHAInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CAPTCHAInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CAPTCHAInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CAPTCHAInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecretKey

`func (o *CAPTCHAInstance) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CAPTCHAInstance) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CAPTCHAInstance) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CAPTCHAInstance) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSiteKey

`func (o *CAPTCHAInstance) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *CAPTCHAInstance) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *CAPTCHAInstance) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.

### HasSiteKey

`func (o *CAPTCHAInstance) HasSiteKey() bool`

HasSiteKey returns a boolean if a field has been set.

### GetType

`func (o *CAPTCHAInstance) GetType() CAPTCHAType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CAPTCHAInstance) GetTypeOk() (*CAPTCHAType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CAPTCHAInstance) SetType(v CAPTCHAType)`

SetType sets Type field to given value.

### HasType

`func (o *CAPTCHAInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *CAPTCHAInstance) GetLinks() ApiTokenLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CAPTCHAInstance) GetLinksOk() (*ApiTokenLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CAPTCHAInstance) SetLinks(v ApiTokenLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CAPTCHAInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


