# RunCloud Assesment

This project is written for RunCloud's Assesment for Junior Level Position.

## Task

Given a situation that you have Windows and run VirtualBox Virtual Machine (VM). The VM will be running Linux (Ubuntu, Centos, etc) and you run a PHP application inside the Windows to receive data from the Linux VM. Inside the Linux VM, you need to send the available disk space of the Linux VM to the PHP application.


To obtain the available disk space, you need to write a piece of code using Go (or any compiled language of your choice) to send the available disk space to a message broker (using RabbitMQ, Apache Kafka or anything you find suitable). Inside the PHP application, you need to subscribe to the message broker to receive the data that was sent by the application that you have written inside the Linux.

### What you need to provide

1. A PHP application that is running inside the host which subscribed to a messaging broker.

2. A config of a messaging broker.

3. The source code of the application running inside the Linux VM to obtain the available disk space.

## Authors

* **Ashraf Kamarudin**

## License

This project is not licensed under anything or anyone and is free to be used

