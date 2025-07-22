package com.bonjoursoftware.mycollections.notification

import groovy.transform.CompileStatic
import jakarta.activation.DataHandler
import jakarta.mail.Authenticator
import jakarta.mail.Message
import jakarta.mail.PasswordAuthentication
import jakarta.mail.Session
import jakarta.mail.internet.InternetAddress
import jakarta.mail.internet.MimeBodyPart
import jakarta.mail.internet.MimeMessage
import jakarta.mail.internet.MimeMultipart
import jakarta.mail.util.ByteArrayDataSource

import static jakarta.mail.Transport.send
import static java.time.LocalDate.now

@CompileStatic
class GmailNotificationService implements NotificationService {

    private static final String ATTACHMENT_EXTENSION = 'json'
    private static final String HTML_CONTENT_TYPE = 'text/html'
    private static final String JSON_CONTENT_TYPE = 'application/json'
    private static final String SOURCE_NAME = 'MyCollections'

    private static final int MAX_BODY_LENGTH = 4096

    private static final Properties PROPERTIES = new Properties().tap {
        put("mail.smtp.host", "smtp.gmail.com")
        put("mail.smtp.port", "587")
        put("mail.smtp.auth", "true")
        put("mail.smtp.starttls.enable", "true")
        put("mail.smtp.ssl.protocols", "TLSv1.2")
        put("mail.smtp.connectiontimeout", "10000") // seconds
        put("mail.smtp.timeout", "10000") // seconds
    }

    private final Session SESSION = Session.getInstance(PROPERTIES, new Authenticator() {
        @Override
        protected PasswordAuthentication getPasswordAuthentication() {
            return new PasswordAuthentication(source, apiKey)
        }
    })

    @Override
    void notify(String title, String body, String recipient) {
        send(new MimeMessage(SESSION).tap {
            setFrom(new InternetAddress(source, SOURCE_NAME))
            addRecipient(Message.RecipientType.TO, new InternetAddress(recipient))
            setSubject(title)

            setContent(new MimeMultipart().tap {
                addBodyPart(new MimeBodyPart().tap {
                    body.size() > MAX_BODY_LENGTH ? setText(title) : setContent(body, HTML_CONTENT_TYPE)
                })

                if (body.size() > MAX_BODY_LENGTH) {
                    addBodyPart(new MimeBodyPart().tap {
                        setDataHandler(new DataHandler(new ByteArrayDataSource(body, JSON_CONTENT_TYPE)))
                        setFileName("${title.toLowerCase().replace(' ', '_')}_${now().toString()}.${ATTACHMENT_EXTENSION}")
                    })
                }
            })
        })
    }
}
