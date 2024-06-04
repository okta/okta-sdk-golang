# OAuth2Claim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysIncludeInToken** | Pointer to **bool** | Specifies whether to include Claims in the token. The value is always &#x60;TRUE&#x60; for access token Claims. If the value is set to &#x60;FALSE&#x60; for an ID token claim, the Claim isn&#39;t included in the ID token when the token is requested with the access token or with the &#x60;authorization_code&#x60;. The client instead uses the access token to get Claims from the &#x60;/userinfo&#x60; endpoint. | [optional] 
**ClaimType** | Pointer to **string** | Specifies whether the Claim is for an access token (&#x60;RESOURCE&#x60;) or an ID token (&#x60;IDENTITY&#x60;) | [optional] 
**Conditions** | Pointer to [**OAuth2ClaimConditions**](OAuth2ClaimConditions.md) |  | [optional] 
**GroupFilterType** | Pointer to **string** | Specifies the type of group filter if &#x60;valueType&#x60; is &#x60;GROUPS&#x60;  If &#x60;valueType&#x60; is &#x60;GROUPS&#x60;, then the groups returned are filtered according to the value of &#x60;group_filter_type&#x60;.  If you have complex filters for Groups, you can [create a Groups allowlist](https://developer.okta.com/docs/guides/customize-tokens-groups-claim/main/) to put them all in a Claim. | [optional] 
**Id** | Pointer to **string** | ID of the Claim | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Claim | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**System** | Pointer to **bool** | When &#x60;true&#x60;, indicates that Okta created the Claim | [optional] 
**Value** | Pointer to **string** | Specifies the value of the Claim. This value must be a string literal if &#x60;valueType&#x60; is &#x60;GROUPS&#x60;, and the string literal is matched with the selected &#x60;group_filter_type&#x60;. The value must be an Okta EL expression if &#x60;valueType&#x60; is &#x60;EXPRESSION&#x60;. | [optional] 
**ValueType** | Pointer to **string** | Specifies whether the Claim is an Okta Expression Language (EL) expression (&#x60;EXPRESSION&#x60;), a set of groups (&#x60;GROUPS&#x60;), or a system claim (&#x60;SYSTEM&#x60;) | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewOAuth2Claim

`func NewOAuth2Claim() *OAuth2Claim`

NewOAuth2Claim instantiates a new OAuth2Claim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClaimWithDefaults

`func NewOAuth2ClaimWithDefaults() *OAuth2Claim`

NewOAuth2ClaimWithDefaults instantiates a new OAuth2Claim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysIncludeInToken

`func (o *OAuth2Claim) GetAlwaysIncludeInToken() bool`

GetAlwaysIncludeInToken returns the AlwaysIncludeInToken field if non-nil, zero value otherwise.

### GetAlwaysIncludeInTokenOk

`func (o *OAuth2Claim) GetAlwaysIncludeInTokenOk() (*bool, bool)`

GetAlwaysIncludeInTokenOk returns a tuple with the AlwaysIncludeInToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysIncludeInToken

`func (o *OAuth2Claim) SetAlwaysIncludeInToken(v bool)`

SetAlwaysIncludeInToken sets AlwaysIncludeInToken field to given value.

### HasAlwaysIncludeInToken

`func (o *OAuth2Claim) HasAlwaysIncludeInToken() bool`

HasAlwaysIncludeInToken returns a boolean if a field has been set.

### GetClaimType

`func (o *OAuth2Claim) GetClaimType() string`

GetClaimType returns the ClaimType field if non-nil, zero value otherwise.

### GetClaimTypeOk

`func (o *OAuth2Claim) GetClaimTypeOk() (*string, bool)`

GetClaimTypeOk returns a tuple with the ClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimType

`func (o *OAuth2Claim) SetClaimType(v string)`

SetClaimType sets ClaimType field to given value.

### HasClaimType

`func (o *OAuth2Claim) HasClaimType() bool`

HasClaimType returns a boolean if a field has been set.

### GetConditions

`func (o *OAuth2Claim) GetConditions() OAuth2ClaimConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *OAuth2Claim) GetConditionsOk() (*OAuth2ClaimConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *OAuth2Claim) SetConditions(v OAuth2ClaimConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *OAuth2Claim) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetGroupFilterType

`func (o *OAuth2Claim) GetGroupFilterType() string`

GetGroupFilterType returns the GroupFilterType field if non-nil, zero value otherwise.

### GetGroupFilterTypeOk

`func (o *OAuth2Claim) GetGroupFilterTypeOk() (*string, bool)`

GetGroupFilterTypeOk returns a tuple with the GroupFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilterType

`func (o *OAuth2Claim) SetGroupFilterType(v string)`

SetGroupFilterType sets GroupFilterType field to given value.

### HasGroupFilterType

`func (o *OAuth2Claim) HasGroupFilterType() bool`

HasGroupFilterType returns a boolean if a field has been set.

### GetId

`func (o *OAuth2Claim) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2Claim) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2Claim) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2Claim) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OAuth2Claim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2Claim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2Claim) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OAuth2Claim) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2Claim) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2Claim) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2Claim) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2Claim) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *OAuth2Claim) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *OAuth2Claim) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *OAuth2Claim) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *OAuth2Claim) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetValue

`func (o *OAuth2Claim) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OAuth2Claim) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OAuth2Claim) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OAuth2Claim) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *OAuth2Claim) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *OAuth2Claim) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *OAuth2Claim) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *OAuth2Claim) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2Claim) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2Claim) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2Claim) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2Claim) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


