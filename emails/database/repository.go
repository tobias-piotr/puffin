package database

import (
	"fmt"
	"strings"

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
	`, convertMapToArgsStr(filters))
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return []emails.Template{}, err
	}

	var templates []emails.Template
	err = stmt.Select(&templates, filters)
	return templates, err
}

func convertMapToArgsStr(filters map[string]any) string {
	args := []string{}
	for key := range filters {
		args = append(args, fmt.Sprintf("%s = :%s", key, key))
	}
	return strings.Join(args, " AND ")
}
