same class name in different files,  if u want to instance in the same file, it
will happen error

//BoardArticle.php
<?php
class Article {
//...  
}
?>


//ForumArticle.php
<?php 
class Article {
//..  
}
?>



How fix???  change the class name

<?php 
class BoardArticle {
  
}


class ForumArticle {
  
}
?>


But it's not good style, see v2


