package com.akarsh.spring.rabbitmq.controller;

import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
public class Hello {

    @Autowired
    private RabbitTemplate template;

    @PostMapping(path = "/")
    public String hello(@RequestParam String message){
        template.convertAndSend("myExchange", "myRoutingKey", message);
        return "Hello World";
    }

}
