package com.camp.promotion.enums;

public enum ProductStatusEnum {

    ONLINE(1),
    OFFLINE(2),
    DELETE(3);

    private final int code;

    ProductStatusEnum(int code) {
        this.code = code;
    }

    public int getCode() {
        return code;
    }
}
