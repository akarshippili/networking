package com.akarsh.spring.rabbitmq.config;

import org.springframework.amqp.core.*;
import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.annotation.PostConstruct;

@Configuration
public class AutoCreateTest {

    @Value("${rabbitmq.auto-created-queue.queue-name}")
    private String queueName;

    @Value("${rabbitmq.auto-created-queue.exchange-name}")
    private String exchangeName;

    @Value("${rabbitmq.auto-created-queue.binding-key}")
    private String bindingKey;

    @Autowired
    private RabbitAdmin admin;

    @Bean
    public Queue autoCreatedQueue() {
        return new Queue(queueName);
    }

    @Bean
    public TopicExchange autoCreatedExchange() {
        return new TopicExchange(exchangeName);
    }

    @Bean
    public Binding autoCreatedBinding(Queue queue, TopicExchange exchange){
        return BindingBuilder.bind(queue).to(exchange).with(bindingKey);
    }

    @PostConstruct
    public void init() {
        admin.declareQueue(autoCreatedQueue());
        admin.declareExchange(autoCreatedExchange());
        admin.declareBinding(autoCreatedBinding(autoCreatedQueue(), autoCreatedExchange()));
    }

}
