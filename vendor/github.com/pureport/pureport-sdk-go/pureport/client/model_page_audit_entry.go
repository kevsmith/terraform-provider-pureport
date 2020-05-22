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

// PageAuditEntry struct for PageAuditEntry
type PageAuditEntry struct {
	Content       []AuditEntry `json:"content,omitempty"`
	PageNumber    int32        `json:"pageNumber,omitempty"`
	PageSize      int32        `json:"pageSize,omitempty"`
	Sort          Sort         `json:"sort,omitempty"`
	TotalElements int32        `json:"totalElements,omitempty"`
}