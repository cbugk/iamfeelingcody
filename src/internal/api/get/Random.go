package get

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cbugk/iamfeelingcody/src/internal/routine"
	"github.com/julienschmidt/httprouter"
)

func Random(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	n, err := strconv.Atoi(r.URL.Query().Get("n"))

	// Could not convert to integer
	if err != nil {
		log.Println(err.Error())
	} else if n <= 0 {
		m, err := json.Marshal(map[string]interface{}{
			"Error":     "Username size must be positive",
			"ErrorCode": 1,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(m)
		}
	} else { // Valid parameter
		// Run in parallel
		found := make(chan string, n)
		routine.Random(n, found)
		// Accumulate results
		result := []map[string]string{}
		for f := range found {
			log.Println("Append: ", f)
			result = append(result, map[string]string{
				"name": f,
			})
		}
		log.Println("closed chan")

		m, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{}"))
			log.Println(err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(m)
		}
	}
}
