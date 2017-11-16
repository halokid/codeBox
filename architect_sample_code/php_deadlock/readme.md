

###PHP死锁问题分析：



######SQL造成的死锁:

```
select product where pid=1 for update;  锁product表
update order set status=1 where uid=2 and pid=1   更新order表



select order where uid=2 for update;    因为这里锁了 order 表， 所以上面的sql就会造成死锁
update product set num=num+1 where 

```



- 另外就是PHP程序执行会有超时的问题，所以设置各种 web服务器， PHP-fpm， php.ini的各种超时策略可以解决部分死锁的问题。

- 还有就是各种远程访问的超时， curl， 流处理（sockopen, file操作等）


'''
$tmCurrent=gettimeofday();

$intUSGone=($tmCurrent['sec']-$tmStart['sec'])*1000000 + ($tmCurrent['usec']-$tmStart['usec']);

if($intUSGone>$this->_intReadTimeoutUS){
  returnfalse;
}
'''


- 另外的一些PHP的超时实现












