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
