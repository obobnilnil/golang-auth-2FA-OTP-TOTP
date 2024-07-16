package servers

import (
	"auth_git/auth/handlers"
	"auth_git/auth/repositories"
	"auth_git/auth/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutesAuth(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositoryAdapter(db)
	s := services.NewServiceAdapter(r)
	h := handlers.NewHanerhandlerAdapter(s)

	router.POST("/api/sendOTPEmail", h.RequestEmailForValidateOTPChicCRMHandlers)
	router.POST("/api/validateOTPEmail", h.ValidateOTPFromRequestEmailChicCRMHandlers)
	router.POST("/api/qrTOTP", h.QrTOTPChicCRMHandlers)
	router.POST("/api/validateQrTOTP", h.ValidateQrTOTPChicCRMHandlers)
	router.DELETE("/api/deleteKeyQrTOTP", h.DeleteKeyQrTOTPChicCRMHandlers)
	router.POST("/api/sendOTPLine", h.RequestEmailForValidateOTPLineChicCRMHandlers)
}
