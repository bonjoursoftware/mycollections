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

import com.bonjoursoftware.mycollections.export.ExportService
import com.bonjoursoftware.mycollections.notification.NotificationService
import com.fasterxml.jackson.databind.ObjectMapper
import groovy.transform.CompileStatic
import io.micronaut.scheduling.annotation.Scheduled

import jakarta.inject.Inject
import jakarta.inject.Singleton

import static com.fasterxml.jackson.annotation.JsonInclude.Include.NON_NULL

@CompileStatic
@Singleton
class BackupJob {

    private static final String BACKUP_NOTIFICATION_TITLE = 'Data Backup'

    @Inject
    private ExportService exportService

    @Inject
    private NotificationService notificationService

    @Scheduled(cron = '0 0 23 ? * *')
    void backup() {
        notificationService.notify(BACKUP_NOTIFICATION_TITLE, buildBackup())
    }

    private String buildBackup() {
        new ObjectMapper()
                .setSerializationInclusion(NON_NULL)
                .writeValueAsString(exportService.run())
    }
}
