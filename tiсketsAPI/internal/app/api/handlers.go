package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tickets/internal/app/models"
)

//Вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

//Возвращает все билеты из бд на данный момент
func (api *API) GetAllTickets(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get All tickets GET /tickets")
	tickets, err := api.storage.Ticket().SelectAll()
	if err != nil {
		api.logger.Info("Error while Tickets.SelectAll : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(tickets)
}

func (api *API) PostTicket(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Ticket POST /tickets")
	var ticket models.Ticket
	err := json.NewDecoder(req.Body).Decode(&ticket)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Ticket().Create(&ticket)
	if err != nil {
		api.logger.Info("Troubles while creating new ticket:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetTicketById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get ticket by ID /tickets/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	ticket, ok, err := api.storage.Ticket().FindTicketById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (tickets) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find ticket with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "ticket with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(ticket)
}

func (api *API) DeleteTicketById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete ticket by Id DELETE /tickets/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok, err := api.storage.Ticket().FindTicketById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (tickets) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find ticket with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Ticket with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.storage.Ticket().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while deleting database elemnt from table (tickets) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Ticket with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) UpdateTicketById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Update ticket by Id UPDATE /tickets/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, err := api.storage.Ticket().FindTicketById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (tickets) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find ticket with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Ticket with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	var ticket models.Ticket
	err = json.NewDecoder(req.Body).Decode(&ticket)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = api.storage.Ticket().UpdateTicket(id, &ticket)
	if err != nil {
		api.logger.Info("Troubles while updating database table (tickets) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	msg := Message{
		StatusCode: 200,
		Message:    fmt.Sprintf("Ticket with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}
