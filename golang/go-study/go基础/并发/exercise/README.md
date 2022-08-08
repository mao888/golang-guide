
# 练习题

1. 使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
    1. 开启一个 goroutine 循环生成int64类型的随机数，发送到`jobChan`
    2. 开启24个 goroutine 从`jobChan`中取出随机数计算各位数的和，将结果发送到`resultChan`
    3. 主 goroutine 从`resultChan`取出结果并打印到终端输出