# \LibraryAndMetadataApi

All URIs are relative to *https://demo.funkwhale.audio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AlbumsGet**](LibraryAndMetadataApi.md#ApiV1AlbumsGet) | **Get** /api/v1/albums/ | List albums
[**ApiV1AlbumsIdGet**](LibraryAndMetadataApi.md#ApiV1AlbumsIdGet) | **Get** /api/v1/albums/{id}/ | Retrieve a single album
[**ApiV1AlbumsIdLibrariesGet**](LibraryAndMetadataApi.md#ApiV1AlbumsIdLibrariesGet) | **Get** /api/v1/albums/{id}/libraries/ | List available user libraries containing tracks from this album
[**ApiV1ArtistsGet**](LibraryAndMetadataApi.md#ApiV1ArtistsGet) | **Get** /api/v1/artists/ | List artists
[**ApiV1ArtistsIdGet**](LibraryAndMetadataApi.md#ApiV1ArtistsIdGet) | **Get** /api/v1/artists/{id}/ | Retrieve a single artist
[**ApiV1ArtistsIdLibrariesGet**](LibraryAndMetadataApi.md#ApiV1ArtistsIdLibrariesGet) | **Get** /api/v1/artists/{id}/libraries/ | List available user libraries containing work from this artist
[**ApiV1LicensesCodeGet**](LibraryAndMetadataApi.md#ApiV1LicensesCodeGet) | **Get** /api/v1/licenses/{code}/ | Retrieve a single license
[**ApiV1LicensesGet**](LibraryAndMetadataApi.md#ApiV1LicensesGet) | **Get** /api/v1/licenses/ | List licenses
[**ApiV1ListenUuidGet**](LibraryAndMetadataApi.md#ApiV1ListenUuidGet) | **Get** /api/v1/listen/{uuid}/ | Download the audio file matching the given track uuid
[**ApiV1TracksGet**](LibraryAndMetadataApi.md#ApiV1TracksGet) | **Get** /api/v1/tracks/ | List tracks
[**ApiV1TracksIdGet**](LibraryAndMetadataApi.md#ApiV1TracksIdGet) | **Get** /api/v1/tracks/{id}/ | Retrieve a single track
[**ApiV1TracksIdLibrariesGet**](LibraryAndMetadataApi.md#ApiV1TracksIdLibrariesGet) | **Get** /api/v1/tracks/{id}/libraries/ | List available user libraries containing given track



## ApiV1AlbumsGet

> ResultPage ApiV1AlbumsGet(ctx, optional)

List albums

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiV1AlbumsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1AlbumsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Search query used to filter albums | 
 **artist** | **optional.Int64**| Only include albums by the requested artist | 

### Return type

[**ResultPage**](ResultPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AlbumsIdGet

> Album ApiV1AlbumsIdGet(ctx, )

Retrieve a single album

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Album**](Album.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AlbumsIdLibrariesGet

> LibraryPage ApiV1AlbumsIdLibrariesGet(ctx, )

List available user libraries containing tracks from this album

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LibraryPage**](LibraryPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ArtistsGet

> ResultPage ApiV1ArtistsGet(ctx, optional)

List artists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiV1ArtistsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1ArtistsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Search query used to filter artists | 

### Return type

[**ResultPage**](ResultPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ArtistsIdGet

> Artist ApiV1ArtistsIdGet(ctx, )

Retrieve a single artist

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Artist**](Artist.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ArtistsIdLibrariesGet

> LibraryPage ApiV1ArtistsIdLibrariesGet(ctx, )

List available user libraries containing work from this artist

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LibraryPage**](LibraryPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1LicensesCodeGet

> License ApiV1LicensesCodeGet(ctx, code)

Retrieve a single license

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**| License code | 

### Return type

[**License**](License.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1LicensesGet

> ResultPage ApiV1LicensesGet(ctx, )

List licenses

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ResultPage**](ResultPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ListenUuidGet

> *os.File ApiV1ListenUuidGet(ctx, uuid, optional)

Download the audio file matching the given track uuid

Given a track uuid (and not ID), return the first found audio file accessible by the user making the request.  In case of a remote upload, this endpoint will fetch the audio file from the remote and cache it before sending the response. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| Track uuid | 
 **optional** | ***ApiV1ListenUuidGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1ListenUuidGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | [**optional.Interface of map[string]interface{}**](.md)| If specified, the endpoint will return a transcoded version of the original audio file.  Since transcoding happens on the fly, it can significantly increase response time, and it&#39;s recommended to request transcoding only for files that are not playable by the client.  This endpoint support bytess-range requests.  | 
 **upload** | [**optional.Interface of string**](.md)| If specified, will return the audio for the given upload uuid.  This is useful for tracks that have multiple uploads available.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TracksGet

> ResultPage ApiV1TracksGet(ctx, optional)

List tracks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiV1TracksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiV1TracksGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| Search query used to filter tracks | 
 **artist** | **optional.Int64**| Only include tracks by the requested artist | 
 **favorites** | **optional.Bool**| filter/exclude tracks favorited by the current user | 
 **album** | **optional.Int64**| Only include tracks from the requested album | 
 **license** | **optional.String**| Only include tracks with the given license | 

### Return type

[**ResultPage**](ResultPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TracksIdGet

> Track ApiV1TracksIdGet(ctx, )

Retrieve a single track

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Track**](Track.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TracksIdLibrariesGet

> LibraryPage ApiV1TracksIdLibrariesGet(ctx, )

List available user libraries containing given track

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LibraryPage**](LibraryPage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

