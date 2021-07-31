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
import com.bonjoursoftware.mycollections.token.TokenService
import groovy.transform.CompileStatic
import io.micronaut.context.annotation.Property

import javax.inject.Inject
import javax.inject.Singleton

import static java.net.URLEncoder.encode

@CompileStatic
@Singleton
class LoginService {

    private static final String AT = '@'
    private static final String LOGIN_LINK_TITLE = 'MyCollections Login Link'
    private static final String LOGIN_LINK_PLACEHOLDER = 'loginLinkPlaceholder'
    private static final String USERNAME_PLACEHOLDER = 'usernamePlaceholder'
    private static final String URL_ENCODE_CHARSET = 'UTF-8'

    @Inject
    private CollectorRepository collectorRepository

    @Inject
    private NotificationService notificationService

    @Inject
    private TokenService tokenService

    @Property(name = 'host.domain')
    private String hostDomain

    void request(String username) {
        tokenService.generate().with { token ->
            collectorRepository.upsert(
                    buildCollector(username, token.hash).tap {
                        it.collection = username
                        it.roles = [CollectorRoles.READ, CollectorRoles.WRITE].toSet()
                    }
            )
            notificationService.notify(LOGIN_LINK_TITLE, loginBody(username, token.secret), username)
        }
    }

    void share(String delegate, Collector collector) {
        tokenService.generate().with { token ->
            collectorRepository.upsert(
                    buildCollector(delegate, token.hash).tap {
                        it.collection = collector.collection
                        it.roles = [CollectorRoles.READ].toSet()
                    }
            )
            notificationService.notify(LOGIN_LINK_TITLE, loginBody(delegate, token.secret), delegate)
        }
    }

    private Collector buildCollector(String username, String hash) {
        new Collector(username: username, friendlyname: resolveFriendlyName(username), hash: hash)
    }

    private String resolveFriendlyName(String username) {
        username.contains(AT) ? username.split(AT).first() : username
    }

    private String loginBody(String username, String secret) {
        LOGIN_BODY_TEMPLATE
                .replace(LOGIN_LINK_PLACEHOLDER, loginLink(username, secret))
                .replace(USERNAME_PLACEHOLDER, username)
    }

    private String loginLink(String username, String secret) {
        "${hostDomain}/#!/login/${urlEncode(username)}/${urlEncode(secret)}"
    }

    private static final String urlEncode(String input) {
        encode(input, URL_ENCODE_CHARSET)
    }

    private static final String LOGIN_BODY_TEMPLATE = '''<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>MyCollections Login Link</title>
</head>
<body>
<p><a href="loginLinkPlaceholder">Click here</a> to log into MyCollections as usernamePlaceholder</p>
</body>
</html>
'''
}
