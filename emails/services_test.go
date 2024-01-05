package emails

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type DummyEmailRepository struct {
	templates []Template
}

func (r *DummyEmailRepository) CreateNewTemplate(data *TemplateData) (*Template, error) {
	return &Template{uuid.New(), time.Now(), data.Name, data.Content}, nil
}

func (r *DummyEmailRepository) GetTemplates() ([]Template, error) {
	return r.templates, nil
}

func (r *DummyEmailRepository) FilterTemplates(filters EmailFilters) ([]Template, error) {
	return r.templates, nil
}

type DummyEmailClient struct{}

func (c *DummyEmailClient) SendEmail(data *EmailData) error {
	return nil
}

func TestCreateTemplate(t *testing.T) {
	srv := NewEmailService(&DummyEmailRepository{}, &DummyEmailClient{})
	result, error := srv.CreateNewTemplate(&TemplateData{"Test Name", "<h1>Test Content</h1>"})
	assert.Nil(t, error)
	assert.Equal(t, "Test Name", result.Name)
	assert.Equal(t, "<h1>Test Content</h1>", result.Content)
}

func TestGetTemplates(t *testing.T) {
	repo := &DummyEmailRepository{[]Template{{uuid.New(), time.Now(), "Test Name", "<h1>Test Content</h1>"}}}
	srv := NewEmailService(repo, &DummyEmailClient{})
	result, error := srv.GetTemplates()
	assert.Nil(t, error)
	assert.Equal(t, 1, len(result))
}
