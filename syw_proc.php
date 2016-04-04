<?php

/**

这个是在客户端已经存有 $_COOKIE['statis']的值之后，客户在
syw站内所有的访问轨迹的数据获取

 * */

//渠道标识
$qudao_arr = array ('baidu',
				    'yahoo',
				    '360',
					'google',
					'other',  //其他所有的访问
				);

$url_arr = array ('/aa/bb',
				  '/cc/dd',
				  '/ee/ff',
				  '/gg/hh'
				);




$ip_arr = array ('10.10.10.29',
				 '55.55.55.22',
				 '87.56.45.59',
				 '48.15.46.54'
				 );

$rds = new Redis();
$rds->connect('127.0.0.1', 6379 );
$rds->select(2);

for( $i=0; $i< 100; $i++ ) {
	$t = microtime(true);
	$rds->zset( $ip_arr[array_rand($ip_arr)],  $t.':'.$url_arr[array_rand($url_arr)].':'.$ip_arr[array_rand($ip_arr)].':0:0:0:0' );	
	sleep(1);
}











