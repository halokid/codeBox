<?php
//这个是原始数据（在实际应用里，这个原始数组可以是从mysql查询出来，或者是从json读出去的，只不过这里以PHP的数组形式来呈现而已）
$ar = array (
  array('id'=>1, 'pid'=>0), 
  array('id'=>2, 'pid'=>0), 
  array('id'=>3, 'pid'=>2), 
  array('id'=>4, 'pid'=>0), 
  array('id'=>5, 'pid'=>3), 
  array('id'=>6, 'pid'=>1), 
  array('id'=>7, 'pid'=>1), 
  array('id'=>8, 'pid'=>6), 
  array('id'=>9, 'pid'=>7), 
  array('id'=>10, 'pid'=>9) 
);

//sort func
function cmd($a, $b) {
  if ($a['pid'] == $b['pid']) return 0;
  return $a['pid'] > $b['pid'] ? 1 : -1;
}

//排序，为避免数据中父节点在子节点后面出现， 这种情况在多次修改数据后经常会发生的
//排序的目的就是防止这种情况造成的混乱
// print_r($ar);

uasort($ar, 'cmd');

print_r($ar);
echo "\n----------------------------------------------------\n";
//define target array
//定义目标数组, 这个数组就是 B+树 的数据形式, 只不过在PHP里面是以数组的形式呈现
$d = array();
//define index array, use to record the position of the node in the target array
//定义索引数组， 用于记录节点在目标数组的位置
//这个就是原来的原始数据编程 B+树 之后，索引数据的组织形式，只不过在这里用PHP的数组形
//式来组织而已，把这个数据序列化之后，写入到文件里面去，这个就是索引文件了
$ind = array();     
foreach ($ar as $v) {
  /**
  $v['child'] = array(); //给每一个节点附加一个child项
  if ($v['pid'] == 0) {
    $i = count($d);
    $d[$i] = $v;
    $ind[$v['id']] = & $d[$i];
  } 
  else {
    $i = count($ind[$v['pid']]['child']);
    $ind[$v['pid']]['child'][$i] = $v;
    $ind[$v['id']] = & $ind[$v['pid']]['child'][$i];
  }
  **/
  // print_r($v);
  // echo "-----------------------------\n";
  
  //--------------------------------
$v['child'] = array(); //给每个节点附加一个child项 
if($v['pid'] == 0) { 
  $i = count($d);     //按pid为0的顺序来
  $d[$i] = $v;   
  //索引文件，为了快速查询， 下标就是原始数据所对应的数组（这里的原始数据也是PHP的数
  //组，尼玛貌似PHP好像就数组一种数据结构）...
  $ind[$v['id']] =& $d[$i];       //引用了 $d 指针内存， 修改 ind 和同时修改 d  
}
else {
  $i = count($ind[$v['pid']]['child']);
  //如果 $v['pid'] 不为0， 则该元素肯定为 $v['pid'] 的 child
  $ind[$v['pid']]['child'][$i] = $v; 
  //赋予节点真实的数据, 便于索引查询数据
  $ind[$v['id']] =& $ind[$v['pid']]['child'][$i]; 
} 
//-------------------------------------
}


// $d 最终组织成的 B+ 树的数据是数组，  $ind 索引数据也是一个数组， 他们的每个元素的
//值都是一样的，但是下标不同， $d按照数组的自然下标 ( 0, 1, 2, 3, 4, 5, 6, ....),
// 但是索引数据的话下面就是原始数据所对应的下标， 所以 $ind 头三个元素的下标跟别是 1, 2, 4 
print_r($d);
echo "\n----------------------------------------------------\n";
print_r($ind);





