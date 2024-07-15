package main

type DataSendingToken struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Template         Template `json:"template"`
}

type Template struct {
	Name       string       `json:"name"`
	Language   Language     `json:"language"`
	Components []Components `json:"components"`
}

type Language struct {
	Code string `json:"code"`
}

type Components struct {
	Type       string       `json:"type"`
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type ReceivedDataMessage struct {
	NombreCliente string `json:"nombre_cliente"`
	Telefono      string `json:"telefono"`
	Token         string `json:"token"`
}

type Error struct {
	Error string `json:"error"`
}
