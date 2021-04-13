//go:generate swag init -o ../docs -g ./api/api.go --dir .././

package api

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/middlewares/permission"
	"github.com/rl-os/api/pkg/transports/http"
)

var ProviderSet = wire.NewSet(
	CreateInitControllersFn,

	providerBeatmapSet,
	providerBeatmapsetSet,
	providerChatSet,
	providerFriendSet,
	providerMeSet,
	providerOAuthClientSet,
	providerOAuthTokenSet,
	providerUserSet,
)

// CreateInitControllersFn
//
// @title osu!lazer API
// @version 2.0
// @description This is a simple server.
// @host localhost:2400
// @BasePath /
//
// @contact.name RL GitHub
// @contact.url https://github.com/rl-os/api
//
// @securitydefinitions.oauth2.password OAuth2
// @in request-body
// @tokenUrl /oauth/token
// @scope.* Grants all access
// @scope.user.* Grants access as user (without system and admin api)
// @scope.admin.* Grants access as admin
// @scope.sys.* Grants access as system user (for example chatbot, worker and etc)
func CreateInitControllersFn(
	user *UserController,
	chat *ChatController,
	friend *FriendController,
	beatmap *BeatmapController,
	beatmapSet *BeatmapSetController,
	current *CurrentUserController,
	oauthToken *OAuthTokenController,
	oauthClient *OAuthClientController,
) http.InitControllers {
	return func(root *echo.Echo) {
		// TODO: move out to external oauth2 server
		root.POST("/api/v2/oauth/token", oauthToken.Create)
		root.POST("/api/v2/oauth/client", oauthClient.Create, permission.MustLogin)

		root.POST("/users", current.Create) // WARNING: LEGACY HANDLER
		root.POST("/api/v2/registration", current.Create)
		root.GET("/api/v2/me", current.Me, permission.MustLogin)
		root.GET("/api/v2/me/:mode", current.Me, permission.MustLogin)
		root.GET("/me/download-quota-check", echo.NotFoundHandler, permission.MustLogin)

		root.GET("/api/v2/users/:user", user.Get)
		root.GET("/api/v2/users/:user/", user.Get)
		root.GET("/api/v2/users/:user/:mode", user.Get)
		root.GET("/api/v2/users/:user/kudosu", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/users/:user/scores/:type", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/users/:user/beatmapsets/:type", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/users/:user/recent_activity", echo.MethodNotAllowedHandler)

		root.GET("/api/v2/friends", friend.GetAll)
		root.PUT("/api/v2/friends", friend.Add)
		root.DELETE("/api/v2/friends", friend.Remove)

		root.POST("/api/v2/chat/new", chat.NewPm, permission.MustLogin)
		root.GET("/api/v2/chat/updates", chat.Updates, permission.MustLogin)
		root.GET("/api/v2/chat/channels/:id/messages", chat.Messages, permission.MustLogin)
		root.POST("/api/v2/chat/channels/:id/messages", chat.Send, permission.MustLogin)
		root.GET("/api/v2/chat/channels", chat.GetAll, permission.MustLogin)
		root.GET("/api/v2/chat/channels/joined", chat.GetJoined, permission.MustLogin)
		root.PUT("/api/v2/chat/channels/:id/users/:user", chat.Join, permission.MustLogin)
		root.DELETE("/api/v2/chat/channels/:id/users/:user", chat.Leave, permission.MustLogin)

		root.GET("/api/v2/beatmapsets/:id", beatmapSet.Get)
		root.GET("/api/v2/beatmapsets/lookup", beatmapSet.Lookup)
		root.GET("/api/v2/beatmapsets/search", beatmapSet.Search)
		root.POST("/api/v2/beatmapsets/:id/favourites", beatmapSet.Favourite, permission.MustLogin)

		root.GET("/api/v2/beatmaps/lookup", beatmap.Lookup)
		root.GET("/api/v2/beatmaps/:id", beatmap.Get)
		root.GET("/api/v2/beatmaps/:id/scores", beatmap.Scores)

		root.GET("/api/v2/ping", echo.MethodNotAllowedHandler)

		root.GET("/api/v2/scores/:mode/:score/download", echo.MethodNotAllowedHandler)

		root.POST("/api/v2/rooms/", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/rooms/:room", echo.MethodNotAllowedHandler)
		root.PUT("/api/v2/rooms/:room/users/:user", echo.MethodNotAllowedHandler)
		root.DELETE("/api/v2/rooms/:room/users/:user", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/rooms/:room/leaderboard", echo.MethodNotAllowedHandler)
		root.POST("/api/v2/rooms/:room/playlist/:playlist/scores", echo.MethodNotAllowedHandler)
		root.PUT("/api/v2/rooms/:room/playlist/:playlist/scores/:score", echo.MethodNotAllowedHandler)

		root.GET("/api/v2/comments/", echo.MethodNotAllowedHandler)
		root.POST("/api/v2/comments/", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/comments/:comment", echo.MethodNotAllowedHandler)
		root.PUT("/api/v2/comments/:comment", echo.MethodNotAllowedHandler)
		root.PATCH("/api/v2/comments/:comment", echo.MethodNotAllowedHandler)
		root.DELETE("/api/v2/comments/:comment", echo.MethodNotAllowedHandler)
		root.POST("/api/v2/comments/:comment/vote", echo.MethodNotAllowedHandler)
		root.DELETE("/api/v2/comments/:comment/vote", echo.MethodNotAllowedHandler)

		root.GET("/api/v2/Notifications/", echo.MethodNotAllowedHandler)
		root.POST("/api/v2/Notifications/mark-read", echo.MethodNotAllowedHandler)

		// === Misc ===
		root.POST("/api/v2/reports", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/changelog", echo.MethodNotAllowedHandler)
		root.GET("/api/v2/changelog/:changelog", echo.MethodNotAllowedHandler)
	}
}
