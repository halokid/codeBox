<?php
//生成阿里云API URL的程序


function percent_encode($res){
    $res=trim(utf8_encode(urlencode($res)));
  //$res=utf8_encode($res);
    $res=str_replace(array('+','*','%7E'), array('%20','%2A','~'), $res);
    return $res;
}
function uuid($prefix = '')  
  {  
    $chars = md5(uniqid(mt_rand(), true));  
    $uuid  = substr($chars,0,8) . '-';  
    $uuid .= substr($chars,8,4) . '-';  
    $uuid .= substr($chars,12,4) . '-';  
    $uuid .= substr($chars,16,4) . '-';  
    $uuid .= substr($chars,20,12);  
    return $prefix . $uuid;  
  } 
function compute_signature($parameters, $access_key_secret){
    ksort($parameters);
    $canonicalizedQueryString = '';
    foreach ($parameters as $k => $v) {
      if(empty($canonicalizedQueryString))
        $canonicalizedQueryString .= percent_encode($k) . '=' . percent_encode($v);
      else
         $canonicalizedQueryString .= '&' . percent_encode($k) . '=' . percent_encode($v);
    }    
    //var_dump($canonicalizedQueryString);
    $stringToSign = 'GET&%2F&' . percent_encode($canonicalizedQueryString);
    //var_dump($stringToSign);
    $signature=base64_encode(hash_hmac("sha1",$stringToSign,$access_key_secret."&",true));
    return $signature;     
}

function compose_url($user_params=''){
    //$timestamp=strftime("%Y-%m-%dT%H:%M:%SZ",time());
    $timestamp=gmdate("Y-m-d\TH:i:s\Z");
    $access_key_id='LTAIpZZDd7RHkjB8';
    $access_key_secret='xxx';
    $parameters = array(
            'Format'        => 'JSON', 
            'Version'       => '2017-03-21', 
            'AccessKeyId'   => $access_key_id, 
            'SignatureVersion'  => '1.0', 
            'SignatureMethod'   => 'HMAC-SHA1', 
            'SignatureNonce'    => uuid(), 
            'TimeStamp'         => $timestamp, 
            'Action'            =>'GetPlayInfo',//DescribeCdnService
            // 'DomainName'        =>'xx.xx.xx',  
            // 'AppName'        =>'AppName',  
            // 'StreamName'        =>'StreamName',  
            // 'StartTime'         => '2016-08-01T00:00:00Z', 
            // 'EndTime'         =>'2016-08-09T08:05:06Z', 

    );
    $signature = compute_signature($parameters, $access_key_secret);
    $canonicalizedQueryString='Signature='.$signature;
    foreach ($parameters as $k => $v) {
         $canonicalizedQueryString .= '&' . $k . '=' . percent_encode($v);
    }   
    var_dump("http://cdn.aliyuncs.com/?".$canonicalizedQueryString);

/*    $parameters['Signature'] = $signature;
    $url="http://cdn.aliyuncs.com/?SignatureVersion=1.0&Format=JSON&TimeStamp=".urlencode($parameters['TimeStamp'])."&AccessKeyId=".$parameters['AccessKeyId']
    ."&SignatureMethod=HMAC-SHA1&Version=".$parameters['Version']."&Signature=".urlencode($parameters['Signature'])."&Action=".$parameters['Action']."&SignatureNonce=".$parameters['SignatureNonce'].'&AppName='.$parameters['AppName'].'&DomainName='.$parameters['DomainName'].'&StartTime='.urlencode($parameters['StartTime']).'&EndTime='.urlencode($parameters['EndTime']).'&StreamName='.$parameters['StreamName'];
var_dump($url);die();*/
    // return $url;
}

compose_url();
?>