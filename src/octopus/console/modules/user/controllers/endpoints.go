package users

import (
	"encoding/json"
	"octopus/console/models"
	"octopus/console/modules/user/aaa"

	"gopkg.in/labstack/echo.v3"
)

func login(ctx echo.Context) error {
	p := &aaa.User{}
	dec := json.NewDecoder(ctx.Request().Body)
	err := dec.Decode(p)
	if err != nil {
		return ctx.JSON(400, struct {
			error error `json:"error"`
		}{
			error: err,
		})
	}

	if valid := models.NewOctManager().GetUser(p.Username, p.Password); !valid {
		return ctx.JSON(401, struct {
			error error `json:"error"`
		}{
			error: err,
		})
	}

	//TODO should redirect to panel
	return nil
}
