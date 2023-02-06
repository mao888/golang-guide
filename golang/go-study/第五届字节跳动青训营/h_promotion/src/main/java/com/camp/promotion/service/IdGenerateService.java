package com.camp.promotion.service;

import com.google.common.base.Preconditions;
import lombok.Getter;
import lombok.Setter;
import lombok.SneakyThrows;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.util.Calendar;
import java.util.Properties;

@Service
public class IdGenerateService {

    public static final long EPOCH;

    private static final long SEQUENCE_BITS = 12L;

    private static final long WORKER_ID_BITS = 10L;
    //mask的作用就是将sequence所代表的最后12位全部置为1，方便后面做位于运算
    private static final long SEQUENCE_MASK = (1 << SEQUENCE_BITS) - 1;

    private static final long WORKER_ID_LEFT_SHIFT_BITS = SEQUENCE_BITS;

    private static final long TIMESTAMP_LEFT_SHIFT_BITS = WORKER_ID_LEFT_SHIFT_BITS + WORKER_ID_BITS;

    private static final long WORKER_ID_MAX_VALUE = 1L << WORKER_ID_BITS;

    private static final long WORKER_ID = 0;

    private static final int MAX_TOLERATE_TIME_DIFFERENCE_MILLISECONDS = 10;

    @Setter
    private static TimeService timeService = TimeService.getInstance();

    @Getter
    @Setter
    private Properties properties = new Properties();

    @Value("${worker.id}")
    private long workId;

    private byte sequenceOffset;
    /**
     * 毫秒内自增序列(0~4095) 既最后的12个bit位，2^12-1
     */
    private long sequence;

    private long lastMilliseconds;

    static {
        Calendar calendar = Calendar.getInstance();
        calendar.set(2022, Calendar.NOVEMBER, 1);
        calendar.set(Calendar.HOUR_OF_DAY, 0);
        calendar.set(Calendar.MINUTE, 0);
        calendar.set(Calendar.SECOND, 0);
        calendar.set(Calendar.MILLISECOND, 0);
        EPOCH = calendar.getTimeInMillis();
    }

    public String getType() {
        return "SNOWFLAKE";
    }

    public synchronized long generateId() {

        long currentMilliseconds = timeService.now();
        //处理时间回溯问题
        if (waitTolerateTimeDifferenceIfNeed(currentMilliseconds)) {
            currentMilliseconds = timeService.now();
        }
        //如果是在同一毫秒中调用此方法，
        if (lastMilliseconds == currentMilliseconds) {
            //这里是一个位运算，因为sequence最多只允许占用最低的12位，一旦+1的步骤达到4096，
            //达到了13位，既2^13,通过和SEQUENCE_MASK做位与运算，重置为0
            if (0L == (sequence = (sequence + 1) & SEQUENCE_MASK)) {
                //一旦重置为0后，说明这一毫秒能生成的key已达到上限，则进入循环等待下一个毫秒。
                // 因为仅仅是一个毫秒的差距，所以没有使用sleep，直接while循环节省cpu调用
                currentMilliseconds = waitUntilNextTime(currentMilliseconds);
            }
        } else {
            //如果不是在同一毫秒中并发调用此方法
            vibrateSequenceOffset();
            sequence = sequenceOffset;
        }
        lastMilliseconds = currentMilliseconds;
        /*
         * (currentMilliseconds - EPOCH) << TIMESTAMP_LEFT_SHIFT_BITS) 时间毫秒数往右移动22位，避开workId和sequence，放到高位的41位
         * (getWorkerId() << WORKER_ID_LEFT_SHIFT_BITS) 将workId左移到随后的10位上
         * 最后是sequence占据低位的12位
         *
         * 最后用或运算将三个部分组合到一个long中,返回最终的结果
         */
        return ((currentMilliseconds - EPOCH) << TIMESTAMP_LEFT_SHIFT_BITS) | (getWorkerId() << WORKER_ID_LEFT_SHIFT_BITS) | sequence;
    }

    /**
     * 如果当前时间没有回溯过，直接返回false，不等待
     * 如果当前时间回溯过，如果在允许的时间范围内，就sleep，返回true，表示等待过
     *                  如果超过了允许的时间范围，直接抛出异常。
     * @param currentMilliseconds
     * @return
     */
    @SneakyThrows
    private boolean waitTolerateTimeDifferenceIfNeed(final long currentMilliseconds) {
        if (lastMilliseconds <= currentMilliseconds) {
            return false;
        }
        //回溯的时间毫秒数
        long timeDifferenceMilliseconds = lastMilliseconds - currentMilliseconds;

        Preconditions.checkState(timeDifferenceMilliseconds < getMaxTolerateTimeDifferenceMilliseconds(),
                "Clock is moving backwards, last time is %d milliseconds, current time is %d milliseconds", lastMilliseconds, currentMilliseconds);
        Thread.sleep(timeDifferenceMilliseconds);
        return true;
    }

    private long getWorkerId() {
        long result = this.workId;
        Preconditions.checkArgument(result >= 0L && result < WORKER_ID_MAX_VALUE);
        return result;
    }

    private int getMaxTolerateTimeDifferenceMilliseconds() {
        return Integer.parseInt(properties.getProperty("max.tolerate.time.difference.milliseconds", String.valueOf(MAX_TOLERATE_TIME_DIFFERENCE_MILLISECONDS)));
    }

    private long waitUntilNextTime(final long lastTime) {
        long result = timeService.now();
        while (result <= lastTime) {
            result = timeService.now();
        }
        return result;
    }

    /**
     * byte的初始值0000 0000
     * ~取反后就是-1，二进制1111 1111
     * 和1做位于运算,0000 0001，可以发现，甭管你是多少，最后都只会留下最低位，1或者0
     * sequenceOffset最开始是0，结果就是1
     * sequenceOffset变成1后，一通操作后结果就是0
     *
     * 所以这个sequenceOffset每次调用循环变成0，1.
     * 这个原理是为了让在分布式环境中低QPS的时候，末尾不总是0，就不会全是偶数。造成切片堆积
     */
    private void vibrateSequenceOffset() {
        sequenceOffset = (byte) (~sequenceOffset & 1);
    }

    public static void main(String[] args) {
        IdGenerateService idService = new IdGenerateService();
        long id = idService.generateId();
        System.out.println(id);
    }
}
