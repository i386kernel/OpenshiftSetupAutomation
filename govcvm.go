package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type makeVM struct {
	RHOSTemplate   string
	datacenter     string
	datastore      string
	vmName         string
	folder         string
	memory         int
	cpu            int
	disksize       int
	mac            string
	ignitionbase64 string
}

func (mv *makeVM) createVM() {
	cloneVM := fmt.Sprintf(`govc vm.clone -vm %s -annotation=%s -c=%d -m=%d -net="VM Network" -net.address="%s" -on=false -folder=%s -ds=%s %s`, mv.RHOSTemplate, mv.vmName, mv.cpu, mv.memory, mv.mac, mv.folder, mv.datastore, mv.vmName)
	cmd := exec.Command("bash", "-c", cloneVM)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal("Start returned: %v\n", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal("Wait returned: %v\n", err)
	}
	enableUUID := fmt.Sprintf(`govc vm.change -e="disk.enableUUID=1" -vm="/%s/vm/%s/%s"`, mv.datacenter,mv.folder, mv.vmName)
	guestigninfochange := fmt.Sprintf(`govc vm.change -e="guestinfo.ignition.config.data=changeme" -vm="/%s/vm/%s/%s"`, mv.datacenter, mv.folder, mv.vmName)
	guestignencoding := fmt.Sprintf(`govc vm.change -e="guestinfo.ignition.config.data.encoding=base64" -vm="/%s/vm/%s/%s"`, mv.datacenter, mv.folder, mv.vmName)
	vmchange := fmt.Sprintf(`govc vm.change -e="guestinfo.ignition.config.data=%s" -vm=%s`, mv.ignitionbase64, mv.vmName)
	diskchange := fmt.Sprintf(`govc vm.disk.change -vm %s  -size %dG`, mv.vmName, mv.disksize)

	for _, v := range []string{enableUUID, guestigninfochange, guestignencoding, vmchange, diskchange} {
		cmd = exec.Command("bash", "-c", v)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (mv *makeVM) powerONVM() {
	cmd := exec.Command("bash", "-c", "govc", "vm.power", "-on", mv.vmName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func (mv *makeVM) powerOffVM() {
	cmd := exec.Command("bash", "-c", "govc", "vm.power", "-off", "-force", mv.vmName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func (mv *makeVM) rebootVM() {
	cmd := exec.Command("bash", "-c", "govc", "vm.power", "-r", mv.vmName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func (mv *makeVM) destroyVM() {
	destroyvm := fmt.Sprintf(`govc vm.destroy %s`, mv.vmName)
	fmt.Println(destroyvm)
	cmd := exec.Command("/bin/bash", "-c", destroyvm)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Start returned: %v\n", err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}
}
