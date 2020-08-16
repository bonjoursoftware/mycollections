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
import com.sendgrid.helpers.mail.objects.Content
import com.sendgrid.helpers.mail.objects.Email
import groovy.transform.CompileStatic

import javax.inject.Singleton

@CompileStatic
@Singleton
class EmailNotificationService implements NotificationService {

    private static final String ENDPOINT = 'mail/send'
    private static final String SOURCE_NAME = 'MyCollections'
    private static final String CONTENT_TYPE = 'text/html'

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
        new Mail(new Email(name: SOURCE_NAME, email: recipient), title, new Email(recipient), new Content(CONTENT_TYPE, body)).build()
    }
}
