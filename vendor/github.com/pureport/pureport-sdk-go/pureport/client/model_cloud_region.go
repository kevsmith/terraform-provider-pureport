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

// CloudRegion struct for CloudRegion
type CloudRegion struct {
	DisplayName        string         `json:"displayName"`
	GeoCoordinates     GeoCoordinates `json:"geoCoordinates,omitempty"`
	Href               string         `json:"href,omitempty"`
	Id                 string         `json:"id,omitempty"`
	Provider           string         `json:"provider"`
	ProviderAssignedId string         `json:"providerAssignedId"`
}
