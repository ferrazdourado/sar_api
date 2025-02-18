package main

import (
    "log"
    "path/filepath"
    "context"
    "github.com/ferrazdourado/sar_api/pkg/config"
    "github.com/ferrazdourado/sar_api/internal/routes"
    "github.com/ferrazdourado/sar_api/internal/controllers"
    "github.com/ferrazdourado/sar_api/internal/services"
    "github.com/ferrazdourado/sar_api/internal/repository/mongodb"
)

func main() {
    // Carregar configurações
    configPath := filepath.Join("config", "config.yaml")
    cfg, err := config.LoadConfig(configPath)
    if err != nil {
        log.Fatalf("Erro ao carregar configurações: %v", err)
    }

    // Inicializar conexão MongoDB
    db := mongodb.NewMongoDB(cfg.Database.URI, cfg.Database.Database)
    if err := db.Connect(context.Background()); err != nil {
        log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
    }
    defer db.Disconnect(context.Background())

    // Inicializar repositórios
    userRepo := mongodb.NewUserRepository(db)
    vpnRepo := mongodb.NewVPNRepository(db)

    // Inicializar serviços
    authService := services.NewAuthService(userRepo, cfg)
    vpnService := services.NewVPNService(vpnRepo)

    // Inicializar controllers
    authController := controllers.NewAuthController(authService)
    vpnController := controllers.NewVPNController(vpnService)

    // Configurar router
    router := routes.NewRouter(vpnController, authController, cfg)
    r := router.SetupRoutes()

    // Iniciar servidor
    log.Printf("Servidor iniciando na porta %s", cfg.Server.Port)
    if err := r.Run(":" + cfg.Server.Port); err != nil {
        log.Fatal("Erro ao iniciar servidor:", err)
    }
}