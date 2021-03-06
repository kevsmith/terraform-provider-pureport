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

// AccessSwitch struct for AccessSwitch
type AccessSwitch struct {
	AvailabilityDomain string `json:"availabilityDomain"`
	Href               string `json:"href,omitempty"`
	Id                 string `json:"id"`
	Pod                Link   `json:"pod"`
	Status             string `json:"status"`
}
