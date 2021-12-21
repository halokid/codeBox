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

import org.springframework.boot.test.context.SpringBootTestContextBootstrapper;
import org.springframework.core.annotation.AnnotatedElementUtils;
import org.springframework.test.context.TestContextBootstrapper;

/**
 * {@link TestContextBootstrapper} for {@link MybatisTest @MybatisTest} support.
 *
 * @author Kazuki Shimizu
 * @since 2.1.0
 */
class MybatisTestContextBootstrapper extends SpringBootTestContextBootstrapper {

  @Override
  protected String[] getProperties(Class<?> testClass) {
    MybatisTest annotation = AnnotatedElementUtils.getMergedAnnotation(testClass, MybatisTest.class);
    return (annotation != null) ? annotation.properties() : null;
  }

}
