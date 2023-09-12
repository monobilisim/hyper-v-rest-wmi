package hyperv

import (
	"hyper-v-rest-wmi/utilities"
	"net/http"

	"github.com/StackExchange/wmi"
	"github.com/gin-gonic/gin"
)

type Msvm_SyntheticEthernetPortSettingData struct {
	InstanceID string
}

type Msvm_GuestNetworkAdapterConfiguration struct {
	IPAddresses []string
}

func queryIp(vmName string) ([]Msvm_GuestNetworkAdapterConfiguration, error) {
	var dst_eth []Msvm_SyntheticEthernetPortSettingData
	q := "ASSOCIATORS OF {Msvm_VirtualSystemSettingData.InstanceID='Microsoft:" + vmName + "'} WHERE ResultClass = Msvm_SyntheticEthernetPortSettingData"
	wmi.QueryNamespace(q, &dst_eth, `root\virtualization\v2`)
	var dst []Msvm_GuestNetworkAdapterConfiguration
	for _, eth_data := range dst_eth {
		var tmp_dst []Msvm_GuestNetworkAdapterConfiguration
		q = "ASSOCIATORS OF {Msvm_SyntheticEthernetPortSettingData.InstanceID='" + eth_data.InstanceID + "'} WHERE ResultClass = Msvm_GuestNetworkAdapterConfiguration"
		utilities.Log.Info(q)
		err := wmi.QueryNamespace(q, &tmp_dst, `root\virtualization\v2`)
		utilities.Log.Info(tmp_dst)
		dst = append(dst, tmp_dst[0])
		utilities.Log.Info(dst)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func Ip(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if !utilities.IsValidUUID(input) {
		c.Data(returnResponse("Invalid VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	result, err := queryIp(input)

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(result, http.StatusOK, "success", "IP info is displayed in data field"))
}
