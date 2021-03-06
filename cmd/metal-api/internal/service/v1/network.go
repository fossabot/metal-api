package v1

import (
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/datastore"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
)

type NetworkBase struct {
	PartitionID *string           `json:"partitionid" description:"the partition this network belongs to" optional:"true"`
	ProjectID   *string           `json:"projectid" description:"the project id this network belongs to, can be empty if globally available" optional:"true"`
	Labels      map[string]string `json:"labels" description:"free labels that you associate with this network."`
}

type NetworkImmutable struct {
	Prefixes            []string `json:"prefixes" modelDescription:"a network which contains prefixes from which IP addresses can be allocated" description:"the prefixes of this network"`
	DestinationPrefixes []string `json:"destinationprefixes" modelDescription:"prefixes that are reachable within this network" description:"the destination prefixes of this network"`
	Nat                 bool     `json:"nat" description:"if set to true, packets leaving this network get masqueraded behind interface ip"`
	PrivateSuper        bool     `json:"privatesuper" description:"if set to true, this network will serve as a partition's super network for the internal machine networks,there can only be one privatesuper network per partition"`
	Underlay            bool     `json:"underlay" description:"if set to true, this network can be used for underlay communication"`
	Vrf                 *uint    `json:"vrf" description:"the vrf this network is associated with" optional:"true"`
	VrfShared           *bool    `json:"vrfshared" description:"if set to true, given vrf can be used by multiple networks, which is sometimes useful for network partioning (default: false)" optional:"true"`
	ParentNetworkID     *string  `json:"parentnetworkid" description:"the id of the parent network"`
}

type NetworkUsage struct {
	AvailableIPs      uint64 `json:"available_ips" description:"the total available IPs" readonly:"true"`
	UsedIPs           uint64 `json:"used_ips" description:"the total used IPs" readonly:"true"`
	AvailablePrefixes uint64 `json:"available_prefixes" description:"the total available Prefixes" readonly:"true"`
	UsedPrefixes      uint64 `json:"used_prefixes" description:"the total used Prefixes" readonly:"true"`
}

type NetworkCreateRequest struct {
	ID *string `json:"id" description:"the unique ID of this entity, auto-generated if left empty" unique:"true"`
	Describable
	NetworkBase
	NetworkImmutable
}

type NetworkAllocateRequest struct {
	Describable
	NetworkBase
}

type NetworkFindRequest struct {
	datastore.NetworkSearchQuery
}

type NetworkUpdateRequest struct {
	Common
	Prefixes []string `json:"prefixes" description:"the prefixes of this network" optional:"true"`
}

type NetworkResponse struct {
	Common
	NetworkBase
	NetworkImmutable
	Usage NetworkUsage `json:"usage" description:"usage of ips and prefixes in this network" readonly:"true"`
	Timestamps
}

func NewNetworkResponse(network *metal.Network, usage *metal.NetworkUsage) *NetworkResponse {
	if network == nil {
		return nil
	}

	var parentNetworkID *string
	if network.ParentNetworkID != "" {
		parentNetworkID = &network.ParentNetworkID
	}

	return &NetworkResponse{
		Common: Common{
			Identifiable: Identifiable{
				ID: network.ID,
			},
			Describable: Describable{
				Name:        &network.Name,
				Description: &network.Description,
			},
		},
		NetworkBase: NetworkBase{
			PartitionID: &network.PartitionID,
			ProjectID:   &network.ProjectID,
			Labels:      network.Labels,
		},
		NetworkImmutable: NetworkImmutable{
			Prefixes:            network.Prefixes.String(),
			DestinationPrefixes: network.DestinationPrefixes.String(),
			Nat:                 network.Nat,
			PrivateSuper:        network.PrivateSuper,
			Underlay:            network.Underlay,
			Vrf:                 &network.Vrf,
			ParentNetworkID:     parentNetworkID,
		},
		Usage: NetworkUsage{
			AvailableIPs:      usage.AvailableIPs,
			UsedIPs:           usage.UsedIPs,
			AvailablePrefixes: usage.AvailablePrefixes,
			UsedPrefixes:      usage.UsedPrefixes,
		},
		Timestamps: Timestamps{
			Created: network.Created,
			Changed: network.Changed,
		},
	}
}
