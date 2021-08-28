package server

import (
	"encoding/json"
	"github.com/facundo-alarcon/url-shortener/urls/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func Routes(
	urlHandler *web.CreateUrlShortenerHandler,
	urlMongoHandler *web.UrlShortenerMongoHandler,
	) *chi.Mux {

	mux := chi.NewMux()

	// globals middleware (se van a usar en todas las rutas que yo defina en el multiplexor)
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs (si hay algun panic dentro de las rutas definidas el server no se va a apagar, sino que se va a recuperar)
	)

	// creacion de rutas
	// mux.method(path, handler)
	//mux.Post("/smartphones", sph.SaveSmartphoneHandler)
	mux.Post("/shorturl", urlHandler.SaveUrlHandler)
	mux.Post("/mongoshorturl",urlMongoHandler.AddUrlShortenerMongoHandler)
	mux.Get("/hello", helloHandler)
	//mux.Post("/reviews", reviewHandler.AddReviewHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// set sobreescribe si es que ya hay
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "facu")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}
