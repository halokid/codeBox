<?php
require 'vendor/autoload.php';

// $article = new Forum\Article();
$article = new Discussion\Forum\Article();
echo $article->msg();
echo "\n";


// $b = new Board\Article();
// echo $b->msg();