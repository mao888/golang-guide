package com.camp.promotion.lock;

import lombok.extern.slf4j.Slf4j;

import java.io.Closeable;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;

@Slf4j
public abstract class AbstractDistributedLock implements Lock, Closeable {

    protected final AtomicBoolean isLock = new AtomicBoolean(false);

    /**
     * 锁轮询间隔,单位毫秒
     */
    protected long intervalTime = 10L;


    /**
     * 锁最大存在时间,单位毫秒
     */
    protected long survivalMillis = 600000L;


    @Override
    public void lock() {
        while (survivalMillis > 0) {
            try {
                if (tryLock()) {
                    return;
                }
                Thread.sleep(intervalTime);
                survivalMillis -= intervalTime;
            } catch (InterruptedException e) {
                log.error("lock thread interrupt", e);
                Thread.currentThread().interrupt();
            }
        }

    }

    @Override
    public void lockInterruptibly() throws InterruptedException {
        lock();
    }

    @Override
    public boolean tryLock() {
        return false;
    }

    @Override
    public boolean tryLock(long time, TimeUnit unit) {

        this.survivalMillis = TimeUnit.MILLISECONDS.convert(time, unit);
        while (survivalMillis > 0) {
            try {
                if (tryLock()) {
                    return true;
                }
                Thread.sleep(this.intervalTime);
                survivalMillis -= intervalTime;
            } catch (InterruptedException e) {
                log.error("lock thread interrupt, ", e);
                Thread.currentThread().interrupt();
            }
        }

        return false;
    }

    @Override
    public void unlock() {

    }

    @Override
    public Condition newCondition() {
        throw new UnsupportedOperationException();
    }

    @Override
    public void close() throws IOException {
        if (this.isLock.get()) {
            unlock();
        }
    }
}
