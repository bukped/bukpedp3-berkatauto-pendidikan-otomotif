package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	berkatbepkg "github.com/berkatauto/berkat-bepkg"
)

func init() {
	functions.HTTP("HelloPostArticle", HelloPostArticle)
}

func HelloPostArticle(w http.ResponseWriter, r *http.Request) {
	// Set header Access-Control-Allow-Origin untuk mengizinkan permintaan dari domain yang spesifik.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set header Access-Control-Allow-Methods untuk mengizinkan metode HTTP yang diizinkan.
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Set header Access-Control-Allow-Headers untuk mengizinkan header yang diizinkan dalam permintaan.
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		// Jika permintaan adalah preflight OPTIONS, Anda hanya perlu mengirimkan header CORS.
		return
	}

	// Tulis respons Anda ke Writer seperti yang Anda lakukan sebelumnya.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := berkatbepkg.GCFPostArticle("MONGOSTRING", "berkatauto", "articleSet", r)
	fmt.Fprintf(w, response)
}
