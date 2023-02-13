package com.camp.promotion.enums;

public enum PaymentTypeEnum {

    ONLINE(1);

    private final int code;

    PaymentTypeEnum(int code) {
        this.code = code;
    }

    public int getCode() {
        return code;
    }
}
