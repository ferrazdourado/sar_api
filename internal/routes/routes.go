package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/ferrazdourado/sar_api/internal/controllers"
    "github.com/ferrazdourado/sar_api/internal/middleware"
    "github.com/ferrazdourado/sar_api/pkg/config"
)

type Router struct {
    vpnController  *controllers.VPNController
    authController *controllers.AuthController
    config         *config.Config
}

func NewRouter(
    vpnController *controllers.VPNController,
    authController *controllers.AuthController,
    cfg *config.Config,
) *Router {
    return &Router{
        vpnController:  vpnController,
        authController: authController,
        config:        cfg,
    }
}

func (r *Router) SetupRoutes() *gin.Engine {
    router := gin.Default()
    authMiddleware := middleware.NewAuthMiddleware(r.config)

    // Grupo de rotas públicas
    public := router.Group("/api/v1")
    {
        auth := public.Group("/auth")
        {
            auth.POST("/login", r.authController.Login)
            auth.POST("/register", r.authController.Register)
        }
    }

    // Grupo de rotas protegidas
    protected := router.Group("/api/v1")
    protected.Use(authMiddleware.Handler())
    {
        vpn := protected.Group("/vpn")
        {
            vpn.GET("/config", r.vpnController.GetVPNConfig)
            vpn.POST("/config", r.vpnController.CreateVPNConfig)
            vpn.GET("/status", r.vpnController.GetVPNStatus)
        }
    }

    return router
}