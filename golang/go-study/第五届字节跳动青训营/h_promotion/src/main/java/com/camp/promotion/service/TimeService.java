package com.camp.promotion.service;


import java.util.concurrent.ScheduledThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

public class TimeService {
    private volatile long now;

    private TimeService() {
        this.now = System.currentTimeMillis();
        scheduleTick();
    }

    private void scheduleTick() {
        new ScheduledThreadPoolExecutor(1, runnable -> {
            Thread thread = new Thread(runnable, "current-time-millis");
            thread.setDaemon(true);
            return thread;
        }).scheduleAtFixedRate(() -> now = System.currentTimeMillis(), 1, 1, TimeUnit.MILLISECONDS);
    }

    public long now() {
        return now;
    }

    public static TimeService getInstance() {
        return TimeService.SingletonHolder.INSTANCE;
    }

    private static class SingletonHolder {
        private static final TimeService INSTANCE = new TimeService();
    }
}
