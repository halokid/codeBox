<?php

$_ = "Hello";
$__ = "World";
$___ = "foo";

echo "$_, $__, $___\n";

echo "{$_}, {$__}, {$___}\n";


foreach( ['a' => 'Alpha', 'b' => 'Beta', 'c' => 'Gamma' ] as $letter => $_ ) {
	echo $letter;
}
echo "\n";
echo $_;

