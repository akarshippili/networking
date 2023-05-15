package com.akarsh.spring.rabbitmq.config;

import org.springframework.amqp.core.*;
import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.annotation.PostConstruct;

@Configuration
public class QueuesConfig {

    @Value("${rabbitmq.queue.myQueue}")
    private String queueName;

    @Value("${rabbitmq.queue.myExchange}")
    private String exchangeName;

    @Value("${rabbitmq.queue.myRoutingKey}")
    private String routingKey;

    @Autowired
    private RabbitAdmin admin;

    @Bean
    public Queue queue(){
        return new Queue(queueName);
    }

    @Bean
    public DirectExchange exchange(){
        return new DirectExchange(exchangeName);
    }

    @Bean
    public Binding binding(Queue queue, DirectExchange exchange){
        return BindingBuilder.bind(queue).to(exchange).with(routingKey);
    }

    @PostConstruct
    public void declare(){
        admin.declareQueue(queue());
        admin.declareExchange(exchange());
        admin.declareBinding(binding(queue(), exchange()));
    }

}
