package server

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

type MyServer struct {
	server *http.Server
}
// new server devolveria un nuevo tipo baso en http.Server
// devuelve un puntero porque dentro de la funcion crea un struct MyServer
// de otra forma no se podria devolver esa estructura crada, necesitas la dir de memoria de la misma

// un caso similar de cuando queres modificar una variable pasada por parametro dentro de una funcion
// cuando se pasa un puntero como parametro es porque
// dentro se va a modificar justamente eso que se pasa como parametro
func NewServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":3000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &MyServer{s}
}

func (s *MyServer) Run() {
	// log Fatal lo que hace es que listen and search devuelva un error
	// ListenAndServe corre detras de una goroutine, lo que hace es, a penas escapa de esa rutina
	// sigue con la linea siguiente
	// si escribimos s.server.ListenAndServe() por fuera de Log.Fatal
	// lo que va a pasar es correr y al mismo tiempo terminar (o en un correr de tiempo muy cercano)
	// entonces no vamos a poder tener el servidor vivo mucho tiempo (solo estaria activo unos milisegundos)
	// entonces no podriamos usarlo ni probar nada
	// Lo que hace entonces log.Fatal es esperar a que ListenAndServe() devuelva un error, entonces
	// hasta que eso pase la linea no se va a terminar de ejecutar toda
	// entonces lo que hace es ejecutar el servidor o esperar a que el error suceda
	log.Fatal(s.server.ListenAndServe())
}