package main

import (
    "log"
    "github.com/ferrazdourado/sar_api/internal/routes"
    "github.com/ferrazdourado/sar_api/internal/controllers"
    "github.com/ferrazdourado/sar_api/internal/services"
    "github.com/ferrazdourado/sar_api/internal/repository/mongodb"
)

func main() {
    // Inicializar repositórios
    userRepo := mongodb.NewUserRepository()
    vpnRepo := mongodb.NewVPNRepository()

    // Inicializar serviços
    authService := services.NewAuthService(userRepo)
    vpnService := services.NewVPNService(vpnRepo)

    // Inicializar controllers
    authController := controllers.NewAuthController(authService)
    vpnController := controllers.NewVPNController(vpnService)

    // Configurar router
    router := routes.NewRouter(vpnController, authController)
    r := router.SetupRoutes()

    // Iniciar servidor
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Erro ao iniciar servidor:", err)
    }
}