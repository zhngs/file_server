package main

import (
	"log"
	"net"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, "./upload/"+filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "\r\nfile %s uploaded successfully\r\n\r\ndownload file: curl -O http://%s:6060/upload/%s\r\n", file.Filename, GetLocalIP(), file.Filename)
	})
	r.StaticFS("/", gin.Dir(".", true))

	log.Printf("file server start %s:%d", GetLocalIP(), 6060)
	r.Run(":6060")
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
