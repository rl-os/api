package friends

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/entity"
	userService "github.com/deissh/osu-lazer/api/pkg/services/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// PutFriendRequestData contain incoming data
type PutFriendRequestData struct {
	TargetId uint `json:"target_id" query:"target_id"`
}

// Put user friends handler
func Put(c echo.Context) (err error) {
	data := c.Get("current_user")

	current, ok := data.(*entity.User)
	if !ok {
		return pkg.NewHTTPError(401, "auth_token_required", "Need auth token in headers")
	}

	params := new(PutFriendRequestData)
	if err := c.Bind(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_params_error", "Failed binding params")
	}

	if err := validator.New().Struct(params); err != nil {
		return pkg.NewHTTPError(http.StatusBadRequest, "request_validate_error", "Failed validate")
	}

	err = userService.SetFriend(current.ID, params.TargetId)
	users, err := userService.GetSubscriptions(current.ID)
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}
