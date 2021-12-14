# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificate**](DomainApi.md#CreateCertificate) | **Put** /api/v1/domains/{domainId}/certificate | Create Certificate
[**CreateDomain**](DomainApi.md#CreateDomain) | **Post** /api/v1/domains | Create Domain
[**DeleteDomain**](DomainApi.md#DeleteDomain) | **Delete** /api/v1/domains/{domainId} | Delete Domain
[**GetDomain**](DomainApi.md#GetDomain) | **Get** /api/v1/domains/{domainId} | Get Domain
[**ListDomains**](DomainApi.md#ListDomains) | **Get** /api/v1/domains | List Domains
[**VerifyDomain**](DomainApi.md#VerifyDomain) | **Post** /api/v1/domains/{domainId}/verify | Verify Domain

# **CreateCertificate**
> CreateCertificate(ctx, body, domainId)
Create Certificate

Creates the Certificate for the Domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainCertificate**](DomainCertificate.md)|  | 
  **domainId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDomain**
> DomainResponse CreateDomain(ctx, body)
Create Domain

Creates your domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Domain**](Domain.md)|  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomain**
> DeleteDomain(ctx, domainId)
Delete Domain

Deletes a Domain by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDomain**
> DomainResponse GetDomain(ctx, domainId)
Get Domain

Fetches a Domain by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomains**
> DomainListResponse ListDomains(ctx, )
List Domains

List all verified custom Domains for the org.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DomainListResponse**](DomainListResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyDomain**
> DomainResponse VerifyDomain(ctx, domainId)
Verify Domain

Verifies the Domain by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**|  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

