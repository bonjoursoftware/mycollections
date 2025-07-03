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

import com.bonjoursoftware.mycollections.collector.Collector
import io.micronaut.security.authentication.Authentication
import spock.lang.Specification

class ItemControllerTest extends Specification {

    private static final Collector A_COLLECTOR = new Collector(friendlyname: 'friend', roles: ['a-role'], username: 'collector')
    private static final Item AN_ITEM = new Item(name: 'an item')
    private static final String AN_ID = 'an id'
    private static final String A_NAME = 'a name'

    private Authentication authentication
    private ItemController itemController
    private ItemRepository itemRepository

    void setup() {
        authentication = Mock(Authentication) {
            getAttributes() >> [collection: A_COLLECTOR.collection, friendlyname: A_COLLECTOR.friendlyname]
            getName() >> A_COLLECTOR.username
            getRoles() >> A_COLLECTOR.roles
        }
        itemRepository = Mock()
        itemController = new ItemController(itemRepository: itemRepository)
    }

    def 'Find by collector delegates to item repository'() {
        when:
        itemController.findByCollector(authentication)

        then:
        1 * itemRepository.findByCollector(A_COLLECTOR)
    }

    def 'Find by id and collector delegates to item repository'() {
        when:
        itemController.findByIdAndCollector(AN_ID, authentication)

        then:
        1 * itemRepository.findByIdAndCollector(AN_ID, A_COLLECTOR)
    }

    def 'Find by name and collector delegates to item repository'() {
        when:
        itemController.findByNameAndCollector(A_NAME, authentication)

        then:
        1 * itemRepository.findByNameAndCollector(A_NAME, A_COLLECTOR)
    }

    def 'Upsert item delegates to item repository'() {
        when:
        itemController.upsert(AN_ITEM, authentication)

        then:
        1 * itemRepository.upsert(AN_ITEM, A_COLLECTOR)
    }

    def 'Delete item delegates to item repository'() {
        when:
        itemController.delete(AN_ITEM, authentication)

        then:
        1 * itemRepository.delete(AN_ITEM, A_COLLECTOR)
    }
}
