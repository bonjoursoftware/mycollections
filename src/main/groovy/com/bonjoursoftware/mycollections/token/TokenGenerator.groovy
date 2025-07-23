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
package com.bonjoursoftware.mycollections.token

import groovy.transform.CompileStatic
import org.springframework.security.crypto.keygen.Base64StringKeyGenerator

import jakarta.inject.Singleton

import static org.springframework.security.crypto.bcrypt.BCrypt.gensalt
import static org.springframework.security.crypto.bcrypt.BCrypt.hashpw

@CompileStatic
@Singleton
class TokenGenerator implements TokenService {

    private static final int KEY_LENGTH = 54
    private static final int SALT_LOG_ROUNDS = 10

    private Base64StringKeyGenerator generator = new Base64StringKeyGenerator(KEY_LENGTH)

    @Override
    Token generate() {
        generator.generateKey().with { secret ->
            new Token(secret: secret, hash: hashpw(secret, gensalt(SALT_LOG_ROUNDS)))
        }
    }
}
