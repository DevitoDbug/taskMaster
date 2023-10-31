package utility

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getTaskID(r *http.Request) (ID int, err error) {
	param := mux.Vars(r)
	id, errVal := strconv.ParseInt(param["id"], 0, 0)
	if errVal != nil {
		return -1, errVal
	}
	return int(id), nil
}
