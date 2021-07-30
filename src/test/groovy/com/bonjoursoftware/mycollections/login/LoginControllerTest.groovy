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
package com.bonjoursoftware.mycollections.login

import com.bonjoursoftware.mycollections.collector.Collector
import com.bonjoursoftware.mycollections.notification.NotificationService
import io.micronaut.security.authentication.Authentication
import spock.lang.Specification
import spock.lang.Subject

class LoginControllerTest extends Specification {

    private static final Collector A_COLLECTOR = new Collector(friendlyname: 'friend', roles: ['a-role'], username: 'collector')
    private static final String SOME_ALLOWED_SUFFIXES = 'allowed-domain.com,trusted-domain.com'

    private Authentication authentication
    private LoginService loginService
    private NotificationService notificationService

    @Subject
    private LoginController loginController

    void setup() {
        authentication = Mock(Authentication) {
            getAttributes() >> [friendlyname: A_COLLECTOR.friendlyname, roles: A_COLLECTOR.roles, username: A_COLLECTOR.username]
        }
        loginService = Mock()
        notificationService = Mock()
        loginController = new LoginController(
                loginService: loginService,
                notificationService: notificationService,
                allowedSuffixes: SOME_ALLOWED_SUFFIXES)
    }

    def 'Request login link delegates to login service when username is allowed'() {
        when:
        loginController.requestLogin(username)

        then:
        1 * loginService.request(username)

        where:
        username                     | _
        'someone@allowed-domain.com' | _
        'someone@trusted-domain.com' | _
    }

    def 'Request login notifies on unauthorised attempt'() {
        when:
        loginController.requestLogin(username)

        then:
        0 * loginService._
        1 * notificationService.notify(_, _)

        where:
        username                               | _
        null                                   | _
        ''                                     | _
        'someone'                              | _
        'someone@'                             | _
        'someone@denied-domain.com'            | _
        '@whatever'                            | _
        '@'                                    | _
        'someone@with@weird-email-address.com' | _
    }

    def 'Request share link delegates to login service when username is valid email'() {
        when:
        loginController.share(username, authentication)

        then:
        1 * loginService.share(username, A_COLLECTOR)

        where:
        username                     | _
        'someone@some-domain.com'    | _
        'someone@trusted-domain.com' | _
    }

    def 'Request share link does nothing when username is invalid email'() {
        when:
        loginController.share(username, authentication)

        then:
        0 * loginService._

        where:
        username                               | _
        null                                   | _
        ''                                     | _
        'someone'                              | _
        'someone@'                             | _
        '@whatever'                            | _
        '@'                                    | _
        'someone@with@weird-email-address.com' | _
    }
}
