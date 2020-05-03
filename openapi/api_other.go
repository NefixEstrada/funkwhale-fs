/*
 * Funkwhale API
 *
 * Interactive documentation for [Funkwhale](https://funkwhale.audio) API.  The API is **not** freezed yet, but we will document breaking changes in our changelog, and try to avoid those as much as possible.  Usage -----  Click on an endpoint name to inspect its properties, parameters and responses.  Use the \"Try it out\" button to send a real world payload to the endpoint and inspect the corresponding response.  Authentication --------------  To authenticate, use the `/token/` endpoint with a username and password, and copy/paste the resulting JWT token in the `Authorize` modal. All subsequent requests made via the interactive documentation will be authenticated.  If you keep the default server (https://demo.funkwhale.audio), the default username and password couple is \"demo\" and \"demo\".  Rate limiting -------------  Depending on server configuration, pods running Funkwhale 0.20 and higher may rate-limit incoming requests to prevent abuse and improve the stability of service. Requests that are dropped because of rate-limiting receive a 429 HTTP response.  The limits themselves vary depending on:  - The client: anonymous requests are subject to lower limits than authenticated requests - The operation being performed: Write and delete operations, as performed with DELETE, POST, PUT and PATCH HTTP methods are subject to lower limits  Those conditions are used to determine the scope of the request, which in turns determine the limit that is applied. For instance, authenticated POST requests are bound to the `authenticated-create` scope, with a default limit of 1000 requests/hour, but anonymous POST requests are bound to the `anonymous-create` scope, with a lower limit of 1000 requests/day.  A full list of scopes with their corresponding description, and the current usage data for the client performing the request is available via the `/api/v1/rate-limit` endpoint.  Additionally, we include HTTP headers on all API response to ensure API clients can understand:  - what scope was bound to a given request - what is the corresponding limit - how much similar requests can be sent before being limited - and how much time they should wait if they have been limited  <table>   <caption>Rate limiting headers</caption>   <thead>     <th>Header</th>     <th>Example value</th>     <th>Description value</th>   </thead>   <tbody>     <tr>       <td><code>X-RateLimit-Limit</code></td>       <td>50</td>       <td>The number of allowed requests whithin a given period</td>     </tr>     <tr>       <td><code>X-RateLimit-Duration</code></td>       <td>3600</td>       <td>The time window, in seconds, during which those requests are accounted for.</td>     </tr>     <tr>       <td><code>X-RateLimit-Scope</code></td>       <td>login</td>       <td>The name of the scope as computed for the request</td>     </tr>     <tr>       <td><code>X-RateLimit-Remaining</code></td>       <td>42</td>       <td>How many requests can be sent with the same scope before the limit applies</td>     </tr>     <tr>       <td><code>Retry-After</code> (if <code>X-RateLimit-Remaining</code> is 0)</td>       <td>3543</td>       <td>How many seconds to wait before a retry</td>     </tr>     <tr>       <td><code>X-RateLimit-Reset</code></td>       <td>1568126089</td>       <td>A timestamp indicating when <code>X-RateLimit-Remaining</code> will return to its higher possible value</td>     </tr>     <tr>       <td><code>X-RateLimit-ResetSeconds</code></td>       <td>3599</td>       <td>How many seconds to wait before <code>X-RateLimit-Remaining</code> returns to its higher possible value</td>     </tr>   </tbody> </table>   Resources ---------  For more targeted guides regarding API usage, and especially authentication, please refer to [https://docs.funkwhale.audio/api.html](https://docs.funkwhale.audio/api.html) 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"os"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// OtherApiService OtherApi service
type OtherApiService service

// ApiV1AttachmentsPostOpts Optional parameters for the method 'ApiV1AttachmentsPost'
type ApiV1AttachmentsPostOpts struct {
    File optional.Interface
}

/*
ApiV1AttachmentsPost Method for ApiV1AttachmentsPost
Upload a new file as an attachment that can be later associated with other objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ApiV1AttachmentsPostOpts - Optional Parameters:
 * @param "File" (optional.Interface of *os.File) - 
*/
func (a *OtherApiService) ApiV1AttachmentsPost(ctx _context.Context, localVarOptionals *ApiV1AttachmentsPostOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/attachments/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormFileName = "file"
	var localVarFile *os.File
	if localVarOptionals != nil && localVarOptionals.File.IsSet() {
		localVarFileOk := false
		localVarFile, localVarFileOk = localVarOptionals.File.Value().(*os.File)
		if !localVarFileOk {
				return nil, reportError("file should be *os.File")
		}
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
ApiV1AttachmentsUuidDelete Delete an attachment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uuid
*/
func (a *OtherApiService) ApiV1AttachmentsUuidDelete(ctx _context.Context, uuid string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/attachments/{uuid}/"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", _neturl.QueryEscape(parameterToString(uuid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

/*
ApiV1AttachmentsUuidGet Retrieve an attachment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uuid
@return Attachment
*/
func (a *OtherApiService) ApiV1AttachmentsUuidGet(ctx _context.Context, uuid string) (Attachment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Attachment
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/attachments/{uuid}/"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", _neturl.QueryEscape(parameterToString(uuid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
