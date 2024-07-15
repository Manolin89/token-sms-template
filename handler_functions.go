package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Token de verificacion para la configuracion
var TOKEN_AUTH = "YAGANASTE"

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/webhook", a.VerifyToken).Methods(http.MethodGet)
	r.HandleFunc("/webhook", a.ReceivedMessages).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) VerifyToken(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query()

	mode := key.Get("hub.mode")
	token := key.Get("hub.verify_token")
	challenge := key.Get("hub.challenge")

	if len(mode) > 0 && len(token) > 0 {

		if mode == "subscribe" && token == TOKEN_AUTH {

			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(challenge))
			if err != nil {
				return
			}

		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode("invalid token")
			if err != nil {
				return
			}
		}
	}
}

func (a *api) ReceivedMessages(w http.ResponseWriter, r *http.Request) {

	log.Printf("Leer mensaje recivido")
	var messageReceivedJson ReceivedDataMessage
	err := json.NewDecoder(r.Body).Decode(&messageReceivedJson)
	log.Printf("mensaje recivido: %v", messageReceivedJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode("error to received message information receivd")
		if err != nil {
			return
		}
		return
	}

	if messageReceivedJson.NombreCliente != "" && messageReceivedJson.Token != "" && messageReceivedJson.Telefono != "" {
		log.Printf("Construir mensaje de primer contacto")
		parameter := []Parameters{
			{
				Type: "text",
				Text: messageReceivedJson.NombreCliente,
			},
			{
				Type: "text",
				Text: messageReceivedJson.Token,
			},
		}
		components := []Components{
			{
				Type:       "body",
				Parameters: parameter,
			},
		}
		languaje := Language{
			Code: "es_MX",
		}
		template := Template{
			Name:       "account_confirm",
			Language:   languaje,
			Components: components,
		}
		dataSendingToken := DataSendingToken{
			MessagingProduct: "whatsapp",
			RecipientType:    "individual",
			To:               messageReceivedJson.Telefono,
			Type:             "template",
			Template:         template,
		}

		SendMessage(w, r, dataSendingToken)
	}

}

func SendMessage(w http.ResponseWriter, _ *http.Request, dataSending DataSendingToken) {
	log.Printf("Sending message")
	ws := &apiWhatsapp{}
	_, status, err := ws.ResquestWhatsapp(dataSending)
	if err != nil || status != 200 {
		return
	}

	log.Printf("Proceso de enviado terminado con estatus: %d", status)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("First Message Code Sending Successfully")
	if err != nil {
		return
	}
}
