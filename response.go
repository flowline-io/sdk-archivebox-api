/*
ArchiveBox API

 <h3>Welcome to your ArchiveBox server's REST API <code>[v1 ALPHA]</code> homepage!</h3> <br/> <i><b>WARNING: This API is still in an early development stage and may change!</b></i> <br/> <ul> <li>⬅️ Manage your server: <a href=\"/admin/api/\"><b>Setup API Keys</b></a>, <a href=\"/admin/\">Go to your Server Admin UI</a>, <a href=\"/\">Go to your Snapshots list</a>  <li>💬 Ask questions and get help here: <a href=\"https://zulip.archivebox.io\">ArchiveBox Chat Forum</a></li> <li>🐞 Report API bugs here: <a href=\"https://github.com/ArchiveBox/ArchiveBox/issues\">Github Issues</a></li> <li>📚 ArchiveBox Documentation: <a href=\"https://github.com/ArchiveBox/ArchiveBox/wiki\">Github Wiki</a></li> <li>📜 See the API source code: <a href=\"https://github.com/ArchiveBox/ArchiveBox/blob/dev/archivebox/api\"><code>archivebox/api/</code></a></li> </ul> <small>Served by ArchiveBox v0.8.5rc50 (<a href=\"https://github.com/ArchiveBox/ArchiveBox/commit/1dff8bae4043efb678c1a6fb517dfa7e7a211150\"><code>1dff8bae</code></a>), API powered by <a href=\"https://django-ninja.dev/\"><code>django-ninja</code></a>.</small> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"net/http"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

// NewAPIResponse returns a new APIResponse object.
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError returns a new APIResponse object with the provided error message.
func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
