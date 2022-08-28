package response

import (
	"log"
	"net/http"
	"testing"
)

func TestResponse(t *testing.T)  {
	var resp          *http.Response
	loc := resp.Header.Get("Location")
	log.Println(loc)
}



