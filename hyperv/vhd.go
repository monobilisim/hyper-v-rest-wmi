package hyperv

import (
	"hyper-v-rest-wmi/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryVHD(vmName string) ([]byte, error) {
	ps := `Get-VHD -Id ` + vmName + ` | ConvertTo-Json`
	output, err := utilities.CommandLine(ps)
	return output, err
}

func VHD(c *gin.Context) {
	input := c.Param("machid")

	if input == "" {
		c.Data(returnResponse("No VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	if !utilities.IsValidUUID(input) {
		c.Data(returnResponse("Invalid VM ID specified", http.StatusBadRequest, "failure", "error"))
		return
	}

	result, err := queryVHD(input)

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	c.Data(returnResponse(result, http.StatusOK, "success", "VHD info is displayed in data field"))
}
