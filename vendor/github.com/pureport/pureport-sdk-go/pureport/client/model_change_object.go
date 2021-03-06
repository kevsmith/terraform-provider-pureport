/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Contact: support@pureport.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// ChangeObject struct for ChangeObject
type ChangeObject struct {
	Current  map[string]interface{} `json:"current,omitempty"`
	Previous map[string]interface{} `json:"previous,omitempty"`
	Property string                 `json:"property,omitempty"`
}
