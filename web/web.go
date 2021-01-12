package web

import "trumpeter/httpd"

func IndexHandlerFunc(c *httpd.Context) {
	c.Write.Write([]byte("Hello ant-go"))
}
