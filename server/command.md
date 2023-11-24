# kafka servers

## Kafka topic

```shell
# list topics
kafka-topics --bootstrap-server=localhost:9092 --list
# create topic
kafka-topics --bootstrap-server=localhost:9092 --topic=OpenAccountEvent --create
# describe topic
kafka-topics --bootstrap-server=localhost:9092 --topic=OpenAccountEvent --describe
```

## Consumer

```shell
# create consumer to subscribe the topic
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=OpenAccountEvent
# create consumer to subscribe the topic with group
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=OpenAccountEvent --group=accountConsumer
# create consumer to subscribe multi topic with group
kafka-console-consumer --bootstrap-server=localhost:9092 --include="OpenAccountEvent|DepositFundEvent|WithdrawFundEvent|CloseAccountEvent" --group=accountConsumer
kafka-console-consumer --bootstrap-server=localhost:9092 --include="OpenAccountEvent|DepositFundEvent|WithdrawFundEvent|CloseAccountEvent" --group=log
# list consumer group
kafka-consumer-groups --bootstrap-server=localhost:9092 --list
```

## Producer

```shell
# create producer to subscribe the topic
kafka-console-producer --bootstrap-server=localhost:9092 --topic=OpenAccountEvent
kafka-console-producer --bootstrap-server=localhost:9092 --topic=DepositFundEvent
```