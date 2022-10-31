package app

import (
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, v interface{}) (bool, error) {
	err := c.ShouldBind(v)
	if err != nil {
		return false, err
	}

	return true, nil
}
