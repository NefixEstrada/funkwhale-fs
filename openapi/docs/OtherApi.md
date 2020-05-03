# \OtherApi

All URIs are relative to *https://demo.funkwhale.audio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AttachmentsPost**](OtherApi.md#ApiV1AttachmentsPost) | **Post** /api/v1/attachments/ | 
[**ApiV1AttachmentsUuidDelete**](OtherApi.md#ApiV1AttachmentsUuidDelete) | **Delete** /api/v1/attachments/{uuid}/ | Delete an attachment
[**ApiV1AttachmentsUuidGet**](OtherApi.md#ApiV1AttachmentsUuidGet) | **Get** /api/v1/attachments/{uuid}/ | Retrieve an attachment



## ApiV1AttachmentsPost

> ApiV1AttachmentsPost(ctx, optional)



Upload a new file as an attachment that can be later associated with other objects.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiV1AttachmentsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1AttachmentsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AttachmentsUuidDelete

> ApiV1AttachmentsUuidDelete(ctx, uuid)

Delete an attachment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AttachmentsUuidGet

> Attachment ApiV1AttachmentsUuidGet(ctx, uuid)

Retrieve an attachment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 

### Return type

[**Attachment**](Attachment.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

