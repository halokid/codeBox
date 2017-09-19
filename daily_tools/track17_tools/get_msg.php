<?php

//$ll = system("python sample.py", $res);
$ll = exec("python sample.py", $res);

echo $ll;
echo "\n-----------------------------\n";
//print_r($res);

echo "\n-----------------------------\n";

$arr = json_decode($ll);
print_r($arr);
