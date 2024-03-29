package service

import (
	"fmt"
	"go-basic/pkg/database/mysql"
	"go-basic/pkg/logger"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	databaseInstance *gorm.DB
	loggerInstance   *logger.Logger
)

func InitServices() {
	loggerInstance = logger.GetLogger("handler")
	databaseInstance = mysql.GetDatabase()
}

type ErrorOutput struct {
	Message    string `json:message`
	StatusCode string `json:statusCode`
}

func outputErrorMessageFormat(name, dataType string) error {
	return fmt.Errorf("param : %s (type: %s) is required", name, dataType)
}

func outputErrorFormat(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}

func outputSuccessFormat(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", operation),
		"data":    data,
	})
}
