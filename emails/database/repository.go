package database

import (
	"fmt"

	"puffin/emails"
	"puffin/libs/database"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// EmailRepository handles all database operations for emails.
type EmailRepository struct {
	db *sqlx.DB
}

func NewEmailRepository(db *sqlx.DB) EmailRepository {
	return EmailRepository{db}
}

func (r EmailRepository) CreateNewTemplate(data *emails.TemplateData) (*emails.Template, error) {
	query := `
	INSERT INTO templates (name, content)
	VALUES (:name, :content)
	RETURNING id, created_at, name, content;
	`
	query, args, err := r.db.BindNamed(query, data)
	if err != nil {
		return &emails.Template{}, err
	}

	var template emails.Template
	err = r.db.QueryRowx(query, args...).StructScan(&template)
	return &template, err
}

func (r EmailRepository) GetTemplates() ([]emails.Template, error) {
	query := `
	SELECT id, created_at, name, content
	FROM templates;
	`
	templates := []emails.Template{}
	err := r.db.Select(&templates, query)
	return templates, err
}

func (r EmailRepository) FilterTemplates(filters emails.EmailFilters) ([]emails.Template, error) {
	query := fmt.Sprintf(`
	SELECT id, created_at, name, content
	FROM templates
	WHERE %s;
	`, database.ConvertMapToArgsStr(filters))
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return []emails.Template{}, err
	}

	var templates []emails.Template
	err = stmt.Select(&templates, filters)
	return templates, err
}

func (r EmailRepository) SaveEmail(data *emails.EmailData) (*emails.Email, error) {
	query := `
	INSERT INTO emails (template_name, recipients, subject, context)
	VALUES (:template_name, :recipients, :subject, :context)
	RETURNING id, created_at, template_name, recipients, subject, context;
	`
	dbData := EmailData{*data, pq.StringArray(data.Recipients), database.JSON(data.Context)}
	query, args, err := r.db.BindNamed(query, dbData)
	if err != nil {
		return &emails.Email{}, err
	}

	var email Email
	err = r.db.QueryRowx(query, args...).StructScan(&email)
	return email.ToEmail(), err
}

func (r EmailRepository) GetEmails() ([]emails.Email, error) {
	query := `
	SELECT id, created_at, template_name, recipients, subject, context
	FROM emails;
	`
	rows, err := r.db.Queryx(query)
	if err != nil {
		return []emails.Email{}, err
	}
	defer rows.Close()

	// Scan emails and convert them to domain objects
	res := []emails.Email{}
	for rows.Next() {
		var email Email
		err = rows.StructScan(&email)
		if err != nil {
			return []emails.Email{}, err
		}
		res = append(res, *email.ToEmail())
	}

	return res, nil
}
