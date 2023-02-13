package com.camp.promotion.limit;

public enum RateEnum {

    RATE_100_PER_SECONDS(100, 1),
    RATE_1000_PER_SECONDS(1000, 1),
    RATE_3000_PER_SECONDS(3000, 1),
    RATE_5000_PER_SECONDS(5000, 1);

    private int count;

    private int duration;

    RateEnum(int count, int duration) {
        this.count = count;
        this.duration = duration;
    }

    public int getCount() {
        return count;
    }

    public int getDuration() {
        return duration;
    }
}
