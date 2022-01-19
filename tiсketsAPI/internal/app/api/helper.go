package api

import (
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"tickets/storage"
)

//Пытаемся откунфигурировать наш API инстанс (а конкретнее - поле logger)
func (a *API) configreLoggerField() error {
	logLevel, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(logLevel)
	return nil
}

//Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configreRouterField() {
	a.router.HandleFunc("/tickets", a.GetAllTickets).Methods("GET")
	a.router.HandleFunc("/ticket/{id}", a.GetTicketById).Methods("GET")
	a.router.HandleFunc("/ticket/{id}", a.DeleteTicketById).Methods("DELETE")
	a.router.HandleFunc("/ticket", a.PostTicket).Methods("POST")
	a.router.HandleFunc("/ticket/{id}", a.UpdateTicketById).Methods("PUT")

}

//Пытаемся отконфигурировать наше хранилище (storage API)
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
