package com.camp.promotion.mq;

import com.alibaba.fastjson2.JSON;
import com.camp.promotion.entity.HOrder;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.client.exception.MQBrokerException;
import org.apache.rocketmq.client.exception.MQClientException;
import org.apache.rocketmq.client.producer.DefaultMQProducer;
import org.apache.rocketmq.client.producer.SendResult;
import org.apache.rocketmq.client.producer.SendStatus;
import org.apache.rocketmq.common.message.Message;
import org.apache.rocketmq.remoting.exception.RemotingException;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.nio.charset.StandardCharsets;

@Slf4j
@Component
public class OrderProducer {

    private DefaultMQProducer producer;

    @Value("${rocketmq.nameserver.addr}")
    private String namesrvAddr;

    @Value("${rocketmq.topic.name}")
    private String topicName;

    @PostConstruct
    public void init() throws MQClientException {

        //初始化Producer
        producer = new DefaultMQProducer();
        producer.setProducerGroup("create_order_producer_group");
        producer.setNamesrvAddr(namesrvAddr);
        producer.setRetryTimesWhenSendFailed(3);
        producer.setSendMsgTimeout(1000);
        producer.start();
    }

    public void send() {

        try {
            for (int i = 0; i < 100; i++) {
                String body = "TEST_ORDER_" + i;
                Message message = new Message(topicName, "create_order" + i, body.getBytes(StandardCharsets.UTF_8));

                //发送消息，需要关注发送结果，并捕获失败等异常。
                SendResult sendResult = producer.send(message);
            }

        } catch (MQBrokerException | MQClientException | RemotingException | InterruptedException e) {
            e.printStackTrace();
        }
    }

    public boolean send(HOrder order) {
        byte[] body = JSON.toJSONBytes(order);
        Message message = new Message(topicName, "create_order" + order.getId(), body);
        SendResult sendResult;
        try {
            sendResult = producer.send(message);
        } catch (MQClientException | RemotingException | MQBrokerException | InterruptedException e) {
            log.error("seng msg to broker error, e = ", e);
            return false;
        }
        return sendResult.getSendStatus() == SendStatus.SEND_OK;
    }
}
