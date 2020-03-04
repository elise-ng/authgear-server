package main

import (
	"net/http"
	"os"

	"github.com/davidbyttow/govips/pkg/vips"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/skygeario/skygear-server/pkg/asset"
	"github.com/skygeario/skygear-server/pkg/asset/config"
	"github.com/skygeario/skygear-server/pkg/asset/handler"
	"github.com/skygeario/skygear-server/pkg/core/cloudstorage"
	coreConfig "github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/db"
	"github.com/skygeario/skygear-server/pkg/core/logging"
	"github.com/skygeario/skygear-server/pkg/core/middleware"
	"github.com/skygeario/skygear-server/pkg/core/redis"
	"github.com/skygeario/skygear-server/pkg/core/sentry"
	"github.com/skygeario/skygear-server/pkg/core/server"
	"github.com/skygeario/skygear-server/pkg/core/validation"
)

/*
	@API Asset Gear
	@Version 1.0.0
	@Server {base_url}/_asset
		Asset Gear URL
		@Variable base_url https://my_app.skygearapis.com
			Skygear App URL

	@SecuritySchemeAPIKey access_key header X-Skygear-API-Key
		Access key used by client app
	@SecuritySchemeAPIKey master_key header X-Skygear-API-Key
		Master key used by admins, can perform administrative operations.
		Can be used as access key as well.
	@SecuritySchemeHTTP access_token Bearer token
		Access token of user
*/
func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	logging.SetModule("asset")
	loggerFactory := logging.NewFactory(
		logging.NewDefaultLogHook(nil),
		&sentry.LogHook{Hub: sentry.DefaultClient.Hub},
	)
	logger := loggerFactory.NewLogger("asset")

	if err := godotenv.Load(); err != nil {
		logger.WithError(err).Debug("Cannot load .env file")
	}

	configuration := config.Configuration{}
	envconfig.Process("", &configuration)
	if err := configuration.Initialize(); err != nil {
		logger.WithError(err).Panic("cannot initialize configuration")
	}

	var storage cloudstorage.Storage
	switch configuration.Storage.Backend {
	case config.StorageBackendAzure:
		storage = cloudstorage.NewAzureStorage(
			configuration.Storage.Azure.ServiceURL,
			configuration.Storage.Azure.StorageAccount,
			configuration.Storage.Azure.AccessKey,
			configuration.Storage.Azure.Container,
		)
	case config.StorageBackendGCS:
		storage = cloudstorage.NewGCSStorage(
			configuration.Storage.GCS.CredentialsJSON,
			configuration.Storage.GCS.ServiceAccount,
			configuration.Storage.GCS.Bucket,
		)
	case config.StorageBackendS3:
		storage = cloudstorage.NewS3Storage(
			configuration.Storage.S3.AccessKey,
			configuration.Storage.S3.SecretKey,
			configuration.Storage.S3.Endpoint,
			configuration.Storage.S3.Region,
			configuration.Storage.S3.Bucket,
		)
	}

	dbPool := db.NewPool()
	redisPool, err := redis.NewPool(configuration.Redis)
	if err != nil {
		logger.Fatalf("fail to create redis pool: %v", err)
	}

	validator := validation.NewValidator("http://v2.skgyear.io")
	validator.AddSchemaFragments(
		handler.PresignUploadRequestSchema,
		handler.SignRequestSchema,
	)

	dependencyMap := &asset.DependencyMap{
		UseInsecureCookie: configuration.UseInsecureCookie,
		Storage:           storage,
		Validator:         validator,
	}

	serverOption := server.DefaultOption()
	serverOption.GearPathPrefix = "/_asset"
	var rootRouter *mux.Router
	var appRouter *mux.Router
	if configuration.Standalone {
		filename := configuration.StandaloneTenantConfigurationFile
		reader, err := os.Open(filename)
		if err != nil {
			logger.WithError(err).Error("Cannot open standalone config")
		}
		tenantConfig, err := coreConfig.NewTenantConfigurationFromYAML(reader)
		if err != nil {
			logger.WithError(err).Fatal("Cannot parse standalone config")
		}

		rootRouter, appRouter = server.NewRouterWithOption(serverOption)
		appRouter.Use(middleware.WriteTenantConfigMiddleware{
			ConfigurationProvider: middleware.ConfigurationProviderFunc(func(_ *http.Request) (coreConfig.TenantConfiguration, error) {
				return *tenantConfig, nil
			}),
		}.Handle)
		appRouter.Use(middleware.RequestIDMiddleware{}.Handle)
		appRouter.Use(middleware.CORSMiddleware{}.Handle)
	} else {
		rootRouter, appRouter = server.NewRouterWithOption(serverOption)
		appRouter.Use(middleware.ReadTenantConfigMiddleware{}.Handle)
	}

	appRouter.Use(middleware.DBMiddleware{Pool: dbPool}.Handle)
	appRouter.Use(middleware.RedisMiddleware{Pool: redisPool}.Handle)
	appRouter.Use(middleware.AuthMiddleware{}.Handle)
	appRouter.Use(middleware.Injecter{
		MiddlewareFactory: middleware.AuthnMiddlewareFactory{},
		Dependency:        dependencyMap,
	}.Handle)

	handler.AttachPresignUploadHandler(appRouter, dependencyMap)
	handler.AttachSignHandler(appRouter, dependencyMap)
	handler.AttachGetHandler(appRouter, dependencyMap)
	handler.AttachListHandler(appRouter, dependencyMap)
	handler.AttachDeleteHandler(appRouter, dependencyMap)
	handler.AttachUploadFormHandler(appRouter, dependencyMap)
	handler.AttachPresignUploadFormHandler(appRouter, dependencyMap)

	srv := &http.Server{
		Addr:    configuration.ServerHost,
		Handler: rootRouter,
	}
	server.ListenAndServe(srv, logger, "Starting asset gear")
}
