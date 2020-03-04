package com.bonjoursoftware.mycollections.item

import groovy.transform.TupleConstructor

import javax.persistence.Entity
import javax.persistence.GeneratedValue
import javax.persistence.Id

@Entity
@TupleConstructor(excludes = 'id')
class Item {
    @Id
    @GeneratedValue
    Long id
    String collector
    String name
}
