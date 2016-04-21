<?php

$s = microtime(true);
echo $s."\n";

$link1 = mysqli_connect('172.16.2.111', 'test', 'xxxx', 'test');
$link2 = mysqli_connect('172.16.2.112', 'test', 'xxxx', 'test');

//mysqli_query("SET NAMES utf8");
mysqli_set_charset ($link1,'utf8');
mysqli_set_charset ($link2,'utf8');


$link1->query("select * from goods", MYSQLI_ASYNC);
$link2->query("select * from goods", MYSQLI_ASYNC);
$all_links = array($link1, $link2);



$processed = 0;

do {
	$links = $errors = $reject = array();
	foreach ( $all_links as $link ) {
		$links[] = $errors[] = $reject[] = $link;
	}

	if ( !mysqli_poll($links, $errors, $reject, 1 ) ) {
		continue;
	}

	// $i=0;	
	$rows = array();
	foreach( $links as $link ) {
		if ( $result = $link->reap_async_query() ) {

			//print_r($result->fetch_row());
		  
			while ( $row = $result->fetch_row() ) {
				//print_r($row);
				// $rows[$i][] = $row;
				$rows[] = $row;
			}
			echo count($rows[$i])."\n";

			if ( is_object($result ) ) {
				mysqli_free_result($result);
			}

			// $i++;
		}
		else die( sprintf("error:	%s\n", mysqli_error($link) ));

		$processed++;
	}

} while( $processed < count($all_links ) );






$e = microtime(true);

echo $e."\n";

echo ($e-$s);


