<?php
include 'BoardArticle.php';
include 'ForumArticle.php';

$article = new Forum\Article();
echo $article->msg();
echo "\n";


$post = new Board\Article();
echo $post->msg();