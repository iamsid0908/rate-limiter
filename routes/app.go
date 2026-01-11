package route

import (
	"my-echo-app/config" // Import your config package
	"my-echo-app/domain"
	"my-echo-app/handler"
	"my-echo-app/service"

	"github.com/redis/go-redis/v9" // Import go-redis
)

// 1. Add the field here so Router can see it
type AppModel struct {
	Health      handler.HealthHandler
	RedisClient *redis.Client
}

func App() AppModel {
	// 2. Initialize Redis using your Config package
	config.RedisInit()          // This connects to Redis
	rdb := config.RedisClient() // This grabs the active client

	healthDomain := &domain.HealthDomainCtx{}

	healthService := service.HealthService{
		HealthDomain: healthDomain,
	}

	healthHandler := handler.HealthHandler{
		HealthService: healthService,
	}

	return AppModel{
		Health:      healthHandler,
		RedisClient: rdb, // 3. Pass it to the struct
	}
}
