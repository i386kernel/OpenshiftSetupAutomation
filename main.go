package main

import (
	"flag"
	"fmt"
	"os"
)

func vmJustconvert(vmproperty VmProp, c chan string) {

	ignChange := changeIgn{
		filename:           vmproperty.IgnFile,
		intfname:           vmproperty.NetInterface,
		dns:                vmproperty.Dns,
		intfIP:             vmproperty.IpAddress,
		gateway:            vmproperty.DefaultGateway,
		contentNetworkName: vmproperty.IpInterfaceName,
	}
	var makeign string
	makeign = ignChange.justConvert()

	makevirtual := makeVM{
		RHOSTemplate:   VmData.Vclone.RhcosTemplate,
		datacenter:     VmData.Vcent.VmDatacenter,
		datastore:      VmData.Vcent.VmDatastore,
		vmName:         vmproperty.VMName,
		folder:         VmData.Vcent.VmFolder,
		memory:         VmData.Vclone.MemoryMB,
		cpu:            VmData.Vclone.CpuNumber,
		disksize:       VmData.Vclone.DiskSizeGB,
		mac:            vmproperty.MacAddress,
		ignitionbase64: makeign,
	}
	makevirtual.createVM()
	c <- "Done Creating VM's and applying relevant configs!!!!"
	close(c)
}

func prepareVM() string {
	c := make(chan string)
	for _, vms := range VmData.VmProperties {
		go vmJustconvert(vms, c)
	}
	fmt.Println(<-c)
	return "Done Provisioning VMs"
}

func destroy() {
	for _, vms := range VmData.VmProperties {
		dvm := makeVM{vmName: vms.VMName}
		fmt.Printf("DESTROYING !!!!: %s", vms.VMName)
		dvm.destroyVM()
	}
}

func start() {
	for _, vms := range VmData.VmProperties {
		pvm := makeVM{vmName: vms.VMName}
		fmt.Printf("Starting: %s", vms.VMName)
		pvm.powerONVM()
	}
}

func turnoff() {
	for _, vms := range VmData.VmProperties {
		pfvm := makeVM{vmName: vms.VMName}
		fmt.Printf("Powering Off: %s", vms.VMName)
		pfvm.powerOffVM()
	}
}

func reboot() {
	for _, vms := range VmData.VmProperties {
		prvm := makeVM{vmName: vms.VMName}
		fmt.Printf("Rebooting: %s", vms.VMName)
		prvm.rebootVM()
	}
}

func main() {
	govcenv := ConvertYamlToStruct()
	for key, value := range govcenv {
		err := os.Setenv(key, value)
		if err != nil {
			fmt.Println(err)
		}
	}
	makeFlag := flag.String("action", "poweron", "VM Actions (-Actions) could be, 'poweron, 'poweroff', 'deploy','destroy'(eg: -action=poweron)")
	flag.Parse()
	switch *makeFlag {
	case "poweron":
		start()
	case "poweroff":
		turnoff()
	case "reboot":
		reboot()
	case "deploy":
		prepareVM()
	case "destroy":
		destroy()
	default:
		start()
	}
}
