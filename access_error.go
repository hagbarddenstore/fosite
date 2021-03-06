package fosite

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Fosite) WriteAccessError(rw http.ResponseWriter, requester AccessRequester, err error) {
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")

	rfcerr := ErrorToRFC6749Error(err)
	js, err := json.Marshal(rfcerr)
	if err != nil {
		http.Error(rw, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	if rfcerr.Name == errServerErrorName {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.Write(js)
}
