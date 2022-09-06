package err

import "github.com/kataras/iris/v12"

// 404
func StatusNotFound(c iris.Context) {
	c.WriteString("404")
}

// 500
func StatusInternalServerError(c iris.Context) {
	c.WriteString("500")
}
