package rest

import (
	"hyper-v-rest-wmi/hyperv"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func StartServer(port int, version string) {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	r := gin.Default()
	r.GET("/vms", hyperv.VMS)
	r.GET("/vms/:machid/summary", hyperv.Summary)
	r.GET("/vms/:machid/memory", hyperv.Memory)
	r.GET("/vms/:machid/vhd", hyperv.VHD)
	r.GET("/vms/:machid/ip", hyperv.Ip)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"Result":  "failure",
			"Message": "Wrong Path",
			"Data":    nil,
		})
	})

	r.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Result":  "success",
			"Message": "Application version",
			"Data":    version,
		})
	})

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	srv.ListenAndServe()

}
