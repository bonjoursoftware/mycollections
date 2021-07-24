FROM gradle:7.1.1-jdk16 AS builder
WORKDIR /home/gradle/project
COPY settings.gradle ./
COPY gradle.properties ./
COPY build.gradle ./
COPY src/ ./src/
RUN gradle test shadowJar --no-daemon --info --gradle-user-home ./

FROM adoptopenjdk:16-jre-hotspot
COPY --from=builder /home/gradle/project/build/libs/*.jar ./mycollections.jar
EXPOSE 8443
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx256m", "--add-exports", "java.base/sun.security.x509=ALL-UNNAMED", "-jar", "mycollections.jar"]
