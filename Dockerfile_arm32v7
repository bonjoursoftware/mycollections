FROM arm32v7/adoptopenjdk:16.0.1_9-jre-hotspot
WORKDIR /home/mycollections
COPY build/libs/*.jar ./mycollections.jar
EXPOSE 8443
CMD ["java", "-Dcom.sun.management.jmxremote", "-Xmx256m", "--add-exports", "java.base/sun.security.x509=ALL-UNNAMED", "-jar", "mycollections.jar"]
