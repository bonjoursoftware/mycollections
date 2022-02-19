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

import com.bonjoursoftware.mycollections.collector.Collector
import com.bonjoursoftware.mycollections.collector.CollectorRepository
import com.bonjoursoftware.mycollections.collector.CollectorRoles
import com.bonjoursoftware.mycollections.item.Item
import com.bonjoursoftware.mycollections.item.ItemRepository
import spock.lang.Specification
import spock.lang.Subject

class ExportServiceTest extends Specification {

    private CollectorRepository collectorRepository
    private ItemRepository itemRepository

    @Subject
    private ExportService exportService

    void setup() {
        collectorRepository = Mock()
        itemRepository = Mock()
        exportService = new ExportService(collectorRepository: collectorRepository, itemRepository: itemRepository)
    }

    def 'The export service collates all collectors with write role and their items'() {
        given:
        def collectors = [
                new Collector(username: 'collector1', roles: [CollectorRoles.READ, CollectorRoles.WRITE]),
                new Collector(username: 'collector2', roles: [CollectorRoles.READ]),
                new Collector(username: 'collector3', roles: [CollectorRoles.READ, CollectorRoles.WRITE]),
        ]

        and:
        def itemsCollector1 = [
                new Item(name: 'first item collector1'),
                new Item(name: 'second item collector1'),
        ]

        and:
        def itemsCollector3 = [
                new Item(name: 'first item collector3'),
                new Item(name: 'second item collector3'),
        ]

        when:
        Export export = exportService.run()

        then:
        1 * collectorRepository.findAll() >> collectors
        1 * itemRepository.findByCollector(collectors.first()) >> itemsCollector1
        1 * itemRepository.findByCollector(collectors.last()) >> itemsCollector3

        and:
        export.collectors.size() == 2
        export.collectors.contains(new EnrichedCollector(collector: collectors.first(), items: itemsCollector1))
        export.collectors.contains(new EnrichedCollector(collector: collectors.last(), items: itemsCollector3))
    }
}
