package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getStudent(ctx *gin.Context) {
	student := server.svc.GetStudent("56654")
	ctx.JSON(http.StatusOK, student)
}
