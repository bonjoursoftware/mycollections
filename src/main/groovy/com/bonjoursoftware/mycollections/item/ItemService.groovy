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

import groovy.transform.CompileStatic
import groovy.transform.InheritConstructors

import javax.inject.Inject
import javax.inject.Singleton

@CompileStatic
@Singleton
class ItemService {

    @Inject
    private ItemRepository itemRepository

    List<ItemDTO> findByCollector(String collector) {
        itemRepository.findByCollector(collector)
    }

    void create(ItemDTO item, String collector) {
        itemRepository.save(new Item(name: item.name, collector: collector))
    }

    void update(ItemDTO item, String collector) {
        verifyOwnershipAndDo(item, collector) {
            itemRepository.update(new Item(id: item.id, name: item.name, collector: collector))
        }
    }

    void delete(ItemDTO item, String collector) {
        verifyOwnershipAndDo(item, collector) {
            itemRepository.deleteById(item.id)
        }
    }

    private void verifyOwnershipAndDo(ItemDTO item, String collector, Closure action) {
        if (itemRepository.existsByIdAndCollector(item.id, collector)) {
            action()
        } else {
            throw new OwnershipException("Collector [${collector}] does not own an item with id [${item.id}]")
        }
    }

    @InheritConstructors
    static class OwnershipException extends RuntimeException {
    }
}
