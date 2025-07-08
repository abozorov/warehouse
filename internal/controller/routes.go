package controller

import (
	"warehouse/internal/configs"
	"warehouse/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServer() error {
	router := gin.Default()

	router.GET("/", Ping)

	// 1 аккаунт админа создаем заранее
	router.POST("/sign-in", SignIn)

	// только админ
	userG := router.Group("/users", CheckAdmin)
	{
		userG.POST("", CreateUser)
		userG.GET("/:id", GetUserByID)
		userG.GET("", GetAllUsers)
		userG.PATCH("/:id", UpdateUserByID)
	}

	// только админ
	cellG := router.Group("/cells", CheckAdmin)
	{
		cellG.POST("", CreateCell)
		cellG.GET("/:id", GetCellByID)
		cellG.GET("", GetAllCells)
		cellG.PATCH("/:id", UpdateCellByID)
	}

	// только админ
	statsG := router.Group("/stats", CheckAdmin)
	{
		// сортировка по дате, фильтрация по юзерам, 
		// вывод статистики: какие товары и сколько принял / отгрузил
		statsG.GET("/users/:id", UserStats)
		
		// сортировка по дате, фильтрация по контрагентам, 
		// вывод статистики: какие товары и сколько принял / отгрузил
		// statsG.GET("/couterparties/:id", CounterpartyStats)

		// сортировка по дате, фильтрация по товару, 
		// вывод статистики: какие товары и сколько принял / отгрузил
		statsG.GET("/products/:id", ProductStats)
		
		// сортировка по дате, фильтрация по типу, 
		// вывод статистики: какие товары и сколько принял / отгрузил
		statsG.GET("/batches", BatchStats)
	}

	// любой авторизованный пользователь
	productG := router.Group("/products", CheckAuthorization)
	{
		productG.POST("", CreateProduct)
		productG.GET("/:id", GetProductByID)
		productG.GET("", GetAllProducts)
		productG.PATCH("/:id", UpdateProductByID)
	}

	// любой авторизованный пользователь
	counterpartyG := router.Group("/counterparties", CheckAuthorization)
	{
		counterpartyG.POST("", CreateCounterparty)
		counterpartyG.GET("/:id", GetCounterpartyByID)
		counterpartyG.GET("", GetAllCounterparties)
		counterpartyG.PATCH("/:id", UpdateCounterpartyByID)
	}

	// любой авторизованный пользователь
	storageG := router.Group("/storages", CheckAuthorization)
	{
		storageG.GET("/:adressCode", GetStorageByAdressCode)
		storageG.GET("", GetAllStorages)
	}

	// любой авторизованный пользователь
	batchG := router.Group("/batches", CheckAuthorization)
	{
		batchG.POST("", CreateBatch)
		batchG.GET("/:id", GetBatchByID)
		batchG.GET("", GetAllBatches)
	}

	// роут для swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(configs.AppSettings.AppParams.PortRun); err != nil {
		logger.Error.Printf("[controller] RunServer():  Error during running HTTP server: %s", err.Error())
		return err
	}

	return nil
}
