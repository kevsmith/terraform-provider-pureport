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

import (
	"time"
)

// ConnectionTimeEgressIngress struct for ConnectionTimeEgressIngress
type ConnectionTimeEgressIngress struct {
	Connection Link      `json:"connection,omitempty"`
	Egress     int64     `json:"egress,omitempty"`
	Ingress    int64     `json:"ingress,omitempty"`
	Time       time.Time `json:"time,omitempty"`
}