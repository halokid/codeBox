<?php
/**
 * User: r00t
 * Date: 2018/2/16
 * Time: 16:32
 */

/*
 * 限流控制
 */
class RateLimit
{
  private $min_num = 60;            //单个用户每分访问数
  private $day_num = 100000;        //单个用户每天总的访问量

  public function min_limit($uid) {
    //每分钟访问的令牌
    $min_num_key = $uid . '_min_num';
    //每天访问的令牌
    $day_num_key = $uid . '_day_num';

    $res_min = $this->getRedis($min_num_key, $this->min_num, 60);
    $res_day = $this->getRedis($day_num_key, $this->day_num, 86400);

    if (!$res_min['status'] || !$res_day['status']) {
      exit($res_min['msg'] . $res_day['msg']);
    }
  }


  public function getRedis($key, $init_num, $expire) {
    $now_time = time();
    $result = ['status' => true, 'msg' => ''];
    $res_obj = $this->di->get('redis');
    $redis->watch($key);
    $limit_val = $redis->get($key);
    if ($limitVal) {
      $limitVal = json_decode($limitVal, true);
      $newNum   = min($init_num, ($limitVal['num'] - 1) + (($init_num / $expire) * ($now_time - $limit_val['time'])));
      if ($newNum > 0) {
        $redis_val = json_encode(['num' => $newNum, 'time' => time()]);
      } else {
        return ['status' => false, 'msg' => '当前时刻令牌消耗完！'];
      }
    } else {
      $redis_val = json_encode(['num' => $initNum, 'time' => time()]);
    }

    $redis->multi();
    $redis->set($key, $redis_val);
    $rob_res = $redis->exec();
    if (!$rob_res) {
      $res = ['status' => false, 'msg' => '访问次数过多'];
    }
    return $res;
  }

}






