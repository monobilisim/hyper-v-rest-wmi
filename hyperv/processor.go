package hyperv

import (
	"hyper-v-rest-wmi/utilities"
	"net/http"

	"github.com/StackExchange/wmi"
	"github.com/gin-gonic/gin"
)

type Msvm_SummaryInformation struct {
	NumberOfProcessors   int16
	GuestOperatingSystem string
}

func queryProcessor(vmName string) ([]Msvm_SummaryInformation, error) {
	var dst []Msvm_SummaryInformation
	q := wmi.CreateQuery(&dst, "WHERE Name='"+vmName+"'")
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}

func Processor(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if !utilities.IsValidUUID(input) {
		c.Data(returnResponse("Invalid VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	result, err := queryProcessor(input)

	for _, a := range result {
		if string(a.NumberOfProcessors) == "" {
			c.Data(returnResponse("VM not found", http.StatusNotFound, "failure", "error"))
			return
		}
	}

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(result, http.StatusOK, "success", "Processor info is displayed in data field"))
}
