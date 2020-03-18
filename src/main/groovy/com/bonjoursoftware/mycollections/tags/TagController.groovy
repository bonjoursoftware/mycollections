package com.bonjoursoftware.mycollections.tags

import groovy.transform.CompileStatic
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get
import io.micronaut.security.annotation.Secured
import io.micronaut.security.authentication.Authentication
import io.micronaut.security.rules.SecurityRule

import javax.inject.Inject

import static com.bonjoursoftware.mycollections.collector.Collector.getCollector

@CompileStatic
@Secured(SecurityRule.IS_AUTHENTICATED)
@Controller('/api/v1/tag')
class TagController {

    @Inject
    private TagRepository tagRepository

    @Get
    List<Tag> findByCollector(Authentication authentication) {
        tagRepository.findByCollector(getCollector(authentication))
    }
}
