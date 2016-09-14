<?php
function myHash($str) {
  // hash(i) = hash(i-1) * 33 + str[i]
  $hash = 0;
  $s = md5($str);
  $seed = 5;
  $len = 32;
  
  for ($i = 0; $i < $len; $i++) {
    $hash = ($hash << $seed) + $hash + ord($s{$i});
  }
  
  return $hash & 0x7FFFFFFF;
}


class ConsistentHash {
  //server list 
  private $_server_list = array();
  // 延迟排序， 因为可能会执行多次 addServer
  private $_lazy_sorted = FALSE;
  
  public function addServer($server) {
    $hash = myHash($server);
    $this->_lazy_sorted = FALSE;
    
    if (!isset($this->_server_list[$hash)) {
      $this->_server_list[$hash] = $server;
    }
    
    return $this;     //引用对象自身，可继续 addServer
  }
  
  public function find($key) {
    // 排序
    if (!$this->_lazy_sorted) {
      asort($this->_server_list);
      $this->_lazy_sorted = TRUE;
    }
    
    $hash = myHash($key);
    $len = sizeof($this->_server_list);
    if ($len == 0) {
      return FALSE;
    }
    
    $keys = array_keys($this->_server_list);      //这个就是 圆形 上面的数值大小
    $values = array_values($this->_server_list);
    
    // 如果不在区间内， 则返回最后一个 server
    if ($hash <= $keys[0] || $hash >= $keys[$len - 1] ) {
      return $values[$len - 1];
    }
    
    foreach ($keys as $key=>$pos) {
      $next_pos = NULL;
      
      if (isset($keys[$key + 1])) {       //像PHP这种先检查是否有，再定义变量的方式真的好吗？？
        $next_pos = $keys[$key +1 ];
      }
      
      if (is_null($next_pos)) {
        return $values[$key];
      }
      
      //区间判断
      if ($hash >= $pos && $hash <= $next_pos) {
        return $values[$key];
      }
    }
    
  }
}


$ch = new ConsistentHash();
$ch->addServer("serv1")->addServer("serv2")->addServer("serv3");

echo "key1 at ".$ch->find("key1");
echo "key2 at ".$ch->find("key2");
echo "key3 at ".$ch->find("key3");













