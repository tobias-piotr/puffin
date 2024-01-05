package database

import (
	"puffin/emails"

	"github.com/jmoiron/sqlx"
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
	RETURNING id, name, content;
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
	SELECT id, name, content
	FROM templates;
	`
	templates := []emails.Template{}
	err := r.db.Select(&templates, query)
	return templates, err
}
