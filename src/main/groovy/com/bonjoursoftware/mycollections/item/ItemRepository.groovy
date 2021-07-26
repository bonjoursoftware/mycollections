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
import com.bonjoursoftware.mycollections.mongo.MongoRepository
import com.bonjoursoftware.mycollections.mongo.StringToObjectIdConverter
import com.mongodb.client.MongoCollection
import groovy.transform.CompileStatic
import org.bson.conversions.Bson

import javax.inject.Singleton

import static com.mongodb.client.model.Filters.eq
import static com.mongodb.client.model.Filters.regex

@CompileStatic
@Singleton
class ItemRepository implements MongoRepository {

    private static final String ID_FIELD = '_id'
    private static final String NAME_FIELD = 'name'
    private static final String TEXT_SEARCH_OPTIONS = 'i'

    private StringToObjectIdConverter oidConverter = new StringToObjectIdConverter()

    List<Item> findByCollector(Collector collector) {
        collection(collector).find().asList()
    }

    Item findByIdAndCollector(String id, Collector collector) {
        collection(collector).find(byId(id)).first()
    }

    List<Item> findByNameAndCollector(String name, Collector collector) {
        collection(collector).find(byName(name)).asList()
    }

    void upsert(Item item, Collector collector) {
        item.id ? replace(item, collector) : create(item, collector)
    }

    void delete(Item item, Collector collector) {
        collection(collector).deleteOne(byId(item))
    }

    private void create(Item item, Collector collector) {
        collection(collector).insertOne(item)
    }

    private void replace(Item item, Collector collector) {
        collection(collector).replaceOne(byId(item), item)
    }

    private Bson byId(String id) {
        eq(ID_FIELD, oidConverter.convert(id))
    }

    private Bson byId(Item item) {
        eq(ID_FIELD, item.id)
    }

    private Bson byName(String name) {
        regex(NAME_FIELD, name, TEXT_SEARCH_OPTIONS)
    }

    private MongoCollection<Item> collection(Collector collector) {
        db().getCollection(collector.username, Item)
    }
}
