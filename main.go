package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//type MyRender struct {
//}
//
//func (m *MyRender) Render(w http.ResponseWriter) error {
//
//}
//func (m *MyRender) WriteContentType(w http.ResponseWriter) {
//
//}
func main() {
	//r := gin.Default()
	//r.Use(func(c *gin.Context) {
	//	fmt.Println("before")
	//	c.Next()
	//	fmt.Println("after")
	//	c.Writer.WriteString("  append")
	//	//fmt.Println(c.Writer.Size())
	//	//fmt.Printf("%#v \n", c.Writer.ResponseWriter)
	//	//conn, w, _ := c.Writer.Hijack()
	//	c.Writer
	//	//buf := make([]byte, 100)
	//	fmt.Println("from hijack--")
	//	//n, _ := w.Read(buf)
	//	//n, _ := conn.Read(buf)
	//	//fmt.Printf("from hijack %s \n", string(buf[:n]))
	//
	//})
	//r.GET("/", func(c *gin.Context) {
	//	fmt.Println("Get ")
	//	c.JSON(200, "abc")
	//})
	//r.Run(":8001")

	r := mgin.Default()
	r.Use(func(c *mgin.Context) {
		fmt.Println("before ")
		c.Next()
		fmt.Println("after ")
	})
	r.GET("/", func(c *mgin.Context) {
		fmt.Println("in get")
		c.JSON(200, "hello")
	})
	r.Run(":8001")
}
