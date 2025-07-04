FROM docker.io/arm64v8/eclipse-temurin:17-jre
WORKDIR /home/mycollections
COPY build/libs/*.jar ./mycollections.jar
EXPOSE 8443
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx256m", "--add-exports", "java.base/sun.security.x509=ALL-UNNAMED", "-jar", "mycollections.jar"]
