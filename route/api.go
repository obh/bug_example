package route

import (
	_e "github.com/obh/bug_example/emitter"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// order handler endpoints
const (
	addRequest = "/request"
	getRequest = "/:request_id"
)

const (
	requestId = "requestId"
)

type APILogRoute struct {
	Emitter _e.Emitter
}

func ConfigureEventHandlerHTTP(e *echo.Group, emit _e.Emitter) {
	handler := &APILogRoute{Emitter: emit}
	handler.AddHandlers(e)
}

func (_this *APILogRoute) AddHandlers(e *echo.Group) {
	e.POST(addRequest, _this.AddRequest)
	//e.GET(getOrder, _this.GetOrder, metrics.HTTPMetric(getOrder), rateLimiter(getOrder))
}

func (_this *APILogRoute) AddRequest(c echo.Context) error {
	event := &_e.Event{MID: 1234, URI: "http://www.google.com", Request: "Sample request", Response: "sample response"}
	err := _this.Emitter.Emit(*event)
	if err != nil {
		log.Error("error in emitting object", err)
		return c.JSON(500, err.Error())
	}
	log.Info("Created event", event)
	return c.JSON(200, &event)
}
