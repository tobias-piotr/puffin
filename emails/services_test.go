package emails

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// DummyEmailRepository is a dummy email repository that has a simplest implementation
// to allow testing business logic.
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

// DummyEmailClient is a dummy email client that saves emails in memory.
type DummyEmailClient struct {
	emails [][]byte
}

func (c *DummyEmailClient) BuildEmail(tmpl string, context map[string]any) ([]byte, error) {
	return []byte(tmpl), nil
}

func (c *DummyEmailClient) SendEmail(receivers []string, subject string, email []byte) error {
	c.emails = append(c.emails, email)
	return nil
}

func TestCreateTemplate(t *testing.T) {
	srv := NewEmailService(&DummyEmailRepository{}, &DummyEmailClient{})
	result, error := srv.CreateNewTemplate(&TemplateData{"Test Name", "<h1>Test Content</h1>"})
	assert.Nil(t, error)
	assert.Equal(t, "Test Name", result.Name)
	assert.Equal(t, "<h1>Test Content</h1>", result.Content)
}

func TestCreateDuplicateTemplate(t *testing.T) {
	repo := &DummyEmailRepository{[]Template{{uuid.New(), time.Now(), "Test Name", "<h1>Test Content</h1>"}}}
	srv := NewEmailService(repo, &DummyEmailClient{})
	_, error := srv.CreateNewTemplate(&TemplateData{"Test Name", "<h1>Test Content</h1>"})
	assert.NotNil(t, error)
	assert.Equal(t, "Template with name Test Name already exists", error.Error())
}

func TestGetTemplates(t *testing.T) {
	repo := &DummyEmailRepository{[]Template{{uuid.New(), time.Now(), "Test Name", "<h1>Test Content</h1>"}}}
	srv := NewEmailService(repo, &DummyEmailClient{})
	result, error := srv.GetTemplates()
	assert.Nil(t, error)
	assert.Equal(t, 1, len(result))
}

func TestSendEmail(t *testing.T) {
	client := &DummyEmailClient{}
	repo := &DummyEmailRepository{[]Template{{uuid.New(), time.Now(), "Test Name", "<h1>Test Content</h1>"}}}
	srv := NewEmailService(repo, client)
	error := srv.SendEmail(&EmailData{TemplateName: "Test Name", Subject: "Test", To: []string{"test@gmail.com"}})
	assert.Nil(t, error)
}

func TestSendEmailNoTemplate(t *testing.T) {
	client := &DummyEmailClient{}
	repo := &DummyEmailRepository{[]Template{}}
	srv := NewEmailService(repo, client)
	error := srv.SendEmail(&EmailData{TemplateName: "Test Name", Subject: "Test", To: []string{"test@gmail.com"}})
	assert.NotNil(t, error)
	assert.Equal(t, "Template with name Test Name does not exist", error.Error())
}
