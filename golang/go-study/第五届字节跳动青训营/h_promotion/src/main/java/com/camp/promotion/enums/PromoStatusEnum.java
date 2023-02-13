package com.camp.promotion.enums;

public enum PromoStatusEnum {

    INIT(0),
    ONLINE(1),
    OFFLINE(2);

    private final int code;

    PromoStatusEnum(int code) {
        this.code = code;
    }

    public int getCode() {
        return code;
    }
}
