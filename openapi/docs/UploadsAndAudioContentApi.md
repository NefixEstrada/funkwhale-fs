# \UploadsAndAudioContentApi

All URIs are relative to *https://demo.funkwhale.audio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ChannelsGet**](UploadsAndAudioContentApi.md#ApiV1ChannelsGet) | **Get** /api/v1/channels/ | List channels
[**ApiV1ChannelsPost**](UploadsAndAudioContentApi.md#ApiV1ChannelsPost) | **Post** /api/v1/channels/ | Create a new channel
[**ApiV1ChannelsUuidDelete**](UploadsAndAudioContentApi.md#ApiV1ChannelsUuidDelete) | **Delete** /api/v1/channels/{uuid}/ | Delete a channel and all associated uploads
[**ApiV1ChannelsUuidGet**](UploadsAndAudioContentApi.md#ApiV1ChannelsUuidGet) | **Get** /api/v1/channels/{uuid}/ | Retrieve a channel
[**ApiV1ChannelsUuidPost**](UploadsAndAudioContentApi.md#ApiV1ChannelsUuidPost) | **Post** /api/v1/channels/{uuid}/ | Update a channel
[**ApiV1LibrariesGet**](UploadsAndAudioContentApi.md#ApiV1LibrariesGet) | **Get** /api/v1/libraries/ | List owned libraries
[**ApiV1LibrariesPost**](UploadsAndAudioContentApi.md#ApiV1LibrariesPost) | **Post** /api/v1/libraries/ | 
[**ApiV1LibrariesUuidDelete**](UploadsAndAudioContentApi.md#ApiV1LibrariesUuidDelete) | **Delete** /api/v1/libraries/{uuid}/ | Delete a library and all associated uploads
[**ApiV1LibrariesUuidGet**](UploadsAndAudioContentApi.md#ApiV1LibrariesUuidGet) | **Get** /api/v1/libraries/{uuid}/ | Retrieve a library
[**ApiV1LibrariesUuidPost**](UploadsAndAudioContentApi.md#ApiV1LibrariesUuidPost) | **Post** /api/v1/libraries/{uuid}/ | Update a library
[**ApiV1UploadsGet**](UploadsAndAudioContentApi.md#ApiV1UploadsGet) | **Get** /api/v1/uploads/ | List owned uploads
[**ApiV1UploadsPost**](UploadsAndAudioContentApi.md#ApiV1UploadsPost) | **Post** /api/v1/uploads/ | 
[**ApiV1UploadsUuidAudioFileMetadataGet**](UploadsAndAudioContentApi.md#ApiV1UploadsUuidAudioFileMetadataGet) | **Get** /api/v1/uploads/{uuid}/audio-file-metadata | Retrieve the tags embedded in the audio file
[**ApiV1UploadsUuidDelete**](UploadsAndAudioContentApi.md#ApiV1UploadsUuidDelete) | **Delete** /api/v1/uploads/{uuid}/ | Delete an upload
[**ApiV1UploadsUuidGet**](UploadsAndAudioContentApi.md#ApiV1UploadsUuidGet) | **Get** /api/v1/uploads/{uuid}/ | Retrieve an upload
[**ApiV1UploadsUuidPatch**](UploadsAndAudioContentApi.md#ApiV1UploadsUuidPatch) | **Patch** /api/v1/uploads/{uuid}/ | Update a draft upload



## ApiV1ChannelsGet

> ResultPage ApiV1ChannelsGet(ctx, )

List channels

### Required Parameters

This endpoint does not need any parameter.

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


## ApiV1ChannelsPost

> ApiV1ChannelsPost(ctx, body)

Create a new channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **ChannelCreate**|  | 

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


## ApiV1ChannelsUuidDelete

> ApiV1ChannelsUuidDelete(ctx, uuid)

Delete a channel and all associated uploads

This will delete the channel, all associated uploads, follows, and broadcast the event on the federation. 

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


## ApiV1ChannelsUuidGet

> Channel ApiV1ChannelsUuidGet(ctx, uuid)

Retrieve a channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 

### Return type

[**Channel**](Channel.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ChannelsUuidPost

> Channel ApiV1ChannelsUuidPost(ctx, uuid, body)

Update a channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 
**body** | **ChannelUpdate**|  | 

### Return type

[**Channel**](Channel.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1LibrariesGet

> ResultPage ApiV1LibrariesGet(ctx, )

List owned libraries

### Required Parameters

This endpoint does not need any parameter.

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


## ApiV1LibrariesPost

> ApiV1LibrariesPost(ctx, body)



Create a new library

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **OwnedLibraryCreate**|  | 

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


## ApiV1LibrariesUuidDelete

> ApiV1LibrariesUuidDelete(ctx, uuid)

Delete a library and all associated uploads

This will delete the library, all associated uploads, follows, and broadcast the event on the federation. 

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


## ApiV1LibrariesUuidGet

> OwnedLibrary ApiV1LibrariesUuidGet(ctx, uuid)

Retrieve a library

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 

### Return type

[**OwnedLibrary**](OwnedLibrary.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1LibrariesUuidPost

> OwnedLibrary ApiV1LibrariesUuidPost(ctx, uuid, body)

Update a library

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 
**body** | **OwnedLibraryCreate**|  | 

### Return type

[**OwnedLibrary**](OwnedLibrary.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UploadsGet

> ResultPage ApiV1UploadsGet(ctx, optional)

List owned uploads

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiV1UploadsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1UploadsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Search query used to filter uploads | 

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


## ApiV1UploadsPost

> ApiV1UploadsPost(ctx, optional)



Upload a new file in a library. The event will be broadcasted on federation, according to the library visibility and followers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiV1UploadsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1UploadsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **library** | [**optional.Interface of string**](string.md)| The library in which the audio should be stored | 
 **importReference** | **optional.String**|  | 
 **source** | **optional.String**|  | 
 **audioFile** | **optional.Interface of *os.File****optional.*os.File**|  | 
 **importStatus** | **optional.String**| Setting import_status to draft will prevent processing, but allow further modifications to audio and metadata. Once ready, use the PATCH method to set import_status to pending. Default to &#x60;pending&#x60; if unspecified. | [default to pending]
 **importMetadata** | [**optional.Interface of ImportMetadata**](ImportMetadata.md)|  | 

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


## ApiV1UploadsUuidAudioFileMetadataGet

> map[string]interface{} ApiV1UploadsUuidAudioFileMetadataGet(ctx, uuid)

Retrieve the tags embedded in the audio file

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UploadsUuidDelete

> ApiV1UploadsUuidDelete(ctx, uuid)

Delete an upload

This will delete the upload from the server and broadcast the event on the federation. 

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


## ApiV1UploadsUuidGet

> OwnedUpload ApiV1UploadsUuidGet(ctx, uuid)

Retrieve an upload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 

### Return type

[**OwnedUpload**](OwnedUpload.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UploadsUuidPatch

> OwnedUpload ApiV1UploadsUuidPatch(ctx, uuid)

Update a draft upload

This will update a draft upload, before it is processed.  All fields supported for `POST /api/v1/uploads` can be updated here.  Setting `import_status` to `pending` will trigger processing, and make future modifications impossible. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)|  | 

### Return type

[**OwnedUpload**](OwnedUpload.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

