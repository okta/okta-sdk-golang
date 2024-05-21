# EmailDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | **string** |  | 
**Domain** | **string** |  | 
**ValidationSubdomain** | Pointer to **string** | Subdomain for the email sender&#39;s custom mail domain. Specify your subdomain when you configure a custom mail domain. | [optional] [default to "mail"]
**DisplayName** | **string** |  | 
**UserName** | **string** |  | 

## Methods

### NewEmailDomain

`func NewEmailDomain(brandId string, domain string, displayName string, userName string, ) *EmailDomain`

NewEmailDomain instantiates a new EmailDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainWithDefaults

`func NewEmailDomainWithDefaults() *EmailDomain`

NewEmailDomainWithDefaults instantiates a new EmailDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *EmailDomain) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *EmailDomain) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *EmailDomain) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.


### GetDomain

`func (o *EmailDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetValidationSubdomain

`func (o *EmailDomain) GetValidationSubdomain() string`

GetValidationSubdomain returns the ValidationSubdomain field if non-nil, zero value otherwise.

### GetValidationSubdomainOk

`func (o *EmailDomain) GetValidationSubdomainOk() (*string, bool)`

GetValidationSubdomainOk returns a tuple with the ValidationSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationSubdomain

`func (o *EmailDomain) SetValidationSubdomain(v string)`

SetValidationSubdomain sets ValidationSubdomain field to given value.

### HasValidationSubdomain

`func (o *EmailDomain) HasValidationSubdomain() bool`

HasValidationSubdomain returns a boolean if a field has been set.

### GetDisplayName

`func (o *EmailDomain) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EmailDomain) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EmailDomain) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetUserName

`func (o *EmailDomain) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EmailDomain) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EmailDomain) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


