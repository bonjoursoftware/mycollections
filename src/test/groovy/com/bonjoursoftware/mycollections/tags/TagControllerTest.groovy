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
package com.bonjoursoftware.mycollections.tags

import com.bonjoursoftware.mycollections.collector.Collector
import io.micronaut.security.authentication.Authentication
import spock.lang.Specification

class TagControllerTest extends Specification {

    private static final Collector A_COLLECTOR = new Collector(username: 'collector')
    private static final String A_TAG = 'tag'

    private Authentication authentication
    private TagController tagController
    private TagRepository tagRepository

    void setup() {
        authentication = Mock(Authentication) {
            getAttributes() >> [username: A_COLLECTOR.username]
        }
        tagRepository = Mock()
        tagController = new TagController(tagRepository: tagRepository)
    }

    def 'Find by collector delegates to tag repository'() {
        when:
        tagController.findByCollector(authentication)

        then:
        1 * tagRepository.findByCollector(A_COLLECTOR)
    }

    def 'Find by tag and collector delegates to tag repository'() {
        when:
        tagController.findByTagAndCollector(A_TAG, authentication)

        then:
        1 * tagRepository.findByTagAndCollector(A_TAG, A_COLLECTOR)
    }
}
