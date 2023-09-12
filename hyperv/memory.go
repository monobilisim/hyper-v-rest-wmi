package hyperv

import (
	"hyper-v-rest-wmi/utilities"
	"net/http"

	"github.com/StackExchange/wmi"
	"github.com/gin-gonic/gin"
)

type Msvm_MemorySettingData struct {
	VirtualQuantity int64
}

func queryMemory(vmName string) ([]Msvm_MemorySettingData, error) {
	var dst []Msvm_MemorySettingData
	q := "ASSOCIATORS OF {Msvm_VirtualSystemSettingData.InstanceID='Microsoft:" + vmName + "'} WHERE ResultClass = Msvm_MemorySettingData"
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}

func Memory(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if !utilities.IsValidUUID(input) {
		c.Data(returnResponse("Invalid VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	result, err := queryMemory(input)

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(result, http.StatusOK, "success", "Memory info is displayed in data field"))
}
