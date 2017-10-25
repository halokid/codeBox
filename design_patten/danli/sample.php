
<?php

class Danli
{
  private static $_instance;


  private function __construct() {
    echo 'this is a __construct method';
  }


  public function __clone() {
    trigger_error('clone is not allow', E_USER_ERROR);
  }


  public static function getInstance() {
    if (!(self::$_instance instanceof self)) {
      self::$_instance = new self;
    }
    return self::$_instance;
  }

 
  public function test() {
    echo 'call method success';
  }

}

//首先获取单例是获取这个类的单例，要记住是这个类的单例
$danli = Danli::getInstance();

//然后再用这个类的单例去调用这个类的方法
$danli->test();









