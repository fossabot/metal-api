package metal

import (
	"reflect"
	"testing"
	"time"
)

// Gerrit and me implemented that monster in a one shot which worked.
func TestSwitch_ConnectDevice(t *testing.T) {
	type fields struct {
		ID          string
		Nics        []Nic
		Connections []Connection
		SiteID      string
		Created     time.Time
		Changed     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		device *Device
	}{
		{
			name: "simple connection",
			fields: fields{
				ID: "switch-1",
				Nics: []Nic{
					Nic{
						Name:       "eth0",
						MacAddress: "00:11:11",
					},
					Nic{
						Name:       "swp1",
						MacAddress: "11:11:11",
					},
					Nic{
						Name:       "swp2",
						MacAddress: "22:11:11",
					},
				},
				SiteID: "nbg1",
				Connections: []Connection{
					Connection{
						Nic: Nic{
							Name:       "swp1",
							MacAddress: "11:11:11",
						},
						DeviceID: "device-1",
					},
					Connection{
						Nic: Nic{
							Name:       "swp2",
							MacAddress: "22:11:11",
						},
						DeviceID: "device-1",
					},
				},
			},
			device: &Device{
				ID: "device-1",
				Hardware: DeviceHardware{
					Nics: []Nic{
						Nic{
							Name: "eth0",
							Neighbors: []Nic{
								Nic{
									MacAddress: "11:11:11",
								},
								Nic{
									MacAddress: "11:11:12",
								},
							},
						},
						Nic{
							Name: "eth1",
							Neighbors: []Nic{
								Nic{
									MacAddress: "22:11:11",
								},
								Nic{
									MacAddress: "11:11:13",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Switch{
				ID:      tt.fields.ID,
				Nics:    tt.fields.Nics,
				SiteID:  tt.fields.SiteID,
				Created: tt.fields.Created,
				Changed: tt.fields.Changed,
			}
			s.ConnectDevice(tt.device)
			if !reflect.DeepEqual(s.Connections, tt.fields.Connections) {
				t.Errorf("expected:%v, got:%v", s.Connections, tt.fields.Connections)
			}
		})
	}
}
