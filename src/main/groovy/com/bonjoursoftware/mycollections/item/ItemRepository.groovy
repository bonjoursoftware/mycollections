package com.bonjoursoftware.mycollections.item

import io.micronaut.data.jdbc.annotation.JdbcRepository
import io.micronaut.data.model.query.builder.sql.Dialect
import io.micronaut.data.repository.CrudRepository

@JdbcRepository(dialect = Dialect.POSTGRES)
interface ItemRepository extends CrudRepository<Item, Long> {
    List<Item> findByName(String name)
    List<Item> findByCollector(String collector)
}