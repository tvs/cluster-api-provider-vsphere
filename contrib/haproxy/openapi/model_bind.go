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

// Bind HAProxy frontend bind configuration
type Bind struct {
	AcceptProxy    bool   `json:"accept_proxy,omitempty"`
	Address        string `json:"address,omitempty"`
	Alpn           string `json:"alpn,omitempty"`
	Name           string `json:"name"`
	Port           *int32 `json:"port,omitempty"`
	Process        string `json:"process,omitempty"`
	Ssl            bool   `json:"ssl,omitempty"`
	SslCafile      string `json:"ssl_cafile,omitempty"`
	SslCertificate string `json:"ssl_certificate,omitempty"`
	TcpUserTimeout *int32 `json:"tcp_user_timeout,omitempty"`
	Transparent    bool   `json:"transparent,omitempty"`
	V4v6           bool   `json:"v4v6,omitempty"`
	Verify         string `json:"verify,omitempty"`
}
