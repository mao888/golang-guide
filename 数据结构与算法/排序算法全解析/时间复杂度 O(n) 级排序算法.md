<a name="TeXMe"></a>
## 计数排序
前文说到，1959 年 77 月，希尔排序通过交换非相邻元素，打破了O(n^2) 的魔咒，使得排序算法的时间复杂度降到了 O(n log⁡n)级，此后的快速排序、堆排序都是基于这样的思想，所以他们的时间复杂度都是 O(n log⁡n)。

那么，排序算法最好的时间复杂度就是 O(n log⁡n)吗？是否有比 O(n log⁡n) 级还要快的排序算法呢？能否在 O(n) 的时间复杂度下完成排序呢

事实上，O(n) 级的排序算法存在已久，但他们只能用于特定的场景。

计数排序就是一种时间复杂度为 O(n) 的排序算法，该算法于 1954 年由 `Harold H. Seward` 提出。在对一定范围内的整数排序时，它的复杂度为 Ο(n+k)（其中 `k` 是整数的范围大小）。

<a name="ZQZkt"></a>
### 伪计数排序
举个例子，我们需要对一列数组排序，这个数组中每个元素都是 [1, 9] 区间内的整数。那么我们可以构建一个长度为 9 的数组用于计数，计数数组的下标分别对应区间内的 9 个整数。然后遍历待排序的数组，将区间内每个整数出现的次数统计到计数数组中对应下标的位置。最后遍历计数数组，将每个元素输出，输出的次数就是对应位置记录的次数。

