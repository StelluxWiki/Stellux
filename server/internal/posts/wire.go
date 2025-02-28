//go:build wireinject
// +build wireinject

package posts

import (
	"server/internal/posts/internal/api"
	"server/internal/posts/internal/repo"
	"server/internal/posts/internal/repo/dao"
	"server/internal/posts/internal/service"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var postsProvider = wire.NewSet(
	api.NewPostHandler,
	service.NewPostsService,
	repo.NewPostsRepo,
	dao.NewPostsDao,
	wire.Bind(new(service.IPostsService), new(*service.PostsService)),
	wire.Bind(new(repo.IPostsRepo), new(*repo.PostsRepo)),
	wire.Bind(new(dao.IPostsDao), new(*dao.PostsDao)),
)

func InitPostsModule(database *mongo.Database) *Module {
	wire.Build(
		postsProvider,
		wire.Struct(new(Module), "Hdl", "Svc"),
	)
	return nil
}
