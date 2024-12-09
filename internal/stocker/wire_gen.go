// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package stocker

import (
	"github.com/google/wire"
	"h11/backend/internal/stocker/application/service"
	"h11/backend/internal/stocker/domain/repository"
	"h11/backend/internal/stocker/infrastructure/database"
	"h11/backend/internal/stocker/infrastructure/repository/implements"
	"h11/backend/internal/stocker/presentation/controller"
)

// Injectors from wire.go:

// コントローラーセット作成
func InitializeController() *ControllersSet {
	jancodeService := service.NewJancodeService()
	jancodeController := controller.NewJancodeController(jancodeService)
	db := database.GetMySQLConnection()
	itemRepositoryImpl := implements.NewItemRepositoryImpl(db)
	itemService := service.NewItemService(itemRepositoryImpl)
	storeRepositoryImpl := implements.NewStoreRepositoryImpl(db)
	storeAuthorizationService := service.NewStoreAuthorizationService(storeRepositoryImpl)
	itemController := controller.NewItemController(itemService, storeAuthorizationService)
	itemStockRepositoryImpl := implements.NewItemStockRepositoryImpl(db)
	itemStockService := service.NewItemStockService(itemStockRepositoryImpl)
	itemStockController := controller.NewItemStockController(itemStockService, storeAuthorizationService)
	userRepositoryImpl := implements.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepositoryImpl)
	authorizationService := service.NewAuthorizationService(userRepositoryImpl)
	userController := controller.NewUserController(userService, authorizationService)
	storeService := service.NewStoreService(userRepositoryImpl, storeRepositoryImpl)
	storeController := controller.NewStoreController(storeService)
	stockInRepositoryImpl := implements.NewStockInRepositoryImpl(db)
	stockInService := service.NewStockInService(stockInRepositoryImpl)
	stockInController := controller.NewStockInController(stockInService)
	controllersSet := &ControllersSet{
		JancodeController:   jancodeController,
		ItemController:      itemController,
		ItemStockController: itemStockController,
		UserController:      userController,
		StoreController:     storeController,
		StockInController:   stockInController,
	}
	return controllersSet
}

// wire.go:

// データベース
var databaseSet = wire.NewSet(database.GetMySQLConnection)

// リポジトリ
var repositorySet = wire.NewSet(implements.NewItemRepositoryImpl, wire.Bind(new(repository.ItemRepository), new(implements.ItemRepositoryImpl)), implements.NewItemStockRepositoryImpl, wire.Bind(new(repository.ItemStockRepository), new(implements.ItemStockRepositoryImpl)), implements.NewUserRepositoryImpl, wire.Bind(new(repository.UserRepository), new(implements.UserRepositoryImpl)), implements.NewStoreRepositoryImpl, wire.Bind(new(repository.StoreRepository), new(implements.StoreRepositoryImpl)), implements.NewStockInRepositoryImpl, wire.Bind(new(repository.StockInRepository), new(implements.StockInRepositoryImpl)))

// サービス
var serviceSet = wire.NewSet(service.NewJancodeService, service.NewItemService, service.NewItemStockService, service.NewUserService, service.NewAuthorizationService, service.NewStoreService, service.NewStoreAuthorizationService, service.NewStockInService)

// コントローラー
var controllerSet = wire.NewSet(controller.NewJancodeController, controller.NewItemController, controller.NewItemStockController, controller.NewUserController, controller.NewStoreController, controller.NewStockInController)

// コントローラーセット
type ControllersSet struct {
	JancodeController   controller.JancodeController
	ItemController      controller.ItemController
	ItemStockController controller.ItemStockController
	UserController      controller.UserController
	StoreController     controller.StoreController
	StockInController   controller.StockInController
}
