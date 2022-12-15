package org.halokid;

import cn.smallbun.screw.core.Configuration;
import cn.smallbun.screw.core.engine.EngineConfig;
import cn.smallbun.screw.core.engine.EngineFileType;
import cn.smallbun.screw.core.engine.EngineTemplateType;
import cn.smallbun.screw.core.execute.DocumentationExecute;
import cn.smallbun.screw.core.process.ProcessConfig;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.ApplicationContext;
import org.springframework.test.context.junit4.SpringRunner;

import javax.sql.DataSource;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

/**
 * <h1>数据库表文档生成</h1>
 */
@SpringBootTest
//@SpringBootTest(classes = DBDocTest.class)        // todo: 错误做法， 装载不了主yaml文件
@RunWith(SpringRunner.class)
public class DBDocTest {

    @Autowired
    private ApplicationContext applicationContext;
    @Test
    public void buildDBDoc(){
        DataSource dataSourceMysql= applicationContext.getBean(DataSource.class);

        //生成文件配置
        EngineConfig engineConfig = EngineConfig.builder()
                // 生成文件路径
//                .fileOutputDir("/Users/mac/Desktop/ideaweb/e-commerce-springcloud/e-commerce-authority-center/src/main/resources/mysql\n")
            .fileOutputDir("/Users/kxy/gitcode/codeBox/ecommercespringcloud/e-commerce-authority-center/src/main/resources/mysql_schema")
            .openOutputDir(false)
                // 文件类型  有html MD
                .fileType(EngineFileType.HTML)
                .produceType(EngineTemplateType.freemarker).build();

        // 生成文档配置，包含有自定义版本号、描述等等
        // 数据库名_description_version.html
        Configuration config= Configuration.builder()
                .version("1.0.0")
                .description("e-commerce-springcloud")
                .dataSource(dataSourceMysql)
                .engineConfig(engineConfig)
                .produceConfig(getProduceConfig())
                .build();
        // 执行生成
        new DocumentationExecute(config).execute();
    }


    /**
     * <h2>配置想要生成的数据表和想要忽略的数据表</h2>
     */
    private ProcessConfig getProduceConfig(){
        // 想要忽略的数据表
        List<String> ignoreTableName = Collections.singletonList("undo_log");
        //  忽略表前缀，忽略 a,b开头的数据库表
        List<String> ignorePrefix = Arrays.asList("a","b");
        // 忽略表后缀
        List<String> ignoreSuffix = Arrays.asList("_test","_Test");

        return ProcessConfig.builder()
                // 根据名称指定表生成
                .designatedTableName(Collections.emptyList())
                // 根据表前缀生成
                .designatedTablePrefix(Collections.emptyList())
                // 根据表后缀生成
                .designatedTableSuffix(Collections.emptyList())
                // 按照表名称忽略
                .ignoreTableName(ignoreTableName)
                // 按照前缀忽略
                .ignoreTablePrefix(ignorePrefix)
                // 按照后缀忽略
                .ignoreTableSuffix(ignoreSuffix)
                .build();

    }
}
