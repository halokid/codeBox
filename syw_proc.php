<?php


$qudao_arr = array ('baidu.com',
									'yahoo.com',
									'360.com',
									'google.com'
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











