// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package routers

import (
	"melon-service/controllers"
	"melon-service/databases"
	"melon-service/interfaces"
	"melon-service/repositories"
	"melon-service/services"

	"github.com/google/wire"
)

// InitializeUserController 声明injector的函数签名
func InitializeUserController(db *databases.DB) controllers.UserController {
	wire.Build(userServiceRepoSet, userRepositorySet)
	return controllers.UserController{} //返回值没有实际意义，只需符合函数签名即可
}

var userServiceRepoSet = wire.NewSet(controllers.NewUserController, wire.Bind(new(interfaces.IUserService), new(services.UserService)), services.NewUserService)

var userRepositorySet = wire.NewSet(wire.Bind(new(interfaces.IUserRepository), new(repositories.UserRepository)), repositories.NewUserRepository)
