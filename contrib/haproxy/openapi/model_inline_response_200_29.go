/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InlineResponse20029 struct for InlineResponse20029
type InlineResponse20029 struct {
	Version int32     `json:"_version,omitempty"`
	Data    LogTarget `json:"data,omitempty"`
}
