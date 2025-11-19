/**
 * MyCollections - Keep track of collections of any kind
 *
 * https://github.com/bonjoursoftware/mycollections
 *
 * Copyright (C) 2020 - 2025 Bonjour Software Limited
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
import spock.lang.Specification

class CollectorAuthenticationExtractorTest extends Specification {

    private static final Collector A_COLLECTOR = new Collector(collection: 'a-collection', friendlyname: 'friend', roles: ['a-role'], username: 'collector')

    def 'Throws an exception when no collector is authenticated'() {
        when:
        CollectorAuthenticationExtractor.getCollector(null)

        then:
        thrown(CollectorAuthenticationExtractor.UnknownCollectorException)
    }

    def 'Throws an exception when collector username is not known'() {
        given:
        def authentication = Mock(Authentication) {
            getAttributes() >> [:]
        }

        when:
        CollectorAuthenticationExtractor.getCollector(authentication)

        then:
        thrown(CollectorAuthenticationExtractor.UnknownCollectorException)
    }

    def 'Returns collector when collector is authenticated and their username is known'() {
        given:
        def authentication = Mock(Authentication) {
            getAttributes() >> [collection: A_COLLECTOR.collection, friendlyname: A_COLLECTOR.friendlyname]
            getName() >> A_COLLECTOR.username
            getRoles() >> A_COLLECTOR.roles
        }

        expect:
        CollectorAuthenticationExtractor.getCollector(authentication) == A_COLLECTOR
    }
}
