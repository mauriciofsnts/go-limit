package health

import (
	"net/http"

	"github.com/mauriciofsnts/gofast/pkg/server/api"
)

func Health(w http.ResponseWriter, r *http.Request) {
	ok := true

	api.Ok(w, ok)
}
