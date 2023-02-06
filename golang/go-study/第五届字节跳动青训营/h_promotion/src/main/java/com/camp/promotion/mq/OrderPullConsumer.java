package com.camp.promotion.mq;

import com.alibaba.fastjson2.JSON;
import com.camp.promotion.common.BizException;
import com.camp.promotion.common.ResponseEnum;
import com.camp.promotion.entity.HOrder;
import com.camp.promotion.service.HOrderService;
import com.camp.promotion.service.HPromoProductService;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.collections4.CollectionUtils;
import org.apache.rocketmq.client.consumer.DefaultLitePullConsumer;
import org.apache.rocketmq.common.consumer.ConsumeFromWhere;
import org.apache.rocketmq.common.message.MessageExt;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.util.ArrayList;
import java.util.List;
import java.util.Random;
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;

@Slf4j
@Component
public class OrderPullConsumer {

    private DefaultLitePullConsumer consumer;

    @Value("${rocketmq.topic.name}")
    private String topicName;

    @Value("${rocketmq.nameserver.addr}")
    private String namesrvAddr;

    @Resource
    private HOrderService hOrderService;

    @Resource
    private HPromoProductService hPromoProductService;

    private final ArrayBlockingQueue<HOrder> deadQueue = new ArrayBlockingQueue<>(1000);

    private ExecutorService executorService;

    private final Random random = new Random();

    private final AtomicBoolean isRunning = new AtomicBoolean(false);

    @PostConstruct
    public void init() throws Exception {
        consumer = new DefaultLitePullConsumer("create_order_consumer_group");
        consumer.setNamesrvAddr(namesrvAddr);
        consumer.subscribe(topicName, "*");
        consumer.setPullBatchSize(100);
        consumer.setConsumeFromWhere(ConsumeFromWhere.CONSUME_FROM_FIRST_OFFSET);
        consumer.start();

        executorService = new ThreadPoolExecutor(10, 10, 120, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(2000), new ThreadFactory() {
            AtomicInteger index = new AtomicInteger();

            @Override
            public Thread newThread(Runnable r) {
                Thread t = new Thread(r);
                t.setName("thread-process-order-" + index.getAndIncrement());
                t.setPriority(Thread.NORM_PRIORITY);
                return t;
            }
        });

        new Thread(() -> {
            try {
                this.isRunning.set(true);
                this.doConsumer();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }, "thread-consumer-start").start();

        Runtime.getRuntime().addShutdownHook(new Thread(this::stop));
    }

    public void doConsumer() throws InterruptedException {
        while (isRunning.get()) {
            List<MessageExt> messageExts = consumer.poll();
            if (CollectionUtils.isEmpty(messageExts)) {
                continue;
            }
            log.info("start consumer msg size = {}", messageExts.size());

            executorService.execute(() -> {
                try {
                    processOrder(messageExts);
                } catch (Exception e) {
                    log.error("process order fail, e = ", e);
                }
            });

            int time = random.nextInt(100);
            Thread.sleep(time);
        }
    }

    @Transactional(rollbackFor = Exception.class)
    public void processOrder(List<MessageExt> messageExts) {
        int decreaseCount = 0;
        List<HOrder> orders = new ArrayList<>(messageExts.size());

        for (MessageExt messageExt : messageExts) {
            byte[] body = messageExt.getBody();
            HOrder order = JSON.parseObject(body, HOrder.class);
            orders.add(order);

            decreaseCount += order.getQuantity();
        }

        if (CollectionUtils.isEmpty(orders)) {
            return;
        }

        int updateStockRes = this.hPromoProductService.decreaseStock(orders.get(0).getActivityId(), orders.get(0).getSkuId(), decreaseCount);
        if (updateStockRes < 1) {
            log.error("process order fail, size = {}, data = {}", orders.size(), JSON.toJSONString(orders));
            this.deadQueue.addAll(orders);
            throw new BizException(ResponseEnum.CONSUMER_FAIL);
        }

        int insertRes = this.hOrderService.batchInsert(orders);
        if (insertRes < orders.size()) {
            log.error("process order fail, size = {}, data = {}", orders.size(), JSON.toJSONString(orders));
            this.deadQueue.addAll(orders);
            throw new BizException(ResponseEnum.CONSUMER_FAIL);
        }
    }

    public void stop() {
        log.info("stop consumer!");
        this.isRunning.set(false);
        this.executorService.shutdown();
    }
}
