package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	cfg "github.com/samsonannan/prizepicks-assessment/pkg/config"
	"github.com/samsonannan/prizepicks-assessment/pkg/db"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent"
	"github.com/samsonannan/prizepicks-assessment/pkg/handlers"
	"github.com/samsonannan/prizepicks-assessment/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	docs "github.com/samsonannan/prizepicks-assessment/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lib/pq"
)

func init() {
	// Make change
	// Load database configuration data using the 'LoadDbConfig' function from the 'cfg' package.
	config, err := cfg.LoadDbConfig("config")
	if err != nil {
		log.Fatalf("failed loading postgres configuration data: %v", err)
	}

	// Unmarshal the loaded configuration data (in JSON format) into the 'db.PostgresCfg' variable.
	err = json.Unmarshal(config, &db.PostgresCfg)
	if err != nil {
		log.Fatalf("failed unmarshaling postgres configuration map: %v", err)
	}

	// Create a new PostgreSQL client connection using the 'ent.Open' function.
	// The connection details are constructed using the values from the 'db.PostgresCfg' variable.
	pgClient, err := ent.Open("postgres", fmt.Sprintf("host=%s port=5432 user=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD")),
	)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool to create the required database schema resources.
	if err := pgClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Set the 'db.PostgresClient' variable to the newly created PostgreSQL client.
	db.PostgresClient = pgClient
}

// @title           Jurrasic Park API
// @version         1.0
// @description     This is a system to keep track of the different cages around the park and the different dinosaurs in each one
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	defer db.PostgresClient.Close()

	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	config.Level.SetLevel(zapcore.InfoLevel)

	err := logger.SetGlobalLogger(&config)
	if err != nil {
		log.Fatalf("failed to build logger: %v", err)
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		dinosaurs := v1.Group("/dinosaurs")
		{
			dinosaurs.GET("/", handlers.GetDinosaurs)
			dinosaurs.GET("/:id", handlers.GetDinosaurById)
			dinosaurs.PUT("/:id", handlers.EditDinosaur)
			dinosaurs.GET("/:id/cage", handlers.GetCageByDinosaurId)

			// TODO: Implement DELETE method for dinosaurs
			// dinosaurs.DELETE("/id", HelloWorld)
		}

		cages := v1.Group("/cages")
		{
			cages.GET("/", handlers.GetCages)
			cages.GET("/:id", handlers.GetCageById)
			cages.GET("/:id/dinosaurs", handlers.GetDinosaursByCageId)
			cages.POST("/", handlers.CreateCage)
			cages.POST("/:id/dinosaur", handlers.CageDinosaur)
			cages.PUT("/:id", handlers.EditCage)

			// TODO: Implement DELETE method for cages
			// cages.DELETE("/id", HelloWorld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
