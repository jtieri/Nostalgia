package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	//fmt.Println("Inside GetIndex() .....")
	c.File("./views/index.html")
}

func GetStream(c *gin.Context) {
	//fmt.Println("Inside GetMedia() .....")
	//c.File("C:\\Users\\justi\\go\\src\\github.com\\jtieri\\Nostalgia\\media\\SLV.mkv")
	//c.File("C:\\Users\\justi\\go\\src\\github.com\\jtieri\\Nostalgia\\media\\now_playing\\prog_index.m3u8")
	//http.ServeFile(c.Writer, c.Request, "C:\\Users\\justi\\go\\src\\github.com\\jtieri\\Nostalgia\\media\\SLV.mkv")

	/*
		log.Println("Making connection to VLC stream...             ")
		resp, err := http.Get("http://127.0.0.1:8081/stream")
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Received response from VLC stream....          ")

		var extraHeaders = map[string]string {
			"Access-Control-Allow-Origin": "*",
		}
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, extraHeaders)

		// Print formatted response from VLC
		log.Println(resp)
		log.Printf("Status: %s \n", resp.Status)
		log.Printf("Status Code: %i \n", resp.StatusCode)
		log.Printf("Headers: %s \n", resp.Header)
		//log.Printf("Body: %s \n", respBody)
	*/

	/*
		remote, err := url.Parse("http://127.0.0.1:8081/stream")
		if err != nil {
			log.Fatal(err.Error())
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		//Define the director func
		//This is a good place to log, for example
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Header.Add("Access-Control-Allow-Origin", "*")
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Param("proxyPath")
		}

		// Print formatted response from VLC
		log.Println(c.Request)
		log.Printf("Host: %s \n", c.Request.Host)
		log.Printf("Headers: %s \n", c.Request.Header)
		log.Printf("Body: %s \n", c.Request.Body)

		proxy.ServeHTTP(c.Writer, c.Request)

	*/
}
