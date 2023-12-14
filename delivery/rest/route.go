package rest

import (
	"clean-architecture/delivery/rest/entity"
	"clean-architecture/libs/httpresponse"
	mwApp "clean-architecture/libs/middleware/app"
	"clean-architecture/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Rest struct {
	goEnv  string
	mwApps mwApp.Middlewares
	entity entity.EntityHandler
}

// New ...
func New(
	svc service.Service,
	mwApps mwApp.Middlewares,
	goEnv string,
) *Rest {
	return &Rest{
		goEnv:  goEnv,
		mwApps: mwApps,
		entity: entity.NewEntityHandler(svc),
	}
}

func (re *Rest) Route(e *echo.Group) {
	re.mwApps.SetCompleteLogMiddleware(e)
	ping := e.Group("/ping")
	ping.Any("", re.Ping)

	tci := e.Group("/entity")
	tci.Use(re.mwApps.VerifyBasiAuth)
	tci.GET("/", re.entity.GetEntityHandler)
}

func (re *Rest) Ping(c echo.Context) error {
	// userID, ok := c.Get("user_id").(string)
	// err := errors.New("masokkkk")
	// if err != nil {
	// 	Error := errs.AssignErr(errs.AddTrace(err), errs.InvalidRequest)
	// 	httpresponse.ErrorTrace(c, http.StatusBadRequest, Error)
	// 	return nil
	// }
	httpresponse.Success(c, http.StatusOK, "clean-architecture:"+re.goEnv)
	return nil
}
