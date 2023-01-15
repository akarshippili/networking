package com.akarsh.spring.rabbitmq.config;

import org.springframework.amqp.core.*;
import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.annotation.PostConstruct;

@Configuration
public class QueuesConfig {

    @Autowired
    private RabbitAdmin admin;

    @Bean
    public Queue queue(){
        return new Queue("myQueue");
    }

    @Bean
    public DirectExchange exchange(){
        return new DirectExchange("myExchange");
    }

    @Bean
    public Binding binding(Queue queue, DirectExchange exchange){
        return BindingBuilder.bind(queue).to(exchange).with("myRoutingKey");
    }

    @PostConstruct
    public void declare(){
        admin.declareQueue(queue());
        admin.declareExchange(exchange());
        admin.declareBinding(binding(queue(), exchange()));
    }

}
