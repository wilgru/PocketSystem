package extensions

import (
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func Register(app *pocketbase.PocketBase) {
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		e.Router.GET("/api/healthz/go", func(re *core.RequestEvent) error {
			return re.JSON(http.StatusOK, map[string]string{
				"status": "ok",
				"source": "go",
			})
		})

		return e.Next()
	})
}
