package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type VM struct {
	Vcent        Vcenter  `yaml:"vCenterDetails"`
	Vclone       VmClone  `yaml:"vmCloneProperties"`
	VmProperties []VmProp `yaml:"vmProperties"`
}

type Vcenter struct {
	VmURL        string `yaml:"vmUrl"`
	VmUsername   string `yaml:"vmUsername"`
	VmPassword   string `yaml:"vmPassword"`
	VmInseucre   string `yaml:"vmInsecure"`
	VmDatastore  string `yaml:"vmDatastore"`
	VmDatacenter string `yaml:"vmDatacenter"`
	VmFolder     string `yaml:"vmFolder"`
}

type VmClone struct {
	RhcosTemplate string `yaml:"rhcosTemplate"`
	CpuNumber     int    `yaml:"cpuNumber"`
	MemoryMB      int    `yaml:"memoryMB"`
	DiskSizeGB    int    `yaml:"diskSizeGB"`
}

type VmProp struct {
	VMName          string `yaml:"vmName"`
	NetInterface    string `yaml:"netInterface"`
	IpAddress       string `yaml:"ipAddress"`
	DefaultGateway  string `yaml:"defaultGateway"`
	Dns             string `yaml:"dns"`
	IpInterfaceName string `yaml:"ipInterfaceName"`
	MacAddress      string `yaml:"macAddress"`
	IgnFile         string `yaml:"ignFile"`
}

var VmData VM

func ConvertYamlToStruct() map[string]string{
	rbs, err := ioutil.ReadFile("vm.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(rbs, &VmData)
	if err != nil {
		fmt.Println(err)
	}
	var Goenv = map[string]string{
		"GOVC_URL":        VmData.Vcent.VmURL,
		"GOVC_USERNAME":   VmData.Vcent.VmUsername,
		"GOVC_PASSWORD":   VmData.Vcent.VmPassword,
		"GOVC_INSECURE":   VmData.Vcent.VmInseucre,
		"GOVC_DATASTORE":  VmData.Vcent.VmDatastore,
		"GOVC_DATACENTER": VmData.Vcent.VmDatacenter,
	}
	return Goenv
}

