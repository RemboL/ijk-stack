FROM maven:3.8.5-eclipse-temurin-17-alpine AS build
COPY src /home/app/src
COPY pom.xml /home/app
RUN mvn -f /home/app/pom.xml clean package

#
# Package stage
#
FROM eclipse-temurin:17.0.1_12-jdk-alpine
COPY --from=build /home/app/target/commerce-0.0.1-SNAPSHOT.jar /usr/local/lib/commerce.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/usr/local/lib/commerce.jar"]
