package com.camp.promotion.common;


public enum ResponseEnum {

    SUCCESS(0,"SUCCESS"),
    ERROR(1001,"ERROR"),
    ILLEGAL_ARGUMENT(1002,"ILLEGAL_ARGUMENT"),
    ACTIVITY_NOT_EXIST(1003, "ACTIVITY_NOT_EXIST"),
    ACTIVITY_NOT_START(1004, "ACTIVITY_NOT_START"),
    ACTIVITY_OFFLINE(1005, "ACTIVITY_OFFLINE"),
    USER_NOT_LOGIN(1006, "USER_NOT_LOGIN"),
    USER_UNDER_RISK(1007, "USER_UNDER_RISK"),
    OUT_OF_STOCK(1008, "OUT_OF_STOCK"),
    INVALID_CACHE_KEY(1009, "INVALID_CACHE_KEY"),
    CREATE_ORDER_FAIL(1010, "CREATE_ORDER_FAIL"),
    CREATE_ACTIVITY_FAIL(1011, "CREATE_ACTIVITY_FAIL"),
    OVER_RATE_LIMIT(1012, "OVER_RATE_LIMIT"),
    CONSUMER_FAIL(1013, "CONSUMER_FAIL"),
    UNEXPECTED(-1, "UNEXPECTED");

    private final int status;
    private final String msg;


    ResponseEnum(int status, String msg){
        this.status = status;
        this.msg = msg;
    }

    public int getStatus(){
        return status;
    }
    public String getMsg(){
        return msg;
    }

}
