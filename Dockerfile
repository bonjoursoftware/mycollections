FROM gradle:6.6.0-jdk14 AS builder
WORKDIR /home/gradle/project
COPY settings.gradle ./
COPY gradle.properties ./
COPY build.gradle ./
COPY src/ ./src/
RUN gradle shadowJar --no-daemon --info --gradle-user-home ./

FROM adoptopenjdk/openjdk14-openj9:jre-14.0.2_12_openj9-0.21.0-alpine
COPY --from=builder /home/gradle/project/build/libs/*.jar ./mycollections.jar
EXPOSE 8443
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx256m", "-XX:+IdleTuningGcOnIdle", "-Xtune:virtualized", "-jar", "mycollections.jar"]