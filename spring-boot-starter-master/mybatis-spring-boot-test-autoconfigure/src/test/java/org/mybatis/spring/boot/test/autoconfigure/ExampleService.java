/*
 *    Copyright 2015-2021 the original author or authors.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
package org.mybatis.spring.boot.test.autoconfigure;

import org.springframework.stereotype.Service;

/**
 * Example component that annotated {@link Service @Service} used with {@link MybatisTest} tests.
 *
 * @author Kazuki Shimizu
 * @since 1.2.1
 */
@Service
public class ExampleService {

  public String getMessage() {
    return "Goodbye!";
  }

}
