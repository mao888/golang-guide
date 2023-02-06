package com.camp.promotion.convertet;

public interface MultiConverter<T, R, U> {

    U convert(ConvertFunction<T, R, U> f);
}
