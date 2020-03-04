package com.bonjoursoftware.mycollections.item

import groovy.transform.CompileStatic
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get
import io.micronaut.security.annotation.Secured
import io.micronaut.security.authentication.Authentication
import io.micronaut.security.rules.SecurityRule

import javax.inject.Inject

@CompileStatic
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller('/item')
class ItemController {

    @Inject
    ItemRepository itemRepository

    @Get
    List<Item> findByCollector(Authentication authentication) {
        itemRepository.findByCollector(authentication.getAttributes().get('email') as String)
    }
}
