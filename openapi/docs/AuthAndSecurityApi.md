# \AuthAndSecurityApi

All URIs are relative to *https://demo.funkwhale.audio*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AuthPasswordResetPost**](AuthAndSecurityApi.md#ApiV1AuthPasswordResetPost) | **Post** /api/v1/auth/password/reset/ | Request a password reset
[**ApiV1AuthRegistrationPost**](AuthAndSecurityApi.md#ApiV1AuthRegistrationPost) | **Post** /api/v1/auth/registration/ | Create an account
[**ApiV1OauthAppsPost**](AuthAndSecurityApi.md#ApiV1OauthAppsPost) | **Post** /api/v1/oauth/apps/ | 
[**ApiV1RateLimitGet**](AuthAndSecurityApi.md#ApiV1RateLimitGet) | **Get** /api/v1/rate-limit/ | Retrive rate-limit information and current usage status
[**ApiV1TokenPost**](AuthAndSecurityApi.md#ApiV1TokenPost) | **Post** /api/v1/token/ | Get an API token
[**ApiV1UsersUsersMeGet**](AuthAndSecurityApi.md#ApiV1UsersUsersMeGet) | **Get** /api/v1/users/users/me/ | Retrive profile information



## ApiV1AuthPasswordResetPost

> ApiV1AuthPasswordResetPost(ctx, inlineObject3)

Request a password reset

Request a password reset. An email with reset instructions will be sent to the provided email, if it's associated with a user account. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject3** | [**InlineObject3**](InlineObject3.md)|  | 

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


## ApiV1AuthRegistrationPost

> ApiV1AuthRegistrationPost(ctx, inlineObject2)

Create an account

Register a new account on this instance. An invitation code will be required if sign up is disabled. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject2** | [**InlineObject2**](InlineObject2.md)|  | 

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


## ApiV1OauthAppsPost

> OAuthApplication ApiV1OauthAppsPost(ctx, inlineObject)



Register an OAuth application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**OAuthApplication**](OAuthApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RateLimitGet

> RateLimitStatus ApiV1RateLimitGet(ctx, )

Retrive rate-limit information and current usage status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RateLimitStatus**](RateLimitStatus.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TokenPost

> ApiV1TokenPost(ctx, inlineObject1)

Get an API token

Obtain a JWT token you can use for authenticating your next requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inlineObject1** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersUsersMeGet

> Me ApiV1UsersUsersMeGet(ctx, )

Retrive profile information

Retrieve profile informations of the current user 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Me**](Me.md)

### Authorization

[jwt](../README.md#jwt), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

