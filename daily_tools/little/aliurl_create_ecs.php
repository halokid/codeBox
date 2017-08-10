<?php
//create aliyun instance 

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


/**
https://ecs.aliyuncs.com/?Action=CreateInstance
&RegionId=cn-hangzhou
&ImageId=_32_23c472_20120822172155_aliguest.vhd
&SecurityGroupId=sg-c0003e8b9
&HostName=Bctest01
&InstanceType=ecs.t1.small
&<公共请求参数>

**/


function compose_url($user_params=''){
    //$timestamp=strftime("%Y-%m-%dT%H:%M:%SZ",time());
    $timestamp=gmdate("Y-m-d\TH:i:s\Z");
    $access_key_id='xxxxxxxxxxxx';
    $access_key_secret='yyyyyyyyyyyyyyyyyyy';
    $parameters = array(
            // 'Action'            =>    'DescribeInstances',
            'Action'            =>    'CreateInstance',
            'Format'            =>    'JSON', 
            'Version'           =>    '2014-05-26', 
            'AccessKeyId'       =>    $access_key_id, 
            'SignatureVersion'  =>    '1.0', 
            'SignatureMethod'   =>    'HMAC-SHA1', 
            'SignatureNonce'    =>    uuid(), 
            'Timestamp'         =>    $timestamp,
            'RegionId'          =>    'cn-shenzhen',
            'ImageId'           =>    'centos_7_03_64_40G_alibase_20170710.vhd',
            'InstanceType'      =>    'ecs.sn1ne.large',
            // 'SecurityGroupId'   =>    'sg-wz9cnteembe2hls8hpm3',
            'SecurityGroupId'   =>    'sg-wz93840wx55ispcc1gvl',
            'HostName'          =>    'jimmy-test-ecs',
            

    );
    $signature = compute_signature($parameters, $access_key_secret);
    $canonicalizedQueryString='Signature='.$signature;
    foreach ($parameters as $k => $v) {
         $canonicalizedQueryString .= '&' . $k . '=' . percent_encode($v);
    }   
    
    
    
    
    $ch = curl_init();

    // 2. 设置请求选项, 包括具体的url
    curl_setopt($ch, CURLOPT_URL, "http://ecs.aliyuncs.com/?".$canonicalizedQueryString);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($ch, CURLOPT_HEADER, 0);
    
    // 3. 执行一个cURL会话并且获取相关回复
    $response = curl_exec($ch);
    echo $response;
    // echo "<hr>";
    
    // 4. 释放cURL句柄,关闭一个cURL会话
    curl_close($ch);

    
}

compose_url();
?>