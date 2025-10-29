# LinksResend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resend** | Pointer to [**HrefObject**](HrefObject.md) | Resends the factor enrollment challenge. See [Resend a factor enrollment](/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/resendEnrollFactor). | [optional] 

## Methods

### NewLinksResend

`func NewLinksResend() *LinksResend`

NewLinksResend instantiates a new LinksResend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksResendWithDefaults

`func NewLinksResendWithDefaults() *LinksResend`

NewLinksResendWithDefaults instantiates a new LinksResend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResend

`func (o *LinksResend) GetResend() HrefObject`

GetResend returns the Resend field if non-nil, zero value otherwise.

### GetResendOk

`func (o *LinksResend) GetResendOk() (*HrefObject, bool)`

GetResendOk returns a tuple with the Resend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResend

`func (o *LinksResend) SetResend(v HrefObject)`

SetResend sets Resend field to given value.

### HasResend

`func (o *LinksResend) HasResend() bool`

HasResend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


