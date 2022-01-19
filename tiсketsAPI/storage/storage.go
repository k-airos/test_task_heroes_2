package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Instance of storage
type Storage struct {
	config *Config
	// DataBase FileDescriptor
	db *sql.DB
	//Subfield for repo interfaceing (model article)
	ticketRepository *TicketRepository
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

//Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}

func (s *Storage) Ticket() *TicketRepository {
	if s.ticketRepository != nil {
		return s.ticketRepository
	}
	s.ticketRepository = &TicketRepository{
		storage: s,
	}
	return s.ticketRepository
}
