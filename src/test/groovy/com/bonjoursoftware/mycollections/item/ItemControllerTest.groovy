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
package com.bonjoursoftware.mycollections.item

import io.micronaut.security.authentication.Authentication
import spock.lang.Specification

class ItemControllerTest extends Specification {

    private static final String A_COLLECTOR = 'collector'
    private static final String A_TAG = 'tag'
    private static final Item AN_ITEM = new Item(name: 'an item')

    private Authentication authentication
    private ItemController itemController
    private ItemRepository itemRepository

    void setup() {
        authentication = Mock(Authentication) {
            getAttributes() >> [username: A_COLLECTOR]
        }
        itemRepository = Mock()
        itemController = new ItemController(itemRepository: itemRepository)
    }

    def 'Find by collector delegates to item service'() {
        when:
        itemController.findByCollector(authentication)

        then:
        1 * itemRepository.findByCollector(A_COLLECTOR)
    }

    def 'Find by tag and collector delegates to item service'() {
        when:
        itemController.findByTagAndCollector(A_TAG, authentication)

        then:
        1 * itemRepository.findByTagAndCollector(A_TAG, A_COLLECTOR)
    }

    def 'Upsert item delegates to item repository'() {
        when:
        itemController.upsert(AN_ITEM, authentication)

        then:
        1 * itemRepository.upsert(AN_ITEM, A_COLLECTOR)
    }

    def 'Delete item delegates to item service'() {
        when:
        itemController.delete(AN_ITEM, authentication)

        then:
        1 * itemRepository.delete(AN_ITEM, A_COLLECTOR)
    }
}
