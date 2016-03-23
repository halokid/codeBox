<?php

/**
统一的唯一标识数据添加,包含内，外部
* */


/**
$url_arr = array ('baidu.com/aa/bb',
				  'yahoo.com/cc/dd',
				  '360.com/ee/ff',
				  'google.com/gg/hh'
                );
**/


$sessid_arr = array ('li9t209jm7mc6m4vmn88o5a7j0',
                     'li9t209jm7mc769vmn88o5a345',
                     'l8hy209jm7mc769vmn88o5a999',
                     'li99j2m7mc769vmnss88o5a875',
                     'li9t27mskklc769vmn88o5appp'
                   );


$qudao_tags = array ('baidu',
				     'yahoo',
				     '360',
				     'google'
					);


$ip_arr = array ('10.10.10.29',
				 '55.55.55.22',
				 '87.56.45.59',
				 '48.15.46.54',
				 '48.15.46.78'
			     );


$sywurl_arr = array( 'www.xxxxxx.com/goods/gooddetail/1020',
                     'www.xxxxxx.com/bsq',
                     'www.xxxxxx.com/koubei',
                     'www.xxxxxx.com/help/999',
                     'www.xxxxxx.com/member/xxx',
                     'www.xxxxxx.com'
                   );


$rds = new Redis();
$rds->connect('127.0.0.1', 6379 );
$rds->select(1);

for( $i=0; $i< 100; $i++ ) {
	$t = microtime(true);
    //$rds->set( $sessid_arr[array_rand($sessid_arr)].':'.$t.':'.$ip_arr[array_rand($ip_arr)].':0',  $t.':'.$url_arr[array_rand($url_arr)].':'.$ip_arr[array_rand($ip_arr)].':0:0:0:0' );	

    $key = $sessid_arr[array_rand($sessid_arr)].':'.$t.':'.$ip_arr[array_rand($ip_arr)].':'.$qudao_tags[array_rand($qudao_tags)].':0';

    //$rds->set( $sessid_arr[array_rand($sessid_arr)].':'.$t.':'.$ip_arr[array_rand($ip_arr)].':'.$qudao_tags[array_rand($qudao_tags)].':0', 'www.xxxxxx.com/goods/gooddetail/1020:'.microtime(true) );	
    //$rds->set( $key, 'www.xxxxxx.com/goods/gooddetail/1020:'.microtime(true) );	


    for( $j=0; $j<100; $j++ ) {
      $res = $rds->zAdd( $key, microtime(true).rand(1,10), $sywurl_arr[array_rand($sywurl_arr)] );   
    }


	//'sleep(1);
}
















