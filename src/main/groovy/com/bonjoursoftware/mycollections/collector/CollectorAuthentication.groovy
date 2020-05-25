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
package com.bonjoursoftware.mycollections.collector

import com.bonjoursoftware.mycollections.notification.NotificationService
import groovy.transform.CompileStatic
import io.micronaut.security.authentication.AuthenticationFailed
import io.micronaut.security.authentication.AuthenticationProvider
import io.micronaut.security.authentication.AuthenticationRequest
import io.micronaut.security.authentication.AuthenticationResponse
import io.micronaut.security.authentication.UserDetails
import io.reactivex.Flowable
import org.reactivestreams.Publisher

import javax.inject.Inject
import javax.inject.Singleton

@CompileStatic
@Singleton
class CollectorAuthentication implements AuthenticationProvider {

    private static final String AUTH_FAILURE = 'Authentication failure'

    @Inject
    private CollectorRepository collectorRepository

    @Inject
    private NotificationService notificationService

    @Override
    Publisher<AuthenticationResponse> authenticate(AuthenticationRequest authenticationRequest) {
        if (collectorRepository.exists(authenticationRequest.identity as String, authenticationRequest.secret as String)) {
            return Flowable.just(new UserDetails((String) authenticationRequest.identity, [])) as Flowable<AuthenticationResponse>
        } else {
            notificationService.notify(AUTH_FAILURE, "Authentication failure for the following idendity: ${authenticationRequest.identity}")
            Flowable.just(new AuthenticationFailed()) as Flowable<AuthenticationResponse>
        }
    }
}
