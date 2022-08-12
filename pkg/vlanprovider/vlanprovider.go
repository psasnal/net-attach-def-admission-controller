package vlanprovider

import (
	"fmt"
)

type Nic struct {
	Name       string `json:"name"`
	MacAddress string `json:"mac-address"`
}

// NIC in JSON format
type JsonNic map[string]interface{}
type NicMap map[string]JsonNic

type NodeTopology struct {
	Bonds      map[string]NicMap
	SriovPools map[string]NicMap
}

type VlanProvider interface {
	Connect() error
	UpdateNodeTopology(string, string) (string, error)
	Attach(string, int, []string) error
	Detach(string, int, []string) error
}

func NewVlanProvider(provider string, config string) (VlanProvider, error) {
	switch provider {
	case "openstack":
		{
			openstack := &OpenstackVlanProvider{
				configFile: config}
			err := openstack.Connect()
			return openstack, err
		}
	case "baremetal":
		{
			fss := &FssVlanProvider{
				configFile: config}
			err := fss.Connect()
			return fss, err
		}
	default:
		return nil, fmt.Errorf("Not supported provider: %q", provider)
	}

}
