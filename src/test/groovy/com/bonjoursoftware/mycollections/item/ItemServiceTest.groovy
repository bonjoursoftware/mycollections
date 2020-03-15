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

import spock.lang.Specification
import spock.lang.Unroll

class ItemServiceTest extends Specification {

    private ItemRepository itemRepository
    private ItemService itemService

    void setup() {
        itemRepository = Mock()
        itemService = new ItemService(itemRepository: itemRepository)
    }

    @Unroll
    def 'Find by collector [#collector] delegates to repository for collector [#collector]'() {
        when:
        itemService.findByCollector(collector)

        then:
        1 * itemRepository.findByCollector(collector)

        where:
        collector << ['bob', 'sam']
    }

    @Unroll
    def 'Create item [#name] for collector [#collector] delegates to repository for item [#name] and collector [#collector]'() {
        given:
        def item = new Item(name: name)

        when:
        itemService.create(item, collector)

        then:
        1 * itemRepository.save({ it.name == name && it.collector == collector } as Item)

        where:
        name           | collector
        'an item'      | 'bob'
        'an item'      | 'sam'
        'another item' | 'bob'
        'another item' | 'sam'
    }

    @Unroll
    def 'Update item with id [#id] for collector [#collector] delegates to repository when collector [#collector] owns an item with id [#id]'() {
        given:
        def item = new Item(id: id, name: name)

        when:
        itemService.update(item, collector)

        then:
        1 * itemRepository.existsByIdAndCollector(id, collector) >> Boolean.TRUE
        1 * itemRepository.update({ it.id == id && it.name == name && it.collector == collector } as Item)

        where:
        id | name           | collector
        1  | 'an item'      | 'bob'
        2  | 'another item' | 'bob'
    }

    @Unroll
    def 'Update item with id [#id] for collector [#collector] throws an exception when collector [#collector] does not own an item with id [#id]'() {
        given:
        def item = new Item(id: id, name: name)

        when:
        itemService.update(item, collector)

        then:
        1 * itemRepository.existsByIdAndCollector(id, collector) >> Boolean.FALSE
        0 * itemRepository.update(_)

        and:
        thrown(ItemService.OwnershipException)

        where:
        id | name           | collector
        1  | 'an item'      | 'bob'
        2  | 'another item' | 'bob'
    }

    @Unroll
    def 'Delete item with id [#id] for collector [#collector] delegates to repository when collector [#collector] owns an item with id [#id]'() {
        given:
        def item = new Item(id: id, name: name)

        when:
        itemService.delete(item, collector)

        then:
        1 * itemRepository.existsByIdAndCollector(id, collector) >> Boolean.TRUE
        1 * itemRepository.deleteById(id)

        where:
        id | name           | collector
        1  | 'an item'      | 'bob'
        2  | 'another item' | 'bob'
    }

    @Unroll
    def 'Delete item with id [#id] for collector [#collector] throws an exception when collector [#collector] does not own an item with id [#id]'() {
        given:
        def item = new Item(id: id, name: name)

        when:
        itemService.delete(item, collector)

        then:
        1 * itemRepository.existsByIdAndCollector(id, collector) >> Boolean.FALSE
        0 * itemRepository.update(_)

        and:
        thrown(ItemService.OwnershipException)

        where:
        id | name           | collector
        1  | 'an item'      | 'bob'
        2  | 'another item' | 'bob'
    }
}
