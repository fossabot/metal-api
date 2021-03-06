package metal

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddMachineId(t *testing.T) {
	tests := []struct {
		name         string
		ip           IP
		expectedTags []string
	}{
		{
			name:         "ip without machine tag",
			ip:           IP{},
			expectedTags: []string{IpTag(TagIPMachineID, "123")},
		},
		{
			name: "ip with empty machine tag",
			ip: IP{
				Tags: []string{TagIPMachineID},
			},
			expectedTags: []string{IpTag(TagIPMachineID, "123")},
		},
		{
			name: "ip with other machine tag",
			ip: IP{
				Tags: []string{IpTag(TagIPMachineID, "1")},
			},
			expectedTags: []string{IpTag(TagIPMachineID, "1"), IpTag(TagIPMachineID, "123")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ip.AddMachineId("123")
			if got := tt.ip.Tags; !cmp.Equal(got, tt.expectedTags) {
				t.Errorf("%v", cmp.Diff(got, tt.expectedTags))
			}
		})
	}
}
func TestRemoveMachineId(t *testing.T) {
	tests := []struct {
		name         string
		ip           IP
		expectedTags []string
	}{
		{
			name:         "ip without machine tag",
			ip:           IP{},
			expectedTags: []string{},
		},
		{
			name: "ip with empty machine tag",
			ip: IP{
				Tags: []string{TagIPMachineID},
			},
			expectedTags: []string{TagIPMachineID},
		},
		{
			name: "ip with other machine tag",
			ip: IP{
				Tags: []string{IpTag(TagIPMachineID, "1")},
			},
			expectedTags: []string{IpTag(TagIPMachineID, "1")},
		},
		{
			name: "ip with matching machine tag",
			ip: IP{
				Tags: []string{IpTag(TagIPMachineID, "123")},
			},
			expectedTags: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ip.RemoveMachineId("123")
			if got := tt.ip.Tags; !cmp.Equal(got, tt.expectedTags) {
				t.Errorf("%v", cmp.Diff(got, tt.expectedTags))
			}
		})
	}
}

func TestGetScope(t *testing.T) {
	tests := []struct {
		name          string
		ip            IP
		expectedScope IPScope
	}{
		{
			name: "empty scope ip",
			ip: IP{
				Tags: []string{IpTag(TagIPMachineID, "102")},
			},
			expectedScope: ScopeEmpty,
		},
		{
			name: "machine ip",
			ip: IP{
				ProjectID: "1",
				Tags:      []string{IpTag(TagIPMachineID, "102")},
			},
			expectedScope: ScopeMachine,
		},
		{
			name: "project ip",
			ip: IP{
				ProjectID: "1",
				Tags:      []string{},
			},
			expectedScope: ScopeProject,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ip.GetScope(); got != tt.expectedScope {
				t.Errorf("IP.GetScope = %v, want %v", got, tt.expectedScope)
			}
		})
	}
}

func TestIPToASN(t *testing.T) {
	ipaddress := IP{
		IPAddress: "10.0.1.2",
	}

	asn, err := ipaddress.ASN()
	if err != nil {
		t.Errorf("no error expected got:%v", err)
	}

	if asn != 4200000258 {
		t.Errorf("expected 4200000258 got: %d", asn)
	}
}
