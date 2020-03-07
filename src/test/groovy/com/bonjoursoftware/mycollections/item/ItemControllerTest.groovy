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

    private static final String A_COLLECTOR = 'collector@dummy-domain.com'
    private static final ItemDTO AN_ITEM = new ItemDTO(name: 'an item')

    private Authentication authentication
    private ItemController itemController
    private ItemService itemService

    void setup() {
        authentication = Mock(Authentication) {
            getAttributes() >> [email: A_COLLECTOR]
        }
        itemService = Mock()
        itemController = new ItemController(itemService: itemService)
    }

    def 'Find by collector delegates to item service'() {
        when:
        itemController.findByCollector(authentication)

        then:
        1 * itemService.findByCollector(A_COLLECTOR)
    }

    def 'Create item delegates to item service'() {
        when:
        itemController.create(AN_ITEM, authentication)

        then:
        1 * itemService.create(AN_ITEM, A_COLLECTOR)
    }

    def 'Update item delegates to item service'() {
        when:
        itemController.update(AN_ITEM, authentication)

        then:
        1 * itemService.update(AN_ITEM, A_COLLECTOR)
    }
}
