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

// DetailedInvoiceItem struct for DetailedInvoiceItem
type DetailedInvoiceItem struct {
	Account            Link                 `json:"account,omitempty"`
	BillableObject     Link                 `json:"billableObject,omitempty"`
	BillableObjectType string               `json:"billableObjectType,omitempty"`
	BillingPlan        BillingPlan          `json:"billingPlan,omitempty"`
	Description        string               `json:"description,omitempty"`
	Quantity           float32              `json:"quantity,omitempty"`
	TimeRanges         []ClosedRangeInstant `json:"timeRanges,omitempty"`
}