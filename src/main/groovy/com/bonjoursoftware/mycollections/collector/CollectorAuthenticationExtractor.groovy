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

import io.micronaut.security.authentication.Authentication

import static java.util.Optional.ofNullable

class CollectorAuthenticationExtractor {

    private static final COLLECTION = 'collection'
    private static final FRIENDLY_NAME = 'friendlyname'
    private static final ROLES = 'roles'
    private static final USERNAME = 'username'

    static Collector getCollector(Authentication authentication) {
        ofNullable(toCollector(authentication)).orElseThrow({ new UnknownCollectorException() })
    }

    private static Collector toCollector(Authentication authentication) {
        authentication?.attributes?.get(USERNAME)?.with { username ->
            new Collector(
                    collection: authentication.attributes.get(COLLECTION),
                    friendlyname: authentication.attributes.get(FRIENDLY_NAME) ?: 'anonymous',
                    roles: authentication.attributes.get(ROLES) as Set<String> ?: Collections.<String>emptySet(),
                    username: username
            )
        }
    }

    static class UnknownCollectorException extends RuntimeException {
    }
}
