package v1

import (
	"time"

	"github.com/calmato/gran-book/api/proto/book"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/gin-gonic/gin"
)

/**
 * ###############################################
 * handler
 * ###############################################
 */
type APIV1Handler interface {
	Routes(rg *gin.RouterGroup)
	RequiredAdminRotues(rg *gin.RouterGroup)
}

type apiV1Handler struct {
	Params
	now func() time.Time
}

type Params struct {
	Admin user.AdminServiceClient
	Auth  user.AuthServiceClient
	Book  book.BookServiceClient
	User  user.UserServiceClient
}

func NewAPIV1Handler(params *Params, now func() time.Time) APIV1Handler {
	return &apiV1Handler{
		Params: *params,
		now:    now,
	}
}

/**
 * ###############################################
 * routes
 * ###############################################
 */
func (h *apiV1Handler) Routes(rg *gin.RouterGroup) {
	apiV1 := rg.Group("/v1")
	h.v1AuthRoutes(apiV1.Group("/auth"))
	h.v1UserRoutes(apiV1.Group("/users"))
	h.v1AdminRoutes(apiV1.Group("/admin"))
	h.v1BookRoutes(apiV1.Group("/book"))
}

func (h *apiV1Handler) RequiredAdminRotues(rg *gin.RouterGroup) {
	apiV1 := rg.Group("/v1")
	h.v1AdminRoutesWithAdmin(apiV1.Group("/admin"))
}

// Auth Service
func (h *apiV1Handler) v1AuthRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.getAuth)
	rg.PATCH("/email", h.updateAuthEmail)
	rg.PATCH("/password", h.updateAuthPassword)
	rg.PATCH("/profile", h.updateAuthProfile)
	rg.POST("/thumbnail", h.uploadAuthThumbnail)
}

// User Service
func (h *apiV1Handler) v1UserRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.listUser)
	rg.GET("/:userID", h.getUser)
}

// Admin Service
func (h *apiV1Handler) v1AdminRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.listAdmin)
	rg.GET("/:userID", h.getAdmin)
}

// Admin Service (with admin role)
func (h *apiV1Handler) v1AdminRoutesWithAdmin(rg *gin.RouterGroup) {
	rg.POST("/:userID", h.createAdmin)
	rg.POST("/:userID/thumbnail", h.uploadAdminThumbnail)
	rg.DELETE("/:userID", h.deleteAdmin)
	rg.PATCH("/:userID/contact", h.updateAdminContact)
	rg.PATCH("/:userID/profile", h.updateAdminProfile)
	rg.PATCH("/:userID/password", h.updateAdminPassword)
}

// Book Service
func (h *apiV1Handler) v1BookRoutes(rg *gin.RouterGroup) {
	rg.DELETE("/:bookID", h.deleteBook)
}
