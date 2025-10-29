# SAMLPayLoadDataAssertion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to [**SAMLPayLoadDataAssertionSubject**](SAMLPayLoadDataAssertionSubject.md) |  | [optional] 
**Authentication** | Pointer to [**SAMLPayLoadDataAssertionAuthentication**](SAMLPayLoadDataAssertionAuthentication.md) |  | [optional] 
**Conditions** | Pointer to [**SAMLPayLoadDataAssertionConditions**](SAMLPayLoadDataAssertionConditions.md) |  | [optional] 
**Claims** | Pointer to [**map[string]SAMLPayLoadDataAssertionClaimsValue**](SAMLPayLoadDataAssertionClaimsValue.md) | Provides a JSON representation of the &#x60;&lt;saml:AttributeStatement&gt;&#x60; element contained in the generated SAML assertion. Contains any optional SAML attribute statements that you have defined for the app using the Admin Console&#39;s **SAML Settings**. | [optional] 
**Lifetime** | Pointer to [**SAMLPayLoadDataAssertionLifetime**](SAMLPayLoadDataAssertionLifetime.md) |  | [optional] 

## Methods

### NewSAMLPayLoadDataAssertion

`func NewSAMLPayLoadDataAssertion() *SAMLPayLoadDataAssertion`

NewSAMLPayLoadDataAssertion instantiates a new SAMLPayLoadDataAssertion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadDataAssertionWithDefaults

`func NewSAMLPayLoadDataAssertionWithDefaults() *SAMLPayLoadDataAssertion`

NewSAMLPayLoadDataAssertionWithDefaults instantiates a new SAMLPayLoadDataAssertion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *SAMLPayLoadDataAssertion) GetSubject() SAMLPayLoadDataAssertionSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SAMLPayLoadDataAssertion) GetSubjectOk() (*SAMLPayLoadDataAssertionSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SAMLPayLoadDataAssertion) SetSubject(v SAMLPayLoadDataAssertionSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SAMLPayLoadDataAssertion) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetAuthentication

`func (o *SAMLPayLoadDataAssertion) GetAuthentication() SAMLPayLoadDataAssertionAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *SAMLPayLoadDataAssertion) GetAuthenticationOk() (*SAMLPayLoadDataAssertionAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *SAMLPayLoadDataAssertion) SetAuthentication(v SAMLPayLoadDataAssertionAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *SAMLPayLoadDataAssertion) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetConditions

`func (o *SAMLPayLoadDataAssertion) GetConditions() SAMLPayLoadDataAssertionConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SAMLPayLoadDataAssertion) GetConditionsOk() (*SAMLPayLoadDataAssertionConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SAMLPayLoadDataAssertion) SetConditions(v SAMLPayLoadDataAssertionConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SAMLPayLoadDataAssertion) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetClaims

`func (o *SAMLPayLoadDataAssertion) GetClaims() map[string]SAMLPayLoadDataAssertionClaimsValue`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SAMLPayLoadDataAssertion) GetClaimsOk() (*map[string]SAMLPayLoadDataAssertionClaimsValue, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SAMLPayLoadDataAssertion) SetClaims(v map[string]SAMLPayLoadDataAssertionClaimsValue)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *SAMLPayLoadDataAssertion) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetLifetime

`func (o *SAMLPayLoadDataAssertion) GetLifetime() SAMLPayLoadDataAssertionLifetime`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *SAMLPayLoadDataAssertion) GetLifetimeOk() (*SAMLPayLoadDataAssertionLifetime, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *SAMLPayLoadDataAssertion) SetLifetime(v SAMLPayLoadDataAssertionLifetime)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *SAMLPayLoadDataAssertion) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


