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
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")

)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}


// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	Url string
	Description string
	Variables map[string]ServerVariable
}

// Configuration stores the configuration of the API client
type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	Servers       []ServerConfiguration
	HTTPClient    *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "https://demo.funkwhale.audio",
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers:       []ServerConfiguration{
			{
				Url: "https://demo.funkwhale.audio",
				Description: "Demo server",
			},
			{
				Url: "https://{domain}",
				Description: "Custom server",
				Variables: map[string]ServerVariable{
					"domain": ServerVariable{
						Description: "Your Funkwhale Domain",
						DefaultValue: "yourdomain",
					},
					"protocol": ServerVariable{
						Description: "No description provided",
						DefaultValue: "https",
						EnumValues: []string{
							"http",
							"https",
						},
					},
				},
			},
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// ServerUrl returns URL based on server settings
func (c *Configuration) ServerUrl(index int, variables map[string]string) (string, error) {
	if index < 0 || len(c.Servers) <= index {
		return "", fmt.Errorf("Index %v out of range %v", index, len(c.Servers) - 1)
	}
	server := c.Servers[index]
	url := server.Url

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("The variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}
