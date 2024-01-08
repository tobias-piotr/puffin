package api

import (
	"net/http"

	"puffin/emails"
	"puffin/libs/api"
)

type EmailHandler struct {
	service emails.EmailService
}

// @Summary Create template
// @Description Create a new email template
// @Tags templates
// @Param data body emails.TemplateData true "Request body"
// @Success 201 {object} emails.Template
// @Failure 400 {object} api.APIError
// @Router /api/v1/templates [post]
func (h *EmailHandler) CreateNewTemplate(w http.ResponseWriter, r *http.Request) {
	data := &emails.TemplateData{}
	if err := api.DecodeAndValidate(r, data); err != nil {
		api.RespondWithErr(w, err)
		return
	}

	t, err := h.service.CreateNewTemplate(data)
	if err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusCreated, t)
}

// @Summary Get templates
// @Description Get a list of existing email templates
// @Tags templates
// @Success 200 {array} emails.Template
// @Failure 400 {object} api.APIError
// @Router /api/v1/templates [get]
func (h *EmailHandler) GetTemplates(w http.ResponseWriter, _ *http.Request) {
	t, err := h.service.GetTemplates()
	if err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusOK, t)
}

// @Summary Send email
// @Description Send an email using a template
// @Tags emails
// @Param data body emails.EmailData true "Request body"
// @Success 200
// @Router /api/v1/emails [post]
func (h *EmailHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	data := &emails.EmailData{}
	if err := api.DecodeAndValidate(r, data); err != nil {
		api.RespondWithErr(w, err)
		return
	}

	if err := h.service.SendEmail(data); err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusOK, nil)
}

// @Summary Get emails
// @Description Get a list of sent emails
// @Tags emails
// @Success 200 {array} emails.Email
// @Router /api/v1/emails [get]
func (h *EmailHandler) GetEmails(w http.ResponseWriter, r *http.Request) {
	emails, err := h.service.GetEmails()
	if err != nil {
		api.RespondWithErr(w, err)
		return
	}

	api.Respond(w, http.StatusOK, emails)
}
