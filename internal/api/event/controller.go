package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"test/internal/event"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type Controller struct {
	dbClient  *redis.Client
	apiClient *echo.Echo
}

func NewController(
	dbClient *redis.Client,
	apiClient *echo.Echo,
) *Controller {
	return &Controller{
		dbClient:  dbClient,
		apiClient: apiClient,
	}
}

func (c *Controller) Create(ctx echo.Context) error {
	var request event.CreateRequest

	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(request.Invitees) > 100 {
		return echo.NewHTTPError(400, "event can have up to 100 invitees")
	}

	id := uuid.NewString()
	entity := event.Entity{
		ID:            id,
		CreateRequest: request,
	}

	data, err := json.Marshal(entity)
	if err != nil {
		return echo.NewHTTPError(500, "failed to marshal event")
	}

	if err = c.dbClient.SetNX(
		ctx.Request().Context(),
		fmt.Sprintf("event#%s", data),
		data,
		0,
	).Err(); err != nil {
		return echo.NewHTTPError(500, "failed to create event")
	}

	return ctx.JSON(200, &entity)
}