<a name="f2bee699"></a>
### 动图演示
![](https://cdn.nlark.com/yuque/0/2022/gif/22219483/1668355005125-1af67e58-edae-4b92-b07d-722b16c2e49a.gif#averageHue=%23020100&clientId=u753af7e9-ae4e-4&crop=0&crop=0&crop=1&crop=1&height=403&id=sPai5&originHeight=607&originWidth=607&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=udf7c52d2-c988-43ad-b2fc-eeffe6a3513&title=&width=403)<br />算法实现如下（以 [1, 9] 为例 ）：
```java
public static void countingSort9(int[] arr) {
    // 建立长度为 9 的数组，下标 0~8 对应数字 1~9
    int[] counting = new int[9];
    // 遍历 arr 中的每个元素
    for (int element : arr) {
        // 将每个整数出现的次数统计到计数数组中对应下标的位置
        counting[element - 1]++;
    }
    int index = 0;
    // 遍历计数数组，将每个元素输出
    for (int i = 0; i < 9; i++) {
        // 输出的次数就是对应位置记录的次数
        while (counting[i] != 0) {
            arr[index++] = i + 1;
            counting[i]--;
        }
    }
}
```

算法非常简单，但这里的排序算法 **并不是** 真正的计数排序。因为现在的实现有一个非常大的弊端：排序完成后，`arr` 中记录的元素已经不再是最开始的那个元素了，他们只是值相等，但却不是同一个对象。

在纯数字排序中，这个弊端或许看起来无伤大雅，但在实际工作中，这样的排序算法几乎无法使用。因为被排序的对象往往都会携带其他的属性，但这份算法将被排序对象的其他属性都丢失了。

就好比业务部门要求我们将 1 号商品，2号商品，3 号商品，4 号商品按照价格排序，它们的价格分别为 8 元、6 元，6 元，9 元。 我们告诉业务部门：排序完成后价格为 6 元、 6 元、8元，9 元，但不知道这些价格对应哪个商品。这显然是不可接受的。

<a name="mcIoY"></a>
### 伪计数排序 2.0
对于这个问题，我们很容易想到一种解决方案：在统计元素出现的次数时，同时把真实的元素保存到列表中，输出时，从列表中取真实的元素。算法实现如下：

```java
public static void countingSort9(int[] arr) {
    // 建立长度为 9 的数组，下标 0~8 对应数字 1~9
    int[] counting = new int[9];
    // 记录每个下标中包含的真实元素，使用队列可以保证排序的稳定性
    HashMap<Integer, Queue<Integer>> records = new HashMap<>();
    // 遍历 arr 中的每个元素
    for (int element : arr) {
        // 将每个整数出现的次数统计到计数数组中对应下标的位置
        counting[element - 1]++;
        if (!records.containsKey(element - 1)) {
            records.put(element - 1, new LinkedList<>());
        }
        records.get(element - 1).add(element);
    }
    int index = 0;
    // 遍历计数数组，将每个元素输出
    for (int i = 0; i < 9; i++) {
        // 输出的次数就是对应位置记录的次数
        while (counting[i] != 0) {
            // 输出记录的真实元素
            arr[index++] = records.get(i).remove();
            counting[i]--;
        }
    }
}
```

在这份代码中，我们通过队列来保存真实的元素，计数完成后，将队列中真实的元素赋到 `arr` 列表中，这就解决了信息丢失的问题，并且使用队列还可以保证排序算法的稳定性。

但是，**这也不是** 真正的计数排序，计数排序中使用了一种更巧妙的方法解决这个问题。

<a name="EsQfp"></a>
### 真正的计数排序
举个例子，班上有 1010 名同学：他们的考试成绩分别是：7,8,9,7,6,7,6,8,6,6，他们需要按照成绩从低到高坐到 0～9共 10 个位置上。<br />用计数排序完成这一过程需要以下几步：

- 第一步仍然是计数，统计出：4 名同学考了 6 分，3 名同学考了 7 分，2名同学考了 8 分，1 名同学考了 9 分；
- 然后从头遍历数组：第一名同学考了 7 分，共有 4 个人比他分数低，所以第一名同学坐在 4 号位置（也就是第 5 个位置）；
- 第二名同学考了 8 分，共有 7 个人（4 + 3）比他分数低，所以第二名同学坐在 7 号位置；
- 第三名同学考了 9 分，共有 9 个人（4 + 3 + 2）比他分数低，所以第三名同学坐在 9 号位置；
- 第四名同学考了 7 分，共有 4 个人比他分数低，并且之前已经有一名考了 7 分的同学坐在了 4 号位置，所以第四名同学坐在 5号位置。
- ...依次完成整个排序

区别就在于计数排序并不是把计数数组的下标直接作为结果输出，而是通过计数的结果，计算出每个元素在排序完成后的位置，然后将元素赋值到对应位置。

代码如下：
```java
public static void countingSort9(int[] arr) {
    // 建立长度为 9 的数组，下标 0~8 对应数字 1~9
    int[] counting = new int[9];
    // 遍历 arr 中的每个元素
    for (int element : arr) {
        // 将每个整数出现的次数统计到计数数组中对应下标的位置
        counting[element - 1]++;
    }
    // 记录前面比自己小的数字的总数
    int preCounts = 0;
    for (int i = 0; i < counting.length; i++) {
        int temp = counting[i];
        // 将 counting 计算成当前数字在结果中的起始下标位置。位置 = 前面比自己小的数字的总数。
        counting[i] = preCounts;
        // 当前的数字比下一个数字小，累计到 preCounts 中
        preCounts += temp;
    }
    int[] result = new int[arr.length];
    for (int element : arr) {
        // counting[element - 1] 表示此元素在结果数组中的下标
        int index = counting[element - 1];
        result[index] = element;
        // 更新 counting[element - 1]，指向此元素的下一个下标
        counting[element - 1]++;
    }
    // 将结果赋值回 arr
    for (int i = 0; i < arr.length; i++) {
        arr[i] = result[i];
    }
}
```

首先我们将每位元素出现的次数记录到 `counting` 数组中。

然后将 `counting[i]` 更新为数字 `i` 在最终排序结果中的起始下标位置。这个位置等于前面比自己小的数字的总数。<br />例如本例中，考 7 分的同学前面有 4 个比自己分数低的同学，所以 7 对应的下标为 4。<br />这一步除了使用 `temp` 变量这种写法以外，还可以通过多做一次减法省去 `temp` 变量：

```java
// 记录前面比自己小的数字的总数
int preCounts = 0;
for (int i = 0; i < counting.length; i++) {
    // 当前的数字比下一个数字小，累计到 preCounts 中
    preCounts += counting[i];
    // 将 counting 计算成当前数字在结果中的起始下标位置。位置 = 前面比自己小的数字的总数。
    counting[i] = preCounts - counting[i];
}
```

接下来从头访问 `arr` 数组，根据 `counting` 中计算出的下标位置，将 `arr` 的每个元素直接放到最终位置上，然后更新 `counting` 中的下标位置。这一步中的 `index` 变量也是可以省略的。

最后将 `result` 数组赋值回 `arr`，完成排序。

这就是计数排序的思想，我们还剩下最后一步，那就是根据 `arr` 中的数字范围计算出计数数组的长度。使得计数排序不仅仅适用于 [1,9]，代码如下：
```java
public static void countingSort(int[] arr) {
    // 判空及防止数组越界
    if (arr == null || arr.length <= 1) return;
    // 找到最大值，最小值
    int max = arr[0];
    int min = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > max) max = arr[i];
        else if (arr[i] < min) min = arr[i];
    }
    // 确定计数范围
    int range = max - min + 1;
    // 建立长度为 range 的数组，下标 0~range-1 对应数字 min~max
    int[] counting = new int[range];
    // 遍历 arr 中的每个元素
    for (int element : arr) {
        // 将每个整数出现的次数统计到计数数组中对应下标的位置，这里需要将每个元素减去 min，才能映射到 0～range-1 范围内
        counting[element - min]++;
    }
    // 记录前面比自己小的数字的总数
    int preCounts = 0;
    for (int i = 0; i < range; i++) {
        // 当前的数字比下一个数字小，累计到 preCounts 中
        preCounts += counting[i];
        // 将 counting 计算成当前数字在结果中的起始下标位置。位置 = 前面比自己小的数字的总数。
        counting[i] = preCounts - counting[i];
    }
    int[] result = new int[arr.length];
    for (int element : arr) {
        // counting[element - min] 表示此元素在结果数组中的下标
        result[counting[element - min]] = element;
        // 更新 counting[element - min]，指向此元素的下一个下标
        counting[element - min]++;
    }
    // 将结果赋值回 arr
    for (int i = 0; i < arr.length; i++) {
        arr[i] = result[i];
    }
}
```

这就是完整的计数排序算法。

<a name="OJJn0"></a>
### 倒序遍历的计数排序
计数排序还有一种写法，在计算元素在最终结果数组中的下标位置这一步，不是计算初始下标位置，而是计算最后一个下标位置。最后倒序遍历 `arr` 数组，逐个将 `arr` 中的元素放到最终位置上。

代码如下：
```java
public static void countingSort(int[] arr) {
    // 防止数组越界
    if (arr == null || arr.length <= 1) return;
    // 找到最大值，最小值
    int max = arr[0];
    int min = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > max) max = arr[i];
        else if (arr[i] < min) min = arr[i];
    }
    // 确定计数范围
    int range = max - min + 1;
    // 建立长度为 range 的数组，下标 0~range-1 对应数字 min~max
    int[] counting = new int[range];
    // 遍历 arr 中的每个元素
    for (int element : arr) {
        // 将每个整数出现的次数统计到计数数组中对应下标的位置，这里需要将每个元素减去 min，才能映射到 0～range-1 范围内
        counting[element - min]++;
    }
    // 每个元素在结果数组中的最后一个下标位置 = 前面比自己小的数字的总数 + 自己的数量 - 1。我们将 counting[0] 先减去 1，后续 counting 直接累加即可
    counting[0]--;
    for (int i = 1; i < range; i++) {
        // 将 counting 计算成当前数字在结果中的最后一个下标位置。位置 = 前面比自己小的数字的总数 + 自己的数量 - 1
        // 由于 counting[0] 已经减了 1，所以后续的减 1 可以省略。
        counting[i] += counting[i - 1];
    }
    int[] result = new int[arr.length];
    // 从后往前遍历数组，通过 counting 中记录的下标位置，将 arr 中的元素放到 result 数组中
    for (int i = arr.length - 1; i >= 0; i--) {
        // counting[arr[i] - min] 表示此元素在结果数组中的下标
        result[counting[arr[i] - min]] = arr[i];
        // 更新 counting[arr[i] - min]，指向此元素的前一个下标
        counting[arr[i] - min]--;
    }
    // 将结果赋值回 arr
    for (int i = 0; i < arr.length; i++) {
        arr[i] = result[i];
    }
}
```

两种算法的核心思想是一致的，并且都是稳定的。第一种写法理解起来简单一些，第二种写法在性能上更好一些。在计算下标位置时，不仅计算量更少，还省去了 `preCounts` 这个变量。在《算法导论》一书中，便是采用的此种写法。

实际上，这个算法最后不通过倒序遍历也能得到正确的排序结果，但这里只有通过倒序遍历的方式，才能保证计数排序的稳定性。

<a name="HSSCb"></a>
### 时间复杂度 & 空间复杂度

从计数排序的实现代码中，可以看到，每次遍历都是进行 `n` 次或者 `k` 次，所以计数排序的时间复杂度为 O(n + k)，`k` 表示数据的范围大小。

用到的空间主要是长度为 `k` 的计数数组和长度为 `n` 的结果数组，所以空间复杂度也是O(n + k)。

需要注意的是，一般我们分析时间复杂度和空间复杂度时，常数项都是忽略不计的。但计数排序的常数项可能非常大，以至于我们无法忽略。不知你是否注意到计数排序的一个非常大的隐患，比如我们想要对这个数组排序：
```java
int[] arr = new int[]{1, Integer.MAX_VALUE};
```

尽管它只包含两个元素，但数据范围是 [1, 2^{31}]，我们知道 `java` 中 `int` 占 44 个字节，一个长度为 2^31 次方的 `int` 数组大约会占 8G8G 的空间。如果使用计数排序，仅仅排序这两个元素，声明计数数组就会占用超大的内存，甚至导致 `OutOfMemory` 异常。

所以计数排序只适用于数据范围不大的场景。例如对考试成绩排序就非常适合计数排序，如果需要排序的数字中存在一位小数，可以将所有数字乘以 10，再去计算最终的下标位置。

<a name="MJEtS"></a>
### 计数排序与 O(nlog⁡n)O(n \log n) 级排序算法的本质区别

前文说到，希尔排序通过交换间隔较远的元素突破了排序算法时间复杂度 O(n^2) 的下界。同样地，我们接下来就一起分析一下，计数排序凭什么能够突破 O(nlog⁡n)的下界呢？它和之前介绍的 O(nlog⁡n)级排序算法的本质区别是什么？

这个问题我们可以从决策树的角度和概率的角度来理解。

<a name="pSOHC"></a>
#### 决策树
决策树是一棵完全二叉树，它可以反映比较排序算法中对所有元素的比较操作。

以包含三个整数的数组 [a, b, c] 为例，基于比较的排序算法的排序过程可以抽象为这样一棵 **决策树**：

![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1668355015980-fe9ef132-a150-4022-8957-db7503d160fc.png#averageHue=%23f3f3f3&clientId=u753af7e9-ae4e-4&crop=0&crop=0&crop=1&crop=1&id=EO5wu&name=image.png&originHeight=1357&originWidth=2474&originalType=binary&ratio=1&rotation=0&showTitle=false&size=166331&status=done&style=none&taskId=u15faa697-17a9-41bd-8edb-10999b00546&title=)

这棵决策树上的每一个叶结点都对应了一种可能的排列，从根结点到任意一个叶结点之间的最短路径（也称为「简单路径」）的长度，表示的是完成对应排列的比较次数。所以从根结点到叶结点之间的最长简单路径的长度，就表示比较排序算法中最坏情况下的比较次数。

设决策树的高度为 `h`，叶结点的数量为 `l`，排序元素总数为 `n` 。

因为叶结点最多有 n! 个，所以我们可以得到：n!≤ l，又因为一棵高度为 `h` 的二叉树，叶结点的数量最多为2^h ，所以我们可以得到：n!≤ l ≤ 2^h

对该式两边取对数，可得：h≥log (n!)

由斯特林（`Stirling`）近似公式，可知 lg⁡(n!)=O(nlog⁡n)

所以 h≥log⁡(n!)=O(nlog⁡n)

于是我们可以得出以下定理：

> 《算法导论》定理 8.1：在最坏情况下，任何比较排序算法都需要做 O(nlog⁡n) 次比较。

由此我们还可以得到以下推论：
> 《算法导论》推论 8.2：堆排序和归并排序都是渐进最优的比较排序算法。


到这里我们就可以得出结论了，如果基于比较来进行排序，无论怎么优化都无法突破 O(nlog⁡n)的下界。计数排序和基于比较的排序算法相比，根本区别就在于：它不是基于比较的排序算法，而是利用了数字本身的属性来进行的排序。整个计数排序算法中没有出现任何一次比较。

<a name="pnZV2"></a>
#### 概率
相信大家都玩过「猜数字」游戏：一方从 [1, 100] 中随机选取一个数字，另一方来猜。每次猜测都会得到「高了」或者「低了」的回答。怎样才能以最少的次数猜中呢？

答案很简单：二分。

二分算法能够保证每次都排除一半的数字。每次猜测不会出现惊喜（一次排除了多于一半的数字），也不会出现悲伤（一次只排除了少于一半的数字），因为答案的每一个分支都是等概率的，所以它在最差的情况下表现是最好的，猜测的一方在 log⁡n次以内必然能够猜中。

基于比较的排序算法与「猜数字」是类似的，每次比较，我们只能得到 a>b 或者 a≤b两种结果，如果我们把数组的全排列比作一块区域，那么每次比较都只能将这块区域分成两份，也就是说每次比较最多排除掉 1/2 的可能性。

再来看计数排序算法，计数排序时申请了长度为 `k` 的计数数组，在遍历每一个数字时，这个数字落在计数数组中的可能性共有 `k` 种，但通过数字本身的大小属性，我们可以「一次」把它放到正确的位置上。相当于一次排除了 (k−1)/k种可能性。

这就是计数排序算法比基于比较的排序算法更快的根本原因。

<a name="XfTVf"></a>
### 练习
算法题：[力扣 912. 排序数组](https://leetcode-cn.com/problems/sort-an-array/)<br />这道题我们已经练习过很多次了，不知你是否注意到每次提交的执行用时和内存消耗总是令人不太满意。事实上，今天我们学习的计数排序才是最适合用来做这道题目的。<br />因为题目中写明了数据的范围不会很大：-50000 <= nums[i] <= 50000，使用计数排序可以超过 99+\%99+% 的用户！
```java
class Solution {
    public int[] sortArray(int[] nums) {
        countingSort(nums);
        return nums;
    }

    public static void countingSort(int[] arr) {
        // 防止数组越界
        if (arr == null || arr.length <= 1) return;
        // 找到最大值，最小值
        int max = arr[0];
        int min = arr[0];
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > max) max = arr[i];
            else if (arr[i] < min) min = arr[i];
        }
        // 确定计数范围
        int range = max - min + 1;
        // 建立长度为 range 的数组，下标 0~range-1 对应数字 min~max
        int[] counting = new int[range];
        // 遍历 arr 中的每个元素
        for (int element : arr) {
            // 将每个整数出现的次数统计到计数数组中对应下标的位置，这里需要将每个元素减去 min，才能映射到 0～range-1 范围内
            counting[element - min]++;
        }
        // 每个元素在结果数组中的最后一个下标位置 = 前面比自己小的数字的总数 + 自己的数量 - 1。我们将 counting[0] 先减去 1，后续 counting 直接累加即可
        counting[0]--;
        for (int i = 1; i < range; i++) {
            // 将 counting 计算成当前数字在结果中的最后一个下标位置。位置 = 前面比自己小的数字的总数 + 自己的数量 - 1
            // 由于 counting[0] 已经减了 1，所以后续的减 1 可以省略。
            counting[i] += counting[i - 1];
        }
        int[] result = new int[arr.length];
        // 从后往前遍历数组，通过 counting 中记录的下标位置，将 arr 中的元素放到 result 数组中
        for (int i = arr.length - 1; i >= 0; i--) {
            // counting[arr[i] - min] 表示此元素在结果数组中的下标
            result[counting[arr[i] - min]] = arr[i];
            // 更新 counting[arr[i] - min]，指向此元素的前一个下标
            counting[arr[i] - min]--;
        }
        // 将结果赋值回 arr
        for (int i = 0; i < arr.length; i++) {
            arr[i] = result[i];
        }
    }
}
```
算法题：[力扣 1122. 数组的相对排序](https://leetcode-cn.com/problems/relative-sort-array/)<br />本题中元素的范围为 [0, 1000][0,1000]，这个范围很小，所以我们可以考虑使用「计数排序」。
```java
class Solution {
    public int[] relativeSortArray(int[] arr1, int[] arr2) {
        int upper = 0;
        for (int x : arr1) {
            upper = Math.max(upper, x);
        }
        int[] frequency = new int[upper + 1];
        for (int x : arr1) {
            ++frequency[x];
        }
        int[] ans = new int[arr1.length];
        int index = 0;
        for (int x : arr2) {
            for (int i = 0; i < frequency[x]; ++i) {
                ans[index++] = x;
            }
            frequency[x] = 0;
        }
        for (int x = 0; x <= upper; ++x) {
            for (int i = 0; i < frequency[x]; ++i) {
                ans[index++] = x;
            }
        }
        return ans;
    }
}
```
<a name="I6QE6"></a>
## 基数排序
想一下我们是怎么对日期进行排序的。比如对这样三个日期进行排序：2014 年 1 月 7 日，2020 年 1 月 9 日，020 年 7 月 10 日。

我们大脑中对日期排序的思维过程是：

- 先看年份，2014 比 2020 要小，所以 2014 年这个日期应该放在其他两个日期前面。
- 另外两个日期年份相等，所以我们比较一下月份，1 比 7 要小，所以 1 月这个日期应该放在 7 月这个日期前面

这种利用多关键字进行排序的思想就是基数排序，和计数排序一样，这也是一种线性时间复杂度的排序算法。其中的每个关键字都被称作一个基数。

比如我们对 999,997,866,666 这四个数字进行基数排序，过程如下：

- 先看第一位基数：6 比 8 小，8 比 9 小，所以 666 是最小的数字，866 是第二小的数字，暂时无法确定两个以 9 开头的数字的大小关系
- 再比较 9 开头的两个数字，看他们第二位基数：9 和 9 相等，暂时无法确定他们的大小关系
- 再比较 99 开头的两个数字，看他们的第三位基数：7 比 9 小，所以 997 小于 999

基数排序有两种实现方式。本例属于「最高位优先法」，简称 `MSD (Most significant digital)`，思路是从最高位开始，依次对基数进行排序。

与之对应的是「最低位优先法」，简称 `LSD (Least significant digital)`。思路是从最低位开始，依次对基数进行排序。使用 `LSD` 必须保证对基数进行排序的过程是稳定的。

通常来讲，`LSD` 比 `MSD` 更常用。以上述排序过程为例，因为使用的是 `MSD`，所以在第二步比较两个以 99 开头的数字时，其他基数开头的数字不得不放到一边。体现在计算机中，这里会产生很多临时变量。

但在采用 `LSD` 进行基数排序时，每一轮遍历都可以将所有数字一视同仁，统一处理。所以 `LSD` 的基数排序更符合计算机的操作习惯。

基数排序最早是用在卡片排序机上的，一张卡片有 80 列，类似一个 80 位的整数。机器通过在卡片不同位置上穿孔表示当前基数的大小。卡片排序机的排序过程就是采用的 `LSD` 的基数排序。

<a name="sjtHp"></a>
### 动图演示
简单起见，我们先只考虑对非负整数排序的情况。<br />![](https://cdn.nlark.com/yuque/0/2022/gif/22219483/1668355005207-d14ee58e-3da9-422a-a3ac-cf14e623a185.gif#averageHue=%23020100&clientId=u753af7e9-ae4e-4&crop=0&crop=0&crop=1&crop=1&id=H2NuF&originHeight=608&originWidth=608&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=uc17a1c5e-3939-49b4-b43e-679b429c70d&title=)

基数排序可以分为以下三个步骤：

- 找出数组中最大的数字的位数 `maxDigitLength`
- 获取数组中每个数字的基数
- 遍历 `maxDigitLength` 轮数组，每轮按照基数对其进行排序

<a name="Jhw7N"></a>
### 找出数组中最大的数字的位数
首先找到数组中的最大值：
```java
public static void radixSort(int[] arr) {
    if (arr == null) return;
    int max = 0;
    for (int value : arr) {
        if (value > max) {
            max = value;
        }
    }
    //...
}
```

通过遍历一次数组，找到了数组中的最大值 `max`，然后我们计算这个最大值的位数：
```java
int maxDigitLength = 0;
while (max != 0) {
    maxDigitLength++;
    max /= 10;
}
```

将 `maxDigitLength` 初始化为 0，然后不断地除以 10，每除一次，`maxDigitLength` 就加一，直到 `max` 为 0。

读者可能会有疑惑，如果 `max` 初始值就是 0 呢？严格来讲，0 在数学上属于 1 位数。

但实际上，基数排序时我们无需考虑 `max` 为 0 的场景，因为 `max` 为 0只有一种可能，那就是数组中所有的数字都为0，此时数组已经有序，我们无需再进行后续的排序过程。

<a name="B7qgo"></a>
### 获取基数

获取基数有两种做法：<br />第一种：
```java
int mod = 10;
int dev = 1;
for (int i = 0; i < maxDigitLength; i++) {
    for (int value : arr) {
        int radix = value % mod / dev;
        // 对基数进行排序
    }
    mod *= 10;
    dev *= 10;
}
```

第二种：
```java
int dev = 1;
for (int i = 0; i < maxDigitLength; i++) {
    for (int value : arr) {
        int radix = value / dev % 10;
        // 对基数进行排序
    }
    dev *= 10;
}
```

两者的区别是先做除法运算还是先做模运算，推荐使用第二种写法，因为它可以节省一个变量。

<a name="CEwTT"></a>
### 对基数进行排序
对基数进行排序非常适合使用我们在上一节中学习的计数排序算法，因为每一个基数都在 [0, 9] 之间，并且计数排序是一种稳定的算法。

`LSD` 方式的基数排序代码如下：
```java
public class RadixSort {

    public static void radixSort(int[] arr) {
        if (arr == null) return;
        // 找出最大值
        int max = 0;
        for (int value : arr) {
            if (value > max) {
                max = value;
            }
        }
        // 计算最大数字的长度
        int maxDigitLength = 0;
        while (max != 0) {
            maxDigitLength++;
            max /= 10;
        }
        // 使用计数排序算法对基数进行排序
        int[] counting = new int[10];
        int[] result = new int[arr.length];
        int dev = 1;
        for (int i = 0; i < maxDigitLength; i++) {
            for (int value : arr) {
                int radix = value / dev % 10;
                counting[radix]++;
            }
            for (int j = 1; j < counting.length; j++) {
                counting[j] += counting[j - 1];
            }
            // 使用倒序遍历的方式完成计数排序
            for (int j = arr.length - 1; j >= 0; j--) {
                int radix = arr[j] / dev % 10;
                result[--counting[radix]] = arr[j];
            }
            // 计数排序完成后，将结果拷贝回 arr 数组
            System.arraycopy(result, 0, arr, 0, arr.length);
            // 将计数数组重置为 0
            Arrays.fill(counting, 0);
            dev *= 10;
        }
    }
}
```

计数排序的思想上一节已经介绍过，这里不再赘述。当每一轮对基数完成排序后，我们将 `result` 数组的值拷贝回 `arr` 数组，并且将 `counting` 数组中的元素都置为 0，以便在下一轮中复用。

<a name="KPu01"></a>
### 对包含负数的数组进行基数排序

如果数组中包含负数，如何进行基数排序呢？<br />我们很容易想到一种思路：将数组中的每个元素都加上一个合适的正整数，使其全部变成非负整数，等到排序完成后，再减去之前加的这个数就可以了。

但这种方案有一个缺点：加法运算可能导致数字越界，所以必须单独处理数字越界的情况。

事实上，有一种更好的方案解决负数的基数排序。那就是在对基数进行计数排序时，申请长度为 19 的计数数组，用来存储 [-9, 9] 这个区间内的所有整数。在把每一位基数计算出来后，加上 9，就能对应上 `counting` 数组的下标了。也就是说，`counting` 数组的下标 [0, 18] 对应基数 [-9, 9]。

代码如下：
```java
public class RadixSort {

    public static void radixSort(int[] arr) {
        if (arr == null) return;
        // 找出最长的数
        int max = 0;
        for (int value : arr) {
            if (Math.abs(value) > max) {
                max = Math.abs(value);
            }
        }
        // 计算最长数字的长度
        int maxDigitLength = 0;
        while (max != 0) {
            maxDigitLength++;
            max /= 10;
        }
        // 使用计数排序算法对基数进行排序，下标 [0, 18] 对应基数 [-9, 9]
        int[] counting = new int[19];
        int[] result = new int[arr.length];
        int dev = 1;
        for (int i = 0; i < maxDigitLength; i++) {
            for (int value : arr) {
                // 下标调整
                int radix = value / dev % 10 + 9;
                counting[radix]++;
            }
            for (int j = 1; j < counting.length; j++) {
                counting[j] += counting[j - 1];
            }
            // 使用倒序遍历的方式完成计数排序
            for (int j = arr.length - 1; j >= 0; j--) {
                // 下标调整
                int radix = arr[j] / dev % 10 + 9;
                result[--counting[radix]] = arr[j];
            }
            // 计数排序完成后，将结果拷贝回 arr 数组
            System.arraycopy(result, 0, arr, 0, arr.length);
            // 将计数数组重置为 0
            Arrays.fill(counting, 0);
            dev *= 10;
        }
    }
}
```

代码中主要做了两处修改：

- 当数组中存在负数时，我们就不能简单的计算数组的最大值了，而是要计算数组中绝对值最大的数，也就是数组中最长的数
- 在获取基数的步骤，将计算出的基数加上 9，使其与 `counting` 数组下标一一对应

<a name="UgeEn"></a>
### LSD VS MSD

前文介绍的基数排序都属于 `LSD`，接下来我们看一下基数排序的 `MSD` 实现。
```java
public class RadixSort {

    public static void radixSort(int[] arr) {
        if (arr == null) return;
        // 找到最大值
        int max = 0;
        for (int value : arr) {
            if (Math.abs(value) > max) {
                max = Math.abs(value);
            }
        }
        // 计算最大长度
        int maxDigitLength = 0;
        while (max != 0) {
            maxDigitLength++;
            max /= 10;
        }
        radixSort(arr, 0, arr.length - 1, maxDigitLength);
    }

    // 对 arr 数组中的 [start, end] 区间进行基数排序
    private static void radixSort(int[] arr, int start, int end, int position) {
        if (start == end || position == 0) return;
        // 使用计数排序对基数进行排序
        int[] counting = new int[19];
        int[] result = new int[end - start + 1];
        int dev = (int) Math.pow(10, position - 1);
        for (int i = start; i <= end; i++) {
            // MSD, 从最高位开始
            int radix = arr[i] / dev % 10 + 9;
            counting[radix]++;
        }
        for (int j = 1; j < counting.length; j++) {
            counting[j] += counting[j - 1];
        }
        // 拷贝 counting，用于待会的递归
        int[] countingCopy = new int[counting.length];
        System.arraycopy(counting, 0, countingCopy, 0, counting.length);
        for (int i = end; i >= start; i--) {
            int radix = arr[i] / dev % 10 + 9;
            result[--counting[radix]] = arr[i];
        }
        // 计数排序完成后，将结果拷贝回 arr 数组
        System.arraycopy(result, 0, arr, start, result.length);
        // 对 [start, end] 区间内的每一位基数进行递归排序
        for (int i = 0; i < counting.length; i++) {
            radixSort(arr, i == 0 ? start : start + countingCopy[i - 1], start + countingCopy[i] - 1, position - 1);
        }
    }

}
```

使用 `MSD` 时，下一轮排序只应该发生在当前轮次基数相等的数字之间，对每一位基数进行递归排序的过程中会产生许多临时变量。

相比 `LSD`，`MSD` 的基数排序显得较为复杂。因为我们每次对基数进行排序后，无法将所有的结果一视同仁地进行下一轮排序，否则下一轮排序会破坏本次排序的结果。

<a name="yI8ve"></a>
### 时间复杂度 & 空间复杂度
无论 `LSD` 还是 `MSD`，基数排序时都需要经历 `maxDigitLength` 轮遍历，每轮遍历的时间复杂度为 O(n + k) ，其中 `k` 表示每个基数可能的取值范围大小。如果是对非负整数排序，则 `k = 10`，如果是对包含负数的数组排序，则 `k = 19`。

所以基数排序的时间复杂度为 O(d(n + k))(`d` 表示最长数字的位数，`k` 表示每个基数可能的取值范围大小)。

使用的空间和计数排序是一样的，空间复杂度为 O(n + k)（`k` 表示每个基数可能的取值范围大小）。

<a name="jnjLb"></a>
### 练习
算法题：[力扣 164. 最大间距](https://leetcode-cn.com/problems/maximum-gap/)<br />题目中明确要求在线性时间复杂度和空间复杂度的条件下解决此问题，非常适合使用本节中使用的基数排序算法。<br />由于数组中所有的元素都是非负整数，所以计数数组的长度只需声明为 10 即可。

代码如下：
```java
class Solution {
    public int maximumGap(int[] nums) {
        if (nums.length < 2) return 0;
        radixSort(nums);
        int result = 0;
        for (int i = 1; i < nums.length; i++) {
            result = Math.max(result, nums[i] - nums[i - 1]);
        }
        return result;
    }

    public static void radixSort(int[] arr) {
        if (arr == null) return;
        // 找出最大值
        int max = 0;
        for (int value : arr) {
            if (value > max) {
                max = value;
            }
        }
        // 计算最大数字的长度
        int maxDigitLength = 0;
        while (max != 0) {
            maxDigitLength++;
            max /= 10;
        }
        // 使用计数排序算法对基数进行排序
        int[] counting = new int[10];
        int[] result = new int[arr.length];
        int dev = 1;
        for (int i = 0; i < maxDigitLength; i++) {
            for (int value : arr) {
                int radix = value / dev % 10;
                counting[radix]++;
            }
            for (int j = 1; j < counting.length; j++) {
                counting[j] += counting[j - 1];
            }
            // 使用倒序遍历的方式完成计数排序
            for (int j = arr.length - 1; j >= 0; j--) {
                int radix = arr[j] / dev % 10;
                result[--counting[radix]] = arr[j];
            }
            // 计数排序完成后，将结果拷贝回 arr 数组
            System.arraycopy(result, 0, arr, 0, arr.length);
            // 将计数数组重置为 0
            Arrays.fill(counting, 0);
            dev *= 10;
        }
    }
}
```
算法题：[力扣 561. 数组拆分 I](https://leetcode-cn.com/problems/array-partition-i/)<br />分析题意可知，只需将数组排序后，以间隔 2 拆分数组即可。排序的过程可以用基数排序实现，由于题目中指定了数据的范围是 [-10^4, 10^4]，包含负数，所以计数数组需要声明为 19。

代码如下：
```java
class Solution {
    public int arrayPairSum(int[] nums) {
        radixSort(nums);
        int result = 0;
        for (int i = 0; i < nums.length; i += 2) {
            result+=nums[i];
        }
        return result;
    }

    public static void radixSort(int[] arr) {
        if (arr == null) return;
        // 找出最长的数
        int max = 0;
        for (int value : arr) {
            if (Math.abs(value) > max) {
                max = Math.abs(value);
            }
        }
        // 计算最长数字的长度
        int maxDigitLength = 0;
        while (max != 0) {
            maxDigitLength++;
            max /= 10;
        }
        // 使用计数排序算法对基数进行排序，下标 [0, 18] 对应基数 [-9, 9]
        int[] counting = new int[19];
        int[] result = new int[arr.length];
        int dev = 1;
        for (int i = 0; i < maxDigitLength; i++) {
            for (int value : arr) {
                // 下标调整
                int radix = value / dev % 10 + 9;
                counting[radix]++;
            }
            for (int j = 1; j < counting.length; j++) {
                counting[j] += counting[j - 1];
            }
            // 使用倒序遍历的方式完成计数排序
            for (int j = arr.length - 1; j >= 0; j--) {
                // 下标调整
                int radix = arr[j] / dev % 10 + 9;
                result[--counting[radix]] = arr[j];
            }
            // 计数排序完成后，将结果拷贝回 arr 数组
            System.arraycopy(result, 0, arr, 0, arr.length);
            // 将计数数组重置为 0
            Arrays.fill(counting, 0);
            dev *= 10;
        }
    }
}
```
<a name="P0FMu"></a>
## 桶排序
桶排序的思想是：

- 将区间划分为 `n` 个相同大小的子区间，每个子区间称为一个桶
- 遍历数组，将每个数字装入桶中
- 对每个桶内的数字单独排序，这里需要采用其他排序算法，如插入、归并、快排等
- 最后按照顺序将所有桶内的数字合并起来

桶排序在实际工作中的应用较少，不仅因为它需要借助于其他排序算法，还因为桶排序算法基于一个假设：所有输入数据都服从均匀分布，也就是说输入数据应该尽可能地均匀分布在每个桶中。只有这个假设成立时，桶排序运行效率才比较高。

在最差的情况下，所有数据都会被装入同一个桶中，此时桶排序算法只会徒增一轮遍历。

使用桶排序算法时，我们需要考虑两个因素：

- 设置多少个桶比较合适
- 桶采用哪种数据结构

这两个因素会直接影响到桶排序的内存和效率。

-  **桶的数量**：桶的数量过少，会导致单个桶内的数字过多，桶排序的时间复杂度就会在很大程度上受桶内排序算法的影响。桶的数量过多，占用的内存就会较大，并且会出现较多的空桶，影响遍历桶的效率。具体设置多少个桶需要根据实际情况决定。 
-  **桶的数据结构**：如果将桶的数据结构设置为数组，那么每个桶的长度必须设置为待排序数组的长度，因为我们需要做好最坏的打算，即所有的数字都被装入了同一个桶中，所以这种方案的空间复杂度会很高。<br />那么是不是将桶的数据结构设置为链表就更好呢？使用链表有一个好处，即所有桶的总长度刚好等于待排序数组的长度，不会造成内存浪费。但使用链表也会有一些问题，我们待会一一分析。 

接下来我们就分别学习一下这两种数据结构实现的桶排序算法。
<a name="TZond"></a>
### 以数组作为桶
首先，找到最大值和最小值：
```java
public static void bucketSort(int[] arr) {
    // 判空及防止数组越界
    if (arr == null || arr.length <= 1) return;
    // 找到最大值，最小值
    int max = arr[0];
    int min = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > max) max = arr[i];
        else if (arr[i] < min) min = arr[i];
    }
    // 确定取值范围
    int range = max - min;
    // ...
}
```

这里需要遍历一轮数组。<br />下一步，开始装桶：
```java
// 设置桶的数量，这里我们设置为 100 个，可以根据实际情况修改。
int bucketAmount = 100;
// 桶和桶之间的间距
double gap = range * 1.0 / (bucketAmount - 1);
// 用二维数组来装桶，第一个维度是桶的编号，第二个维度是桶中的数字。每个桶的长度必须设置为 arr.length，因为我们要做好最坏的打算：所有的数字都被装入了一个桶中。
int[][] buckets = new int[bucketAmount][arr.length];
// 单独采用一个数组来记录每个桶当前的长度，也就是当前桶内共有多少个数字。
int[] bucketLength = new int[bucketAmount];
// 装桶
for (int value : arr) {
    // 找到 value 属于哪个桶
    int index = (int) ((value - min) / gap);
    // 装桶后，更新 bucketLength[index]
    buckets[index][bucketLength[index]++] = value;
}
```

我们将桶的数量设置为 100 个，这个值可以根据输入数据的实际情况修改。所有的桶是一个二维数组，第一个维度代表桶的编号，第二个维度代表桶内的数字，每个桶中都有一组数字。

由于每个桶的长度都等于待排序数组的长度，所以我们还需要一个单独的数组来记录当前桶内的有效数字数量。

装桶时需要做一些简单的运算：先通过第一步找到的取值范围计算出每个桶之间的间距，再通过当前数字与最小值的距离除以间距计算出桶的编号，最后根据编号把当前数字放入对应的桶中。

下一步是对每个桶内的数字进行单独排序，这一步需要借助其他排序算法：
```java
// 对每个桶内的数字进行单独排序
int index = 0;
for (int i = 0; i < bucketAmount; i++) {
    if (bucketLength[i] == 0) continue;
    // 取出桶内的数组
    int[] arrInBucket = Arrays.copyOf(buckets[i], bucketLength[i]);
    // 这里需要结合其他排序算法，例如：插入排序
    insertSort(arrInBucket);
    // 排序完成后将桶内的结果收集起来
    System.arraycopy(arrInBucket, 0, arr, index, bucketLength[i]);
    index += bucketLength[i];
}
```

我们以插入排序为例，`insertSort` 函数如下：
```java
// 插入排序
public static void insertSort(int[] arr) {
    // 从第二个数开始，往前插入数字
    for (int i = 1; i < arr.length; i++) {
        int currentNumber = arr[i];
        int j = i - 1;
        // 寻找插入位置的过程中，不断地将比 currentNumber 大的数字向后挪
        while (j >= 0 && currentNumber < arr[j]) {
            arr[j + 1] = arr[j];
            j--;
        }
        // 两种情况会跳出循环：1. 遇到一个小于或等于 currentNumber 的数字，跳出循环，currentNumber 就坐到它后面。
        // 2. 已经走到数列头部，仍然没有遇到小于或等于 currentNumber 的数字，也会跳出循环，此时 j 等于 -1，currentNumber 就坐到数列头部。
        arr[j + 1] = currentNumber;
    }
}
```

这就是以数组作为桶实现的桶排序，它最大的缺点就是每个桶都和待排序数组一样长，非常消耗内存，容易导致「超出内存限制」错误。

我们可以在这份代码的基础上做一个优化：声明时所有的数组都为空，当需要添加数字时，不断扩容，并加入新数字。完整代码如下：
```java
public static void bucketSort(int[] arr) {
    // 判空及防止数组越界
    if (arr == null || arr.length <= 1) return;
    // 找到最大值，最小值
    int max = arr[0];
    int min = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > max) max = arr[i];
        else if (arr[i] < min) min = arr[i];
    }
    // 确定取值范围
    int range = max - min;
    // 设置桶的数量，这里我们设置为 100 个，可以根据实际情况修改。
    int bucketAmount = 100;
    // 桶和桶之间的间距
    double gap = range * 1.0 / (bucketAmount - 1);
    // 用二维数组来装桶，第一个维度是桶的编号，第二个维度是桶中的数字。初始化长度为 0
    int[][] buckets = new int[bucketAmount][];
    // 装桶
    for (int value : arr) {
        // 找到 value 属于哪个桶
        int index = (int) ((value - min) / gap);
        buckets[index] = add(buckets[index], value);
    }
    int index = 0;
    // 对每个桶内的数字进行单独排序
    for (int i = 0; i < bucketAmount; i++) {
        if (buckets[i] == null || buckets[i].length == 0) continue;
        // 这里需要结合其他排序算法，例如：插入排序
        insertSort(buckets[i]);
        // 排序完成后将桶内的结果收集起来
        System.arraycopy(buckets[i], 0, arr, index, buckets[i].length);
        index += buckets[i].length;
    }
}
// 数组扩容
public static int[] add(int[] arr, int num) {
    if (arr == null) return new int[]{num};
    int[] newArr = Arrays.copyOf(arr, arr.length + 1);
    newArr[arr.length] = num;
    return newArr;
}
// 插入排序
public static void insertSort(int[] arr) {
    // 从第二个数开始，往前插入数字
    for (int i = 1; i < arr.length; i++) {
        int currentNumber = arr[i];
        int j = i - 1;
        // 寻找插入位置的过程中，不断地将比 currentNumber 大的数字向后挪
        while (j >= 0 && currentNumber < arr[j]) {
            arr[j + 1] = arr[j];
            j--;
        }
        // 两种情况会跳出循环：1. 遇到一个小于或等于 currentNumber 的数字，跳出循环，currentNumber 就坐到它后面。
        // 2. 已经走到数列头部，仍然没有遇到小于或等于 currentNumber 的数字，也会跳出循环，此时 j 等于 -1，currentNumber 就坐到数列头部。
        arr[j + 1] = currentNumber;
    }
}
```

优化之后，以数组作为桶就不会造成太大的内存消耗了，并且我们不再需要 `bucketLength` 数组来记录桶的长度。

这里的扩容算法和 `ArrayList` 扩容很相似，先开辟一个更长的新数组，并将原数组拷贝过来，再加入新数字。但 `ArrayList` 扩容时，数组长度是先从 0 扩容到 10，后续再不断乘以 1.5 倍，这会造成一定的内存浪费。

无论是数组还是 `ArrayList`，扩容过程都比较耗时，所以这个优化属于用时间换空间。

<a name="NZBvn"></a>
### 以链表作为桶
以链表作为桶的桶排序和以数组作为桶的桶排序思路是类似的，代码如下：
```java
public static void bucketSort(int[] arr) {
    // 判空及防止数组越界
    if (arr == null || arr.length <= 1) return;
    // 找到最大值，最小值
    int max = arr[0];
    int min = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > max) max = arr[i];
        else if (arr[i] < min) min = arr[i];
    }
    // 确定取值范围
    int range = max - min;
    // 设置桶的数量，这里我们设置为 100 个，可以任意修改。
    int bucketAmount = 100;
    // 桶和桶之间的间距
    double gap = range * 1.0 / (bucketAmount - 1);
    HashMap<Integer, LinkedList<Integer>> buckets = new HashMap<>();
    // 装桶
    for (int value : arr) {
        // 找到 value 属于哪个桶
        int index = (int) ((value - min) / gap);
        if (!buckets.containsKey(index)) {
            buckets.put(index, new LinkedList<>());
        }
        buckets.get(index).add(value);
    }
    int index = 0;
    // 对每个桶内的数字进行单独排序
    for (int i = 0; i < bucketAmount; i++) {
        LinkedList<Integer> bucket = buckets.get(i);
        if (bucket == null) continue;
        // 这里需要结合其他排序算法，例如：插入排序
        insertSort(bucket);
        // 排序完成后将桶内的结果收集起来
        for (int num : bucket) {
            arr[index++] = num;
        }
    }
}
// 对链表插入排序
public static void insertSort(LinkedList<Integer> arr) {
    // 从第二个数开始，往前插入数字
    for (int i = 1; i < arr.size(); i++) {
        int currentNumber = arr.get(i);
        int j = i - 1;
        // 寻找插入位置的过程中，不断地将比 currentNumber 大的数字向后挪
        while (j >= 0 && currentNumber < arr.get(j)) {
            arr.set(j + 1, arr.get(j));
            j--;
        }
        // 两种情况会跳出循环：1. 遇到一个小于或等于 currentNumber 的数字，跳出循环，currentNumber 就坐到它后面。
        // 2. 已经走到数列头部，仍然没有遇到小于或等于 currentNumber 的数字，也会跳出循环，此时 j 等于 -1，currentNumber 就坐到数列头部。
        arr.set(j + 1, currentNumber);
    }
}
```

首先，我们仍然是找到数组中的最大值和最小值，确定出数据的取值范围，然后划分 100100 个桶，计算出间距。并且把所有的数字都放入 `LinkedList` 链表中。装桶后，再对链表进行插入排序即可。

采用 `LinkedList`，装桶时不会有额外的空间浪费，但装桶后排序会比较耗时，因为访问 `LinkedList` 链表时，`get` 和 `set` 方法都需要从链表头部开始，逐个向后寻找结点，效率较低。

使用链表排序还有一个问题：由于链表中不能存储基本类型，所以我们不得不将链表类型声明为 `LinkedList<Integer>`，`int` 和 `Integer` 互相转换的过程被称为 「装箱」和「拆箱」，这也会造成额外的性能消耗。

<a name="rUwOG"></a>
### 折中的方案：装桶时用链表，桶内排序用数组
为了解决以上两种数据结构各自的痛点，我们可以采用一种折中的方案：装桶时使用 `LinkedList`，避免扩容问题，桶内排序时将链表转换为数组，再进行排序，避免 `LinkedList` 排序较慢的问题和大量 「装箱」和「拆箱」的性能消耗（整个链表中的 `Integer` 都只需要拆箱一次）。

代码如下：
```java
public static void bucketSort(int[] arr) {
    // 判空及防止数组越界
    if (arr == null || arr.length <= 1) return;
    // 找到最大值，最小值
    int max = arr[0];
    int min = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > max) max = arr[i];
        else if (arr[i] < min) min = arr[i];
    }
    // 确定取值范围
    int range = max - min;
    // 设置桶的数量，这里我们设置为 100 个，可以任意修改。
    int bucketAmount = 100;
    // 桶和桶之间的间距
    double gap = range * 1.0 / (bucketAmount - 1);
    HashMap<Integer, Queue<Integer>> buckets = new HashMap<>();
    // 装桶
    for (int value : arr) {
        // 找到 value 属于哪个桶
        int index = (int) ((value - min) / gap);
        if (!buckets.containsKey(index)) {
            buckets.put(index, new LinkedList<>());
        }
        buckets.get(index).add(value);
    }
    int index = 0;
    // 对每个桶内的数字进行单独排序
    for (int i = 0; i < bucketAmount; i++) {
        Queue<Integer> bucket = buckets.get(i);
        if (bucket == null) continue;
        // 将链表转换为数组，提升排序性能
        int[] arrInBucket = bucket.stream().mapToInt(Integer::intValue).toArray();
        // 这里需要结合其他排序算法，例如：插入排序
        insertSort(arrInBucket);
        // 排序完成后将桶内的结果收集起来
        System.arraycopy(arrInBucket, 0, arr, index, arrInBucket.length);
        index += arrInBucket.length;
    }
}
// 插入排序
public static void insertSort(int[] arr) {
    // 从第二个数开始，往前插入数字
    for (int i = 1; i < arr.length; i++) {
        int currentNumber = arr[i];
        int j = i - 1;
        // 寻找插入位置的过程中，不断地将比 currentNumber 大的数字向后挪
        while (j >= 0 && currentNumber < arr[j]) {
            arr[j + 1] = arr[j];
            j--;
        }
        // 两种情况会跳出循环：1. 遇到一个小于或等于 currentNumber 的数字，跳出循环，currentNumber 就坐到它后面。
        // 2. 已经走到数列头部，仍然没有遇到小于或等于 currentNumber 的数字，也会跳出循环，此时 j 等于 -1，currentNumber 就坐到数列头部。
        arr[j + 1] = currentNumber;
    }
}
```

可以看到，我们在桶内排序前，通过这行代码：
```java
int[] arrInBucket = bucket.stream().mapToInt(Integer::intValue).toArray();
```

将 `LinkedList` 转换为 `int[]`，然后再进行插入排序。代价是这里多了一个中间变量 `arrInBucket`，它会占用 O(n) 的空间，并且 `LinkedList` 转换为 `int[]` 的过程需要遍历一次数组，这会增加 O(n) 的时间。

总结一下桶排序算法中，采用各个数据结构作为桶的特点：

- 以数组作为桶，初始化每个桶的长度为 `n`：时间上做到了最好，但空间占用很高。
- 以数组作为桶，初始化每个桶的长度为 0：空间上做到了最好，但装桶时对数组扩容比较耗时。
- 以 `LinkedList` 作为桶：空间上做到了最好，并且装桶时无需扩容，但对 `LinkedList` 排序比较耗时。
- 装桶时采用 `LinkedList`，排序时采用数组：时间和空间上都是一种折中的方案，但 `LinkedList` 转换 `int[]` 的过程需要遍历一次数组，增加了 O(n) 的时间，转换后会占用 O(n) 的空间。

<a name="jyec7"></a>
### 时间复杂度 & 空间复杂度
桶排序的时间复杂度为 O(n)，需要注意的是，这里 n 的常数项是比较大的，意味着桶排序不一定比 O(nlogn) 级的排序算法快。空间复杂度为 O(n)。
<a name="H1m6U"></a>
### 桶排序 VS (计数排序 || 基数排序)
桶排序也是一种线性时间复杂度的排序算法。许多文章中说计数排序和基数排序都是桶排序的一种特殊情况，但笔者认为这种说法不太准确。

桶排序 VS 计数排序：虽然计数排序也有划分子区间的操作，但是计数排序在统计了每个数字出现的次数后，主要是通过计算每个数字在排序完成后的数组中的最终位置来完成排序，并没有真正把数字装到桶中。而桶排序则是将所有数字装入了桶里，最后从桶里取出每个数字。桶排序的过程比较像我们在计数排序的文章中介绍的「伪计数排序 2.02.0 版本」。

桶排序 VS 基数排序：如果把基数排序看作桶排序，那么基数排序的过程就是不断地装桶，基数排序并没有桶内排序这一步。而桶排序主要分为两步：装桶和桶内排序，桶内排序时需要借助其他排序算法。

并且桶排序基于输入数据均匀分布的假设，计数排序和基数排序则没有这样的限制。

所以笔者认为桶排序和这两种算法还是有明显区别的，在《算法导论》一书中也没有说计数排序和基数排序是桶排序的特例。

<a name="dZvGo"></a>
### 练习
算法题：[力扣 164. 最小差值 I](https://leetcode-cn.com/problems/smallest-range-i/)<br />解析：分析题目可知，A 数组中的每个数可以波动 [−K,K]，其中的最小差值就是 最大值 - k和 最小值 + k的差值，如果差值 <0，则说明每个数字都可以波动成相等的数字，直接返回 0 即可。
```java
class Solution {
    public int smallestRangeI(int[] A, int K) {
        // 判空及防止数组越界
        if (A == null || A.length <= 1) return 0;
        // 找到最大值，最小值
        int max = A[0];
        int min = A[0];
        for (int i = 1; i < A.length; i++) {
            if (A[i] > max) max = A[i];
            else if (A[i] < min) min = A[i];
        }
        int range = (max - K) - (min + K);
        return Math.max(range, 0);
    }
}
```
算法题：[力扣 164. 最大间距](https://leetcode-cn.com/problems/maximum-gap/)<br />解析：设长度为 N 的数组中的最大值和最小值为 max、min，则不难发现相邻数字的最大间距不会小于(max−min)/(N−1)。

因此，我们可以选取整数 d=(max−min)/(N−1)。随后，我们将整个区间划分为若干个大小为 d 的桶，并找出每个整数所在的桶。根据前面的结论，能够知道，元素之间的最大间距一定不会出现在某个桶的内部，而一定会出现在不同桶当中。

因此，在找出每个元素所在的桶之后，我们可以维护每个桶内元素的最大值与最小值。随后，只需从前到后不断比较相邻的桶，用后一个桶的最小值与前一个桶的最大值之差作为两个桶的间距，最终就能得到所求的答案。

代码如下：
```java
class Solution {
    public int maximumGap(int[] nums) {
        int n = nums.length;
        if (n < 2) {
            return 0;
        }
        int minVal = Arrays.stream(nums).min().getAsInt();
        int maxVal = Arrays.stream(nums).max().getAsInt();
        int d = Math.max(1, (maxVal - minVal) / (n - 1));
        int bucketSize = (maxVal - minVal) / d + 1;

        int[][] bucket = new int[bucketSize][2];
        for (int i = 0; i < bucketSize; ++i) {
            Arrays.fill(bucket[i], -1); // 存储 (桶内最小值，桶内最大值) 对， (-1, -1) 表示该桶是空的
        }
        for (int i = 0; i < n; i++) {
            int idx = (nums[i] - minVal) / d;
            if (bucket[idx][0] == -1) {
                bucket[idx][0] = bucket[idx][1] = nums[i];
            } else {
                bucket[idx][0] = Math.min(bucket[idx][0], nums[i]);
                bucket[idx][1] = Math.max(bucket[idx][1], nums[i]);
            }
        }

        int ret = 0;
        int prev = -1;
        for (int i = 0; i < bucketSize; i++) {
            if (bucket[i][0] == -1) {
                continue;
            }
            if (prev != -1) {
                ret = Math.max(ret, bucket[i][0] - bucket[prev][1]);
            }
            prev = i;
        }
        return ret;
    }
}
```
