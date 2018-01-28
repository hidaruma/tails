package mediaretrieval

import (
	"net/http"
	"context"
	"github.com/hidaruma/tails/subsonic"
	"github.com/hidaruma/tails/data"
)

type AvaterContainer struct {
	Avater []byte
}

func GetAvater(w http.ResponseWriter, r *http.Request) {

}