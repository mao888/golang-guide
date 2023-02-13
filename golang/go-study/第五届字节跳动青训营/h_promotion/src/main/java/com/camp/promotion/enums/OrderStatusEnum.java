package com.camp.promotion.enums;

public enum OrderStatusEnum {

    CANCEL(0),
    NON_PAYMENT(1),
    PAID(2),
    SEND(3),
    SUCCESS(4),
    CLOSE(5);

    private final int code;

    OrderStatusEnum(int code) {
        this.code = code;
    }

    public int getCode() {
        return code;
    }
}
