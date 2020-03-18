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

import com.bonjoursoftware.mycollections.mongo.MongoRepository
import com.mongodb.client.MongoCollection
import groovy.transform.CompileStatic

import javax.inject.Singleton

@CompileStatic
@Singleton
class TagRepository implements MongoRepository {

    List<String> findByCollector(String collector) {
        collection(collector).distinct(TAGS_FIELD, String).asList()
    }

    private MongoCollection collection(String collector) {
        db().getCollection(collector, String)
    }
}
