package org.r00txx.ali;

/**
 * Created by r00xx on 2017/8/10.
 */

import com.aliyuncs.profile.DefaultProfile;
import com.aliyuncs.profile.IClientProfile;
import com.aliyuncs.DefaultAcsClient;
import com.aliyuncs.IAcsClient;
import com.aliyuncs.exceptions.ClientException;
import com.aliyuncs.exceptions.ServerException;
import com.aliyuncs.ecs.model.v20140526.*;

public class Sample {
  public static void main(String[] args) {

    //create DefaultAcsClient & init
    DefaultProfile profile = DefaultProfile.getProfile(
                              "your region id",
                              "access key",
                              "secret");
    IAcsClient client = new DefaultAcsClient(profile);

    //create API request & setting params
    DescribeInstancesRequest request = new DescribeInstancesRequest();
    request.setPageNumber(10);

    //make request & handle exception
    try {
      DescribeInstancesResponse response = client.getAcsResponse(request);
    }
    catch (ServerException e) {
      e.printStackTrace();
    }
    catch (ClientException e) {
      e.printStackTrace();
    }

  }
}

















