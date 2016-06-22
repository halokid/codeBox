<?php 
include 'BoardArticle.php';
include 'ForumArticle.php';

use Forum\Article;

$article = new Article();
echo $article->msg()."\n";
