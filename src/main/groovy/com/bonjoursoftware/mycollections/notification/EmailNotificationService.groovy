/**
 * MyCollections - Keep track of collections of any kind
 *
 * https://github.com/bonjoursoftware/mycollections
 *
 * Copyright (C) 2020 Bonjour Software Limited
 *
 * https://bonjoursoftware.com/
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see
 * https://github.com/bonjoursoftware/mycollections/blob/master/LICENSE
 */
package com.bonjoursoftware.mycollections.notification

import com.sendgrid.Method
import com.sendgrid.Request
import com.sendgrid.SendGrid
import com.sendgrid.helpers.mail.Mail
import com.sendgrid.helpers.mail.objects.Attachments
import com.sendgrid.helpers.mail.objects.Content
import com.sendgrid.helpers.mail.objects.Email
import groovy.transform.CompileStatic

import javax.inject.Singleton

import static java.time.LocalDate.now
import static java.util.Base64.encoder

@CompileStatic
@Singleton
class EmailNotificationService implements NotificationService {

    private static final String ENDPOINT = 'mail/send'
    private static final String SOURCE_NAME = 'MyCollections'
    private static final String HTML_CONTENT_TYPE = 'text/html'
    private static final String CHARSET = 'UTF-8'
    private static final String ATTACHMENT_EXTENSION = 'json'
    private static final String ATTACHMENT_DISPOSITION = 'attachment'
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
        new SendGrid(apiKey).api(buildRequest(title, body, recipient))
    }

    private Request buildRequest(String title, String body, String recipient) {
        new Request().tap {
            setMethod(Method.POST)
            setEndpoint(ENDPOINT)
            setBody(buildPayload(title, body, recipient))
        }
    }

    private String buildPayload(String title, String body, String recipient) {
        body.size() < MAX_BODY_LENGTH ? buildPayloadWithBody(title, body, recipient) : buildPayloadWithAttachment(title, body, recipient)
    }

    private String buildPayloadWithBody(String title, String body, String recipient) {
        new Mail(
                new Email(name: SOURCE_NAME, email: source),
                title,
                new Email(recipient),
                new Content(HTML_CONTENT_TYPE, body)
        ).build()
    }

    private String buildPayloadWithAttachment(String title, String body, String recipient) {
        new Mail(
                new Email(name: SOURCE_NAME, email: source),
                title,
                new Email(recipient),
                new Content(HTML_CONTENT_TYPE, title)
        ).tap {
            addAttachments(new Attachments().tap {
                setContent(getEncoder().encodeToString(body.getBytes(CHARSET)))
                setContentId(title)
                setDisposition(ATTACHMENT_DISPOSITION)
                setFilename("${title.toLowerCase().replace(' ', '_')}_${now().toString()}.${ATTACHMENT_EXTENSION}")
                setType(JSON_CONTENT_TYPE)
            })
        }.build()
    }
}
