package storage

import (
	"fmt"
	"log"
	"tickets/internal/app/models"
)

//Instance of Article repository (model interface)
type TicketRepository struct {
	storage *Storage
}

var (
	tableTickets string = "tickets"
)

func (tr *TicketRepository) Create(a *models.Ticket) (*models.Ticket, error) {
	query := fmt.Sprintf("INSERT INTO %s (film, screen_writer, ticket_content) VALUES ($1, $2, $3) RETURNING id", tableTickets)
	if err := tr.storage.db.QueryRow(query, a.Film, a.ScreenWriter, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}

	return a, nil

}

func (tr *TicketRepository) DeleteById(id int) (*models.Ticket, error) {
	ticket, ok, err := tr.FindTicketById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableTickets)
		_, err := tr.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return ticket, nil
}

func (tr *TicketRepository) FindTicketById(id int) (*models.Ticket, bool, error) {
	tickets, err := tr.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var ticketFinded *models.Ticket
	for _, a := range tickets {
		if a.ID == id {
			ticketFinded = a
			founded = true
			break
		}
	}
	return ticketFinded, founded, nil
}

func (tr *TicketRepository) SelectAll() ([]*models.Ticket, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableTickets)
	rows, err := tr.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := make([]*models.Ticket, 0)
	for rows.Next() {
		a := models.Ticket{}
		err := rows.Scan(&a.ID, &a.Film, &a.ScreenWriter, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		tickets = append(tickets, &a)
	}
	return tickets, nil
}

func (tr *TicketRepository) UpdateTicket(id int, t *models.Ticket) error {
	query := fmt.Sprintf("UPDATE %s SET film=$1, screen_writer=$2, ticket_content=$3  WHERE id=$4", tableTickets)
	_, err := tr.storage.db.Exec(query, t.Film, t.ScreenWriter, t.Content, id)
	if err != nil {
		return err
	}

	return nil
}
