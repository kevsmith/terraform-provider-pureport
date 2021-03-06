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

// PacketFabricConnection Connection using Packet Fabric as an access point provider.
type PacketFabricConnection struct {
	ActiveAt time.Time `json:"activeAt,omitempty"`
	// If the connection is advertising internal routes, which allows the customer the option of probing and tracing these routes.
	AdvertiseInternalRoutes bool         `json:"advertiseInternalRoutes,omitempty"`
	BillingPlan             *BillingPlan `json:"billingPlan,omitempty"`
	// The provider used for billing this connection.
	BillingProvider string `json:"billingProvider,omitempty"`
	// The licensed billing term for the connection.
	BillingTerm string    `json:"billingTerm"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	// The customer side ASN. This can either be a public or private ASN. If this is a public ASN, you must own it to prevent conflicts.
	CustomerASN int64 `json:"customerASN,omitempty"`
	// Set of customer Networks for this connection.
	CustomerNetworks []CustomerNetwork `json:"customerNetworks,omitempty"`
	DeletedAt        time.Time         `json:"deletedAt,omitempty"`
	// The user defined description for the connection.
	Description string `json:"description,omitempty"`
	// Error Code assigned to the connection if it is an error state.
	ErrorCode string `json:"errorCode,omitempty"`
	// Error message assigned to the connection if it is an error state.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether this connection has redundant gateways for failover.
	HighAvailability bool   `json:"highAvailability"`
	Href             string `json:"href,omitempty"`
	// The id is a unique identifier representing the connection. This can be provided during creation, but if left empty, will be generated.
	Id       string `json:"id,omitempty"`
	Location Link   `json:"location"`
	// The user specified name for the connection.
	Name             string           `json:"name"`
	Nat              *NatConfig       `json:"nat,omitempty"`
	Network          Link             `json:"network,omitempty"`
	PrimaryGateway   *StandardGateway `json:"primaryGateway,omitempty"`
	SecondaryGateway *StandardGateway `json:"secondaryGateway,omitempty"`
	// The connection speed in Mbps.
	Speed int32 `json:"speed"`
	// The current state of the connection.
	State string            `json:"state,omitempty"`
	Tags  map[string]string `json:"tags,omitempty"`
	// The connection type.
	Type                     string                    `json:"type"`
	BgpPasswordConfiguration *BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	CloudRegion              Link                      `json:"cloudRegion"`
	Customer                 PacketFabricCustomer      `json:"customer"`
	// An ID representing the primary connection with the 3rd party provider selected in the “customer” field.
	CustomerPrimaryId string `json:"customerPrimaryId"`
	// An ID representing the secondary connection with the 3rd party provider selected in the “customer” field.
	CustomerSecondaryId string                `json:"customerSecondaryId,omitempty"`
	Peering             *PeeringConfiguration `json:"peering"`
	// The primary VLAN ID.
	PrimaryVlan int32 `json:"primaryVlan,omitempty"`
	// The secondary VLAN ID if this is an HA connection.
	SecondaryVlan int32 `json:"secondaryVlan,omitempty"`
}
