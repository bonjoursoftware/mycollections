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
import spock.lang.Specification

class CollectorTest extends Specification {

    private static final String COLLECTOR_EMAIL = 'collector@dummy-domain.com'

    def 'Throws an exception when no collector is authenticated'() {
        when:
        Collector.getCollector(null)

        then:
        thrown(Collector.UnknownCollectorException)
    }

    def 'Throws an exception when collector email is not known'() {
        given:
        def authentication = Mock(Authentication) {
            getAttributes() >> [:]
        }

        when:
        Collector.getCollector(authentication)

        then:
        thrown(Collector.UnknownCollectorException)
    }

    def 'Returns collector email when collector is authenticated and their email is known'() {
        given:
        def authentication = Mock(Authentication) {
            getAttributes() >> [email: COLLECTOR_EMAIL]
        }

        expect:
        Collector.getCollector(authentication) == COLLECTOR_EMAIL
    }
}
