package com.camp.promotion.convertet;

public interface ConvertFunction<T, R, U> {
    U apply(T t, R r);
}
