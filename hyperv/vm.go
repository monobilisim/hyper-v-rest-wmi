package hyperv

import (
	"net/http"
	"time"

	"github.com/StackExchange/wmi"
	"github.com/gin-gonic/gin"
)

type MSVM_ComputerSystem struct {
	ElementName string
	InstallDate time.Time
	Name        string
	ProcessID   int32
}

func queryVMS() ([]MSVM_ComputerSystem, error) {
	var dst []MSVM_ComputerSystem
	q := wmi.CreateQuery(&dst, "WHERE Caption='Virtual Machine'")
	err := wmi.QueryNamespace(q, &dst, `root\virtualization\v2`)
	return dst, err
}

func VMS(c *gin.Context) {
	result, err := queryVMS()

	for _, a := range result {
		if string(a.Name) == "" {
			c.Data(returnResponse("VM not found", http.StatusNotFound, "failure", "error"))
			return
		}
	}

	if err != nil {
		c.Data(returnResponse(err.Error(), http.StatusInternalServerError, "failure", "error"))
		return
	}

	if len(result) == 0 {
		c.Data(returnResponse("No VM found", http.StatusNotFound, "failure", "error"))
		return
	}

	c.Data(returnResponse(result, http.StatusOK, "success", "VM info is displayed in data field"))
}
