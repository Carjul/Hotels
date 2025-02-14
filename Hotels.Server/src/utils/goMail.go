package utils

import (
	"fmt"
	"os"

	"gopkg.in/mail.v2"
)

func SendMail(correo string, asunto string, htmlContent string) string {
	email_password := os.Getenv("EMAIL_PASSWORD")
	Email := os.Getenv("EMAIL")

	if email_password == "" || Email == "" {
		return "No se ha configurado el correo electrónico"
	}

	email := mail.NewMessage()
	email.SetHeader("From", "\"Soporte técnico hotels\""+"<"+Email+">")
	email.SetHeader("To", correo)
	email.SetHeader("Subject", asunto)
	email.SetBody("text/html", htmlContent)

	// Configuración del servidor SMTP
	smtpServer := mail.NewDialer("smtp.gmail.com", 587, Email, email_password)

	// Intentar enviar el correo electrónico
	err := smtpServer.DialAndSend(email)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return "Error al enviar el correo electrónico"
	}

	return "Correo electrónico enviado correctamente"
}
