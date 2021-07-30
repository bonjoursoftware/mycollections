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
import com.bonjoursoftware.mycollections.collector.CollectorRoles
import com.bonjoursoftware.mycollections.notification.NotificationService
import com.bonjoursoftware.mycollections.token.Token
import com.bonjoursoftware.mycollections.token.TokenService
import spock.lang.Specification
import spock.lang.Subject

class LoginServiceTest extends Specification {

    private static final Collector A_COLLECTOR = new Collector(friendlyname: 'username', roles: ['a-role'], username: 'username@domain.com')
    private static final String A_DELEGATE = 'delegate@other-domain.com'
    private static final String A_HOST_DOMAIN = 'domain'
    private static final Token A_TOKEN = new Token(secret: 'a secret/', hash: 'a hash')

    private CollectorRepository collectorRepository
    private NotificationService notificationService
    private TokenService tokenService

    @Subject
    private LoginService loginService

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
        loginService.request(A_COLLECTOR.username)

        then:
        1 * collectorRepository.upsert({ Collector collector ->
            collector.username == A_COLLECTOR.username
                    && collector.hash == A_TOKEN.hash
                    && collector.friendlyname == A_COLLECTOR.friendlyname
        })
        1 * notificationService.notify(
                'MyCollections Login Link',
                { String body -> body.contains("${A_HOST_DOMAIN}/#!/login/username%40domain.com/a+secret%2F") },
                A_COLLECTOR.username
        )
    }

    def 'Request share link refreshes secret and notifies delegate collector'() {
        when:
        loginService.share(A_DELEGATE, A_COLLECTOR)

        then:
        1 * collectorRepository.upsert({ Collector delegate ->
            delegate.username == A_DELEGATE
                    && delegate.hash == A_TOKEN.hash
                    && delegate.friendlyname == 'delegate'
                    && delegate.collection == A_COLLECTOR.collection
                    && delegate.roles == [CollectorRoles.READ].toSet()
        })
        1 * notificationService.notify(
                'MyCollections Login Link',
                { String body -> body.contains("${A_HOST_DOMAIN}/#!/login/delegate%40other-domain.com/a+secret%2F") },
                A_DELEGATE
        )
    }
}
