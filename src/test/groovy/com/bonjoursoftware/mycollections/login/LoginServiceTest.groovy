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
import com.bonjoursoftware.mycollections.collector.CollectorRepository
import com.bonjoursoftware.mycollections.notification.NotificationService
import com.bonjoursoftware.mycollections.token.Token
import com.bonjoursoftware.mycollections.token.TokenService
import spock.lang.Specification

class LoginServiceTest extends Specification {

    private static final String A_USERNAME = 'username@domain.com'
    private static final String A_FRIENDLYNAME = 'username'
    private static final String A_HOST_DOMAIN = 'domain'
    private static final Token A_TOKEN = new Token(secret: 'a secret/', hash: 'a hash')

    private LoginService loginService
    private CollectorRepository collectorRepository
    private NotificationService notificationService
    private TokenService tokenService

    void setup() {
        collectorRepository = Mock()
        notificationService = Mock()
        tokenService = Mock(TokenService) {
            generate() >> A_TOKEN
        }
        loginService = new LoginService(
                collectorRepository: collectorRepository,
                notificationService: notificationService,
                tokenService: tokenService,
                hostDomain: A_HOST_DOMAIN)
    }

    def 'Request authentication link refreshes secret and notifies collector'() {
        when:
        loginService.request(A_USERNAME)

        then:
        1 * collectorRepository.upsert({ Collector collector ->
            collector.username == A_USERNAME
                    && collector.hash == A_TOKEN.hash
                    && collector.friendlyname == A_FRIENDLYNAME
        })
        1 * notificationService.notify(
                'MyCollections Login Link',
                { String body -> body.contains("${A_HOST_DOMAIN}/#!/login/username%40domain.com/a+secret%2F") },
                A_USERNAME)
    }
}
