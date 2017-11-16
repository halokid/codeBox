<?php
/*
 * @Author: r00x.tactx 
 * @Date: 2017-10-20 15:51:20 
 * @Last Modified by: xx.tactx
 * @Last Modified time: 2017-10-20 17:52:39
 * PHP的超时实现

 思路很简单：链接一个后端，然后设置为非阻塞模式，如果没有连接上就一直循环，判断当前时间和超时时间之间的差异。

 phpsocket中实现原始的超时：(每次循环都当前时间去减，性能会很差，cpu占用会较高)
 */


$host="127.0.0.1";
$port=”80″;

$timeout=15;//timeoutinseconds

$socket=socket_create(AF_INET,SOCK_STREAM,SOL_TCP) or die("Unabletocreatesocketn");

//务必设置为阻塞模式
socket_set_nonblock($socket) or die("Unabletosetnonblockonsocketn");

$time=time();   //这个明显有问题， 应该放在 while 里面

//循环的时候每次都减去相应值

while(!@socket_connect($socket,$host,$port))//如果没有连接上就一直死循环

{

$err=socket_last_error($socket);

if($err==115||$err==114)

{

if((time()-$time)>=$timeout)//每次都需要去判断一下是否超时了

{

socket_close($socket);

die("Connectiontimedout.n");

}

sleep(1);

continue;

}

die(socket_strerror($err)."n");

}

//还原阻塞模式
socket_set_block($this->socket) or die("Unabletosetblockonsocketn");


