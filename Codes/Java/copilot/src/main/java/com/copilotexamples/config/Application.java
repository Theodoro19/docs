package com.copilotexamples.config;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@EnableJpaRepositories("com.copilotexamples.persistence.repo.*")
@EntityScan("com.copilotexamples.persistence.model.*")
@SpringBootApplication(scanBasePackages = {"com.copilotexamples.controller.*", "com.copilotexamples.service.*"})
public class Application {

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
