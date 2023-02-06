package com.camp.promotion.handler;

import com.camp.promotion.common.BizException;
import com.camp.promotion.common.ResponseData;
import com.camp.promotion.common.ResponseEnum;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@Slf4j
//@RestControllerAdvice
public class BizExceptionHandler {

//    @ExceptionHandler(value = Exception.class)
    public ResponseData<?> exceptionHandler(Exception e) {
        if (e instanceof BizException) {
            log.error("biz exec error, e = ", e);
            return ResponseData.Create((BizException) e);
        }

        log.error("biz exec error, e = ", e);
        return ResponseData.Create(ResponseEnum.UNEXPECTED);
    }
}
