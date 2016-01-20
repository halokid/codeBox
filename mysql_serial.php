<?php

$s = microtime(true);
echo $s."\n";



$link1 = mysql_connect('172.16.2.111', 'test', 'xxxx');
mysql_select_db('test');

mysql_query("SET NAMES utf8");


$sql = "select * from goods";
$res = mysql_query($sql);

while( $row = mysql_fetch_array($res) ) {
	$rows[] = $row;
}
echo count($rows)."\n";




$link2 = mysql_connect('172.16.2.111', 'test', 'xxxx');
mysql_select_db('test');

mysql_query("SET NAMES utf8");


$sql2 = "select * from goods";
$res2 = mysql_query($sql2);

while( $row2 = mysql_fetch_array($res2) ) {
	$rows2[] = $row2;
}
echo count($rows2)."\n";










$e = microtime(true);

echo $e."\n";

echo ($e-$s);





