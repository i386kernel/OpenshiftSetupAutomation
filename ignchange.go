package openShiftSetup

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type changeIgn struct {
	filename           string
	intfname           string
	dns                string
	intfIP             string
	gateway            string
	contentNetworkName string
}

// Takes input of Filename InterfaceName, DNS, IP Address and Gateway
func (ci *changeIgn) ignitionchange() string {
	var igUnMarsh map[string]interface{}
	file, err := ioutil.ReadFile(ci.filename)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(file, &igUnMarsh)
	if err != nil {
		fmt.Println(err)
	}
	contentData := fmt.Sprintf("[Match]\nName=%s\n\n[Network]\nDNS=%s\nAddress=%s\nGateway=%s", ci.intfname, ci.dns, ci.intfIP, ci.gateway)
	igUnMarsh["networkd"] = map[string]map[string]string{"units": {"contents": contentData, "name": ci.contentNetworkName}}
	bs, err := json.MarshalIndent(igUnMarsh, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	base64str := base64.StdEncoding.EncodeToString(bs)
	return base64str
}

func (ci *changeIgn) bootstrapIgn() string {
	file, err := ioutil.ReadFile(ci.filename)
	if err != nil {
		fmt.Println(err)
	}
	base64str := base64.StdEncoding.EncodeToString(file)
	return base64str
}

func (ci *changeIgn) justConvert() string {
	file, err := ioutil.ReadFile(ci.filename)
	if err != nil {
		fmt.Println(err)
	}
	base64str := base64.StdEncoding.EncodeToString(file)
	return base64str
}
