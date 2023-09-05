package wmi

import (
	"github.com/StackExchange/wmi"
)

type Msvm_SyntheticEthernetPortSettingData struct {
	InstanceID string
}

type Msvm_GuestNetworkAdapterConfiguration struct {
	IPAddresses []string
}

func Ip(vmName string) ([]Msvm_GuestNetworkAdapterConfiguration, error) {
	var dst_eth []Msvm_SyntheticEthernetPortSettingData
	q := "ASSOCIATORS OF {Msvm_VirtualSystemSettingData.InstanceID='Microsoft:" + vmName + "'} WHERE ResultClass = Msvm_SyntheticEthernetPortSettingData"
	wmi.QueryNamespace(q, &dst_eth, `root\virtualization\v2`)
	var dst []Msvm_GuestNetworkAdapterConfiguration
	q = "ASSOCIATORS OF {Msvm_SyntheticEthernetPortSettingData.InstanceID='" + dst_eth[0].InstanceID + "'} WHERE ResultClass = Msvm_GuestNetworkAdapterConfiguration"
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}
