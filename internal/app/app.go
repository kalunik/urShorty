package app

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kalunik/urShorty/config"
	"github.com/kalunik/urShorty/internal/api"
	r "github.com/kalunik/urShorty/internal/repository"
	"github.com/kalunik/urShorty/internal/usecase"
	"github.com/kalunik/urShorty/pkg/logger"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type App struct {
	r                *api.Router
	redisClient      *redis.Client
	clickhouseClient driver.Conn
	log              logger.Logger
	conf             config.AppConfig
}

func NewApp(redis *redis.Client, clickhouse driver.Conn, logger logger.Logger, config *config.AppConfig) *App {
	return &App{r: nil, redisClient: redis, clickhouseClient: clickhouse, log: logger, conf: *config}
}

func (a *App) Run() {
	redisRepo := r.NewRedisRepository(a.redisClient)
	clickhouseRepo := r.NewClickhouseRepository(a.clickhouseClient)

	urlService := usecase.NewPathMetaUsecase(redisRepo, clickhouseRepo, a.log)

	urlPairHandlers := api.NewPathMetaHandlers(urlService, a.log)

	a.r = api.NewRouter()
	//middleware
	a.r.Mux.Use(middleware.RequestID)

	//check if I need return for UrlPairRouter(),
	//I expect mux* (pointer) do all stuff
	a.r.PathMetaRoutes(urlPairHandlers)

	a.log.Infof("server will start on %s port", a.conf.Server.Port)
	http.ListenAndServe(a.conf.Server.Port, a.r.Mux)
	//shutdown
}
