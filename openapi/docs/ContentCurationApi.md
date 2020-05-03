# \ContentCurationApi

All URIs are relative to *https://demo.funkwhale.audio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1FavoritesTracksGet**](ContentCurationApi.md#ApiV1FavoritesTracksGet) | **Get** /api/v1/favorites/tracks/ | 
[**ApiV1FavoritesTracksPost**](ContentCurationApi.md#ApiV1FavoritesTracksPost) | **Post** /api/v1/favorites/tracks/ | Mark the given track as favorite
[**ApiV1FavoritesTracksRemovePost**](ContentCurationApi.md#ApiV1FavoritesTracksRemovePost) | **Post** /api/v1/favorites/tracks/remove/ | Remove the given track from favorites



## ApiV1FavoritesTracksGet

> ResultPage ApiV1FavoritesTracksGet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiV1FavoritesTracksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1FavoritesTracksGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Search query used to filter favorites | 
 **user** | [**optional.Interface of map[string]interface{}**](.md)| Limit results to favorites belonging to the given user | 

### Return type

[**ResultPage**](ResultPage.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FavoritesTracksPost

> InlineResponse201 ApiV1FavoritesTracksPost(ctx, inlineObject5)

Mark the given track as favorite

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject5** | [**InlineObject5**](InlineObject5.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FavoritesTracksRemovePost

> ApiV1FavoritesTracksRemovePost(ctx, inlineObject6)

Remove the given track from favorites

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject6** | [**InlineObject6**](InlineObject6.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

