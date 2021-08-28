package web

import (
	"encoding/json"
	"github.com/facundo-alarcon/url-shortener/internal/database"
	"github.com/facundo-alarcon/url-shortener/urls/gateway/mysql"
	url "github.com/facundo-alarcon/url-shortener/urls/models"
	"net/http"
)

type CreateUrlShortenerHandler struct {
	mysql.UrlCreateGateway
}

func NewCreateUrlShortenerHandler (client *database.MySqlClient) *CreateUrlShortenerHandler{
	return &CreateUrlShortenerHandler{mysql.NewUrlCreateGateway(client)}
}


//func NewReviewHandler(mongo *database.Mongo) *ReviewHandler {
//	return &ReviewHandler{gateway.NewReviewGateway(mongo)}
//}

func (u *CreateUrlShortenerHandler) SaveUrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// parseamos el body del json que mandamos
	cmd := parseRequest(r)
	// creamos una entrada en la tabla
	res, err :=u.Create(cmd)

	// ejempo sin el Create, podemos hacer un json aca hardcodeado
	//res, err := u.Create(&url.CreateURLCMD{
	//	Original: "https://google.com",
	//	Short: "https://localhos/g1",
	//})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg":"error in create short url"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

func parseRequest(r *http.Request) *url.CreateURLCMD{
	body := r.Body
	// con el defer esperamos a que termine la ejecucion de toda una funcion para que se llame
	defer body.Close()
	var cmd url.CreateURLCMD
	// decodificamos el cuerpo del body y no asignamos el error en este caso
	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}


// todo generateshorturl string
// todo get shorturl