package codebox.yy.keepalive;

import java.io.Serializable;
import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * 维持连接的消息对象。
 * <p>
 * 创建时间：2010-7-18 上午12:22:09
 * @author HouLei
 * @since 1.0
 */
public class KeepAlive implements Serializable{

    private static final long serialVersionUID = -2813120366138988480L;

    /* 覆盖该方法，仅用于测试使用。
     * @see java.lang.Object#toString()
     */
    @Override
    public String toString() {
        return new SimpleDateFormat("yyyy-MM-dd HH:mm:ss").format(new Date())+"\t维持连接包";
    }

}



