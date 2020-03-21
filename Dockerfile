FROM gradle:6.2.2-jdk13 AS builder
WORKDIR /home/gradle/project
COPY settings.gradle ./
COPY gradle.properties ./
COPY build.gradle ./
COPY src/ ./src/
RUN gradle shadowJar --no-daemon --info --gradle-user-home ./

FROM adoptopenjdk/openjdk13-openj9:jre-13.0.2_8_openj9-0.18.0-alpine
COPY --from=builder /home/gradle/project/build/libs/*.jar ./mycollections.jar
EXPOSE 8080
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx256m", "-XX:+IdleTuningGcOnIdle", "-Xtune:virtualized", "-jar", "mycollections.jar"]
