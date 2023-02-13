package com.camp.promotion.mq;

import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.client.consumer.DefaultLitePullConsumer;
import org.apache.rocketmq.client.consumer.DefaultMQPushConsumer;
import org.apache.rocketmq.client.consumer.listener.ConsumeConcurrentlyContext;
import org.apache.rocketmq.client.consumer.listener.ConsumeConcurrentlyStatus;
import org.apache.rocketmq.client.consumer.listener.MessageListenerConcurrently;
import org.apache.rocketmq.common.consumer.ConsumeFromWhere;
import org.apache.rocketmq.common.message.MessageExt;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.List;

@Slf4j
@Component
public class OrderPushConsumer {

    private DefaultMQPushConsumer consumer;


    @Value("${rocketmq.topic.name}")
    private String topicName;

    @Value("${rocketmq.nameserver.addr}")
    private String namesrvAddr;

//    @PostConstruct
    public void init() throws Exception {
        consumer = new DefaultMQPushConsumer("create_order_consumer_group2");
        consumer.setNamesrvAddr(namesrvAddr);
        consumer.subscribe(topicName,"*");
        consumer.setConsumeMessageBatchMaxSize(100);
        consumer.setConsumeFromWhere(ConsumeFromWhere.CONSUME_FROM_FIRST_OFFSET);

        consumer.registerMessageListener((MessageListenerConcurrently) (msgs, context) -> {
            log.info("consume msg list size = {}", msgs.size());
            return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
        });
        consumer.start();

    }
}
