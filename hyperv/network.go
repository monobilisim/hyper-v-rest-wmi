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

func queryNetwork(vmName string) ([]Msvm_GuestNetworkAdapterConfiguration, error) {
	var dst_eth []Msvm_SyntheticEthernetPortSettingData
	q := "ASSOCIATORS OF {Msvm_VirtualSystemSettingData.InstanceID='Microsoft:" + vmName + "'} WHERE ResultClass = Msvm_SyntheticEthernetPortSettingData"
	wmi.QueryNamespace(q, &dst_eth, `root\virtualization\v2`)
	var dst []Msvm_GuestNetworkAdapterConfiguration
	q = "ASSOCIATORS OF {Msvm_SyntheticEthernetPortSettingData.InstanceID='" + dst_eth[0].InstanceID + "'} WHERE ResultClass = Msvm_GuestNetworkAdapterConfiguration"
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}

func Network(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if !utilities.IsValidUUID(input) {
		c.Data(returnResponse("Invalid VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	result, err := queryNetwork(input)

	for _, a := range result {
		for _, b := range a.IPAddresses {
			if string(b) == "" {
				c.Data(returnResponse("VM not found", http.StatusNotFound, "failure", "error"))
				return
			}
		}
	}

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(result, http.StatusOK, "success", "Network info is displayed in data field"))
}
