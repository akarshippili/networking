package com.akarsh.spring.rabbitmq.listeners;

import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

@Component
@Slf4j
public class RabbitMQMessageListener {
    @RabbitListener(queues = "${rabbitmq.queue.myQueue}")
    public void processMessage(String message) {
        log.info("Received RabbitMQ message {}", message);
    }
}
