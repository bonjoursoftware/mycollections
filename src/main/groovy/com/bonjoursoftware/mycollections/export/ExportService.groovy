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
package com.bonjoursoftware.mycollections.export

import com.bonjoursoftware.mycollections.collector.CollectorRepository
import com.bonjoursoftware.mycollections.item.ItemRepository
import groovy.transform.CompileStatic

import javax.inject.Inject
import javax.inject.Singleton
import java.time.LocalDateTime

@CompileStatic
@Singleton
class ExportService {

    @Inject
    private CollectorRepository collectorRepository

    @Inject
    private ItemRepository itemRepository

    Export run() {
        new Export(
                date: LocalDateTime.now(),
                collectors: collectorRepository.findAll().collect() {
                    new EnrichedCollector(
                            collector: it,
                            items: itemRepository.findByCollector(it.username))
                }
        )
    }
}
