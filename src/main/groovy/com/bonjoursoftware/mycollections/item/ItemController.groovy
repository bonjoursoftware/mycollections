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
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Post
import io.micronaut.http.annotation.QueryValue
import io.micronaut.security.annotation.Secured
import io.micronaut.security.authentication.Authentication
import io.micronaut.security.rules.SecurityRule

import javax.inject.Inject

import static com.bonjoursoftware.mycollections.collector.CollectorAuthenticationExtractor.getUsername

@CompileStatic
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller('/api/v1/item')
class ItemController {

    @Inject
    private ItemRepository itemRepository

    @Get
    List<Item> findByCollector(Authentication authentication) {
        itemRepository.findByCollector(getUsername(authentication))
    }

    @Get('/{id}')
    Item findByIdAndCollector(@QueryValue('id') String id, Authentication authentication) {
        itemRepository.findByIdAndCollector(id, getUsername(authentication))
    }

    @Post
    void upsert(Item item, Authentication authentication) {
        item.tap { tags = tags?.collect { it.trim() } }
        itemRepository.upsert(item, getUsername(authentication))
    }

    @Delete
    void delete(Item item, Authentication authentication) {
        itemRepository.delete(item, getUsername(authentication))
    }
}
