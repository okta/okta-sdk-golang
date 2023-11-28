# CsrMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**CsrMetadataSubject**](CsrMetadataSubject.md) |  | [optional] 
**SubjectAltNames** | Pointer to [**CsrMetadataSubjectAltNames**](CsrMetadataSubjectAltNames.md) |  | [optional] 

## Methods

### NewCsrMetadata

`func NewCsrMetadata() *CsrMetadata`

NewCsrMetadata instantiates a new CsrMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsrMetadataWithDefaults

`func NewCsrMetadataWithDefaults() *CsrMetadata`

NewCsrMetadataWithDefaults instantiates a new CsrMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CsrMetadata) GetSubject() CsrMetadataSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CsrMetadata) GetSubjectOk() (*CsrMetadataSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CsrMetadata) SetSubject(v CsrMetadataSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CsrMetadata) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectAltNames

`func (o *CsrMetadata) GetSubjectAltNames() CsrMetadataSubjectAltNames`

GetSubjectAltNames returns the SubjectAltNames field if non-nil, zero value otherwise.

### GetSubjectAltNamesOk

`func (o *CsrMetadata) GetSubjectAltNamesOk() (*CsrMetadataSubjectAltNames, bool)`

GetSubjectAltNamesOk returns a tuple with the SubjectAltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNames

`func (o *CsrMetadata) SetSubjectAltNames(v CsrMetadataSubjectAltNames)`

SetSubjectAltNames sets SubjectAltNames field to given value.

### HasSubjectAltNames

`func (o *CsrMetadata) HasSubjectAltNames() bool`

HasSubjectAltNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


