package com.bonjoursoftware.mycollections.tag

import com.bonjoursoftware.mycollections.item.Item
import com.fasterxml.jackson.annotation.JsonIgnore
import groovy.transform.EqualsAndHashCode

import javax.persistence.Entity
import javax.persistence.GeneratedValue
import javax.persistence.GenerationType
import javax.persistence.Id
import javax.persistence.ManyToMany
import javax.persistence.SequenceGenerator

@Entity
@EqualsAndHashCode(includes = 'name')
class Tag {
    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = 'tag_id_seq')
    @SequenceGenerator(name = 'tag_id_seq', allocationSize = 1)
    Long id

    String name

    @JsonIgnore
    @ManyToMany(mappedBy = 'tags')
    Set<Item> items
}
