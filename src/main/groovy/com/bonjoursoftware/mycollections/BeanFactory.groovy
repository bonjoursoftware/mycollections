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
package com.bonjoursoftware.mycollections

import com.mongodb.ConnectionString
import com.mongodb.MongoClientSettings
import com.mongodb.client.MongoClient
import com.mongodb.client.MongoClients
import groovy.transform.CompileStatic
import io.micronaut.context.annotation.Factory
import io.micronaut.context.annotation.Property
import org.bson.codecs.configuration.CodecRegistries
import org.bson.codecs.pojo.PojoCodecProvider

import javax.inject.Singleton

@Factory
@CompileStatic
class BeanFactory {

    @Singleton
    MongoClient mongoClient(
            @Property(name = 'mongo.user') String user,
            @Property(name = 'mongo.password') String password,
            @Property(name = 'mongo.host') String host
    ) {
        MongoClients.create(MongoClientSettings.builder()
                .applyConnectionString(new ConnectionString("mongodb+srv://${user}:${password}@${host}/?retryWrites=true&w=majority"))
                .codecRegistry(CodecRegistries.fromRegistries(MongoClientSettings.getDefaultCodecRegistry(), CodecRegistries.fromProviders(PojoCodecProvider.builder().automatic(true).build())))
                .build())
    }
}
