FROM gradle:6.7.0-jdk14 AS builder
WORKDIR /home/gradle/project
COPY settings.gradle ./
COPY gradle.properties ./
COPY build.gradle ./
COPY src/ ./src/
RUN gradle test shadowJar --no-daemon --info --gradle-user-home ./

FROM adoptopenjdk:14-jre-hotspot
COPY --from=builder /home/gradle/project/build/libs/*.jar ./mycollections.jar
EXPOSE 8443
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx256m", "-jar", "mycollections.jar"]
