# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateNetworkZone**](NetworkZoneApi.md#ActivateNetworkZone) | **Post** /api/v1/zones/{zoneId}/lifecycle/activate | Activate Network Zone
[**CreateNetworkZone**](NetworkZoneApi.md#CreateNetworkZone) | **Post** /api/v1/zones | Add Network Zone
[**DeactivateNetworkZone**](NetworkZoneApi.md#DeactivateNetworkZone) | **Post** /api/v1/zones/{zoneId}/lifecycle/deactivate | Deactivate Network Zone
[**DeleteNetworkZone**](NetworkZoneApi.md#DeleteNetworkZone) | **Delete** /api/v1/zones/{zoneId} | Delete Network Zone
[**GetNetworkZone**](NetworkZoneApi.md#GetNetworkZone) | **Get** /api/v1/zones/{zoneId} | Get Network Zone
[**ListNetworkZones**](NetworkZoneApi.md#ListNetworkZones) | **Get** /api/v1/zones | List Network Zones
[**UpdateNetworkZone**](NetworkZoneApi.md#UpdateNetworkZone) | **Put** /api/v1/zones/{zoneId} | Update Network Zone

# **ActivateNetworkZone**
> NetworkZone ActivateNetworkZone(ctx, zoneId)
Activate Network Zone

Activate Network Zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**|  | 

### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNetworkZone**
> NetworkZone CreateNetworkZone(ctx, body)
Add Network Zone

Adds a new network zone to your Okta organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkZone**](NetworkZone.md)|  | 

### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateNetworkZone**
> NetworkZone DeactivateNetworkZone(ctx, zoneId)
Deactivate Network Zone

Deactivates a network zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**|  | 

### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkZone**
> DeleteNetworkZone(ctx, zoneId)
Delete Network Zone

Removes network zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkZone**
> NetworkZone GetNetworkZone(ctx, zoneId)
Get Network Zone

Fetches a network zone from your Okta organization by `id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**|  | 

### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNetworkZones**
> []NetworkZone ListNetworkZones(ctx, optional)
List Network Zones

Enumerates network zones added to your organization with pagination. A subset of zones can be returned that match a supported filter expression or query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworkZoneApiListNetworkZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworkZoneApiListNetworkZonesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **optional.String**| Specifies the pagination cursor for the next page of network zones | 
 **limit** | **optional.Int32**| Specifies the number of results for a page | [default to -1]
 **filter** | **optional.String**| Filters zones by usage or id expression | 

### Return type

[**[]NetworkZone**](NetworkZone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkZone**
> NetworkZone UpdateNetworkZone(ctx, body, zoneId)
Update Network Zone

Updates a network zone in your organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkZone**](NetworkZone.md)|  | 
  **zoneId** | **string**|  | 

### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

