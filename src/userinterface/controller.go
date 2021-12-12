package userinterface

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Handler() gin.HandlerFunc
}
