package com.bonjoursoftware.mycollections.notification

import jakarta.activation.DataHandler
import jakarta.mail.Authenticator
import jakarta.mail.Message
import jakarta.mail.PasswordAuthentication
import jakarta.mail.Session
import jakarta.mail.Transport
import jakarta.mail.internet.InternetAddress
import jakarta.mail.internet.MimeBodyPart
import jakarta.mail.internet.MimeMessage
import jakarta.mail.internet.MimeMultipart
import jakarta.mail.util.ByteArrayDataSource

import static java.time.LocalDate.now

class GmailNotificationService implements NotificationService {

    private static final String HTML_CONTENT_TYPE = 'text/html'
    private static final String ATTACHMENT_EXTENSION = 'json'
    private static final String JSON_CONTENT_TYPE = 'application/json'

    private static final int MAX_BODY_LENGTH = 4096

    @Override
    void notify(String title, String body) {
        sendMail(title, body)
    }

    @Override
    void notify(String title, String body, String recipient) {
        sendMail(title, body, recipient)
    }

    private sendMail(String title, String body, String recipient = target) {
        Properties props = new Properties()
        props.put("mail.smtp.host", "smtp.gmail.com")
        props.put("mail.smtp.port", "587")
        props.put("mail.smtp.auth", "true")
        props.put("mail.smtp.starttls.enable", "true")
        props.put("mail.smtp.ssl.protocols", "TLSv1.2")
        props.put("mail.smtp.connectiontimeout", "5000") // seconds
        props.put("mail.smtp.timeout", "5000") // seconds

        def sesh = Session.getInstance(props, new Authenticator() {
            @Override
            protected PasswordAuthentication getPasswordAuthentication() {
                return new PasswordAuthentication(source, apiKey)
            }
        })

        def msg = new MimeMessage(sesh)
        msg.setFrom(new InternetAddress(source))
        msg.addRecipient(Message.RecipientType.TO, new InternetAddress(recipient))
        msg.setSubject(title)

        def multipart = new MimeMultipart()

        def msgBodyPart = new MimeBodyPart()

        if (body.size() > MAX_BODY_LENGTH) {
            msgBodyPart.setText(title)
        } else {
            msgBodyPart.setContent(body, HTML_CONTENT_TYPE)
        }

        multipart.addBodyPart(msgBodyPart)

        if (body.size() > MAX_BODY_LENGTH) {
            def attachPart = new MimeBodyPart()
            def dataSource = new ByteArrayDataSource(body, JSON_CONTENT_TYPE)
            attachPart.setDataHandler(new DataHandler(dataSource))
            attachPart.setFileName("${title.toLowerCase().replace(' ', '_')}_${now().toString()}.${ATTACHMENT_EXTENSION}")
            multipart.addBodyPart(attachPart)
        }

        msg.setContent(multipart)
        Transport.send(msg)
    }
}
