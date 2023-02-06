package com.camp.promotion.convertet;

import java.util.function.Function;

public interface Converter<T, R> {

    R convert(Function<T, R> f);
}
