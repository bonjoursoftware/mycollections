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

import com.bonjoursoftware.mycollections.collector.CollectorRoles
import com.bonjoursoftware.mycollections.notification.NotificationService
import groovy.transform.CompileStatic
import groovy.util.logging.Slf4j
import io.micronaut.context.annotation.Property
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Post
import io.micronaut.http.annotation.QueryValue
import io.micronaut.security.annotation.Secured
import io.micronaut.security.authentication.Authentication
import io.micronaut.security.rules.SecurityRule

import javax.annotation.security.RolesAllowed
import jakarta.inject.Inject

import static com.bonjoursoftware.mycollections.collector.CollectorAuthenticationExtractor.getCollector

@CompileStatic
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller('/api/v1/login')
@Slf4j
class LoginController {

    private static final String AT = '@'
    private static final String SUFFIX_SEPARATOR = ','
    private static final String LOGIN_REQUEST_NOTIFICATION = 'Unauthorised Login Request'

    @Inject
    private LoginService loginService

    @Inject
    private NotificationService notificationService

    @Property(name = 'username.allowed-suffixes')
    private String allowedSuffixes

    @Get('/{username}')
    @Secured(SecurityRule.IS_ANONYMOUS)
    void requestLogin(@QueryValue('username') String username) {
        if (looksLikeEmailAddress(username) && hasAllowedSuffix(username))
            loginService.request(username)
        else
            notificationService.notify(LOGIN_REQUEST_NOTIFICATION, "[${username}] made an unauthorised login request")
    }

    @Post('/{delegate}')
    @RolesAllowed(CollectorRoles.WRITE)
    void share(String delegate, Authentication authentication) {
        if (looksLikeEmailAddress(delegate))
            loginService.share(delegate, getCollector(authentication))
    }

    private boolean looksLikeEmailAddress(String username) {
        username?.split(AT)?.findAll()?.size() == 2
    }

    private boolean hasAllowedSuffix(String username) {
        username?.endsWithAny(allowedSuffixes.split(SUFFIX_SEPARATOR))
    }
}
