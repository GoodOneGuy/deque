# deque

# 简单暴力实现
直接用slice，头部插入/删除时全部元素移动。这种实现的主要问题是双端队列头部插入和删除的操作很频繁，时间复杂度O(n)。
# list实现
list本质是双向链表，可以直接实现双端队列。这种实现的问题在于随机访问元素的时间复杂度是O(n), 并且指针占用空间。
# list+slice实现或者二维slice
如图，和c++中的deque容器实现基本思路一致，处理算法本身比较复杂外，比较完美的算法。对内存最友好。
![image](https://github.com/GoodOneGuy/deque/assets/131322884/2438b4c4-d9f3-432e-8bcc-d351b7b7e638)

# slice实现环形队列，必要时扩容迁移数据
算法相对简单，内存一旦申请之后，不主动退还。本项目采用这种实现方式
![image](https://github.com/GoodOneGuy/deque/assets/131322884/c3f38f63-a0e1-4619-82db-ab00f62683cd)

