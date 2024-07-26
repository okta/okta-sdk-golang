# GetFactorTransactionStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorResult** | Pointer to **string** | Result of the verification transaction | [optional] 
**Profile** | Pointer to [**UserFactorPushTransactionRejectedAllOfProfile**](UserFactorPushTransactionRejectedAllOfProfile.md) |  | [optional] 
**Links** | Pointer to [**UserFactorPushTransactionRejectedAllOfLinks**](UserFactorPushTransactionRejectedAllOfLinks.md) |  | [optional] 

## Methods

### NewGetFactorTransactionStatus200Response

`func NewGetFactorTransactionStatus200Response() *GetFactorTransactionStatus200Response`

NewGetFactorTransactionStatus200Response instantiates a new GetFactorTransactionStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFactorTransactionStatus200ResponseWithDefaults

`func NewGetFactorTransactionStatus200ResponseWithDefaults() *GetFactorTransactionStatus200Response`

NewGetFactorTransactionStatus200ResponseWithDefaults instantiates a new GetFactorTransactionStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorResult

`func (o *GetFactorTransactionStatus200Response) GetFactorResult() string`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *GetFactorTransactionStatus200Response) GetFactorResultOk() (*string, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *GetFactorTransactionStatus200Response) SetFactorResult(v string)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *GetFactorTransactionStatus200Response) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetProfile

`func (o *GetFactorTransactionStatus200Response) GetProfile() UserFactorPushTransactionRejectedAllOfProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetFactorTransactionStatus200Response) GetProfileOk() (*UserFactorPushTransactionRejectedAllOfProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetFactorTransactionStatus200Response) SetProfile(v UserFactorPushTransactionRejectedAllOfProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetFactorTransactionStatus200Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetLinks

`func (o *GetFactorTransactionStatus200Response) GetLinks() UserFactorPushTransactionRejectedAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFactorTransactionStatus200Response) GetLinksOk() (*UserFactorPushTransactionRejectedAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFactorTransactionStatus200Response) SetLinks(v UserFactorPushTransactionRejectedAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetFactorTransactionStatus200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


