package main

import "net/http"

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPaylaod mailMessage

	err := app.readJson(w, r, requestPaylaod)

	if err != nil {
		app.errorJSON(w, err)

		return
	}

	msg := Message{
		From:    requestPaylaod.From,
		To:      requestPaylaod.To,
		Subject: requestPaylaod.Subject,
		Data:    requestPaylaod.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)

	if err != nil {
		app.errorJSON(w, err)

		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to " + requestPaylaod.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
