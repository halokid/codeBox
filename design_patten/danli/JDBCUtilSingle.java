import java.sql.*;

public final class JDBCUtilSingle {
  static String url = "jdbc://mysql:///test";
  static String name = "root";
  static String password = "xxx";
  static Connection conn = null;

  private static JDBCUtilSingle jdbcUtilSingle = null;

  private static JDBCUtilSingle getInstance() {
    if (jdbcUtilSingle == null) {
      synchronized (JDBCUtilSingle.class) {
        if (jdbcUtilSingle == null) {
          jdbcUtilSingle = new JDBCUtilSingle();
        }
      }
    }
    return jdbcUtilSingle;
  }

  //构造函数为 private， 防止 new
  private JDBCUtilSingle() {

  }

  static {
    try {
      Class.forName("com.mysql.jdbc.Driver");
    } catch (ClassNotFoundException e) {
      e.printStackTrace();
    }
  }

  public Connection getConnection() {
    try {
      conn = DriverManager.getConnection(url, name, password);
    } catch (SQLException e) {
      e.printStackTrace();
    }
    return conn;
  }

  public void closeConnection(ResultSet rs, Statement statement, Connection con) {
      try {
        if (rs != null) {
          rs.close();
        }
      } catch (SQLException e) {
        e.printStackTrace();
      } finally {
        try {
          if (statement != null) {
            statement.close();
          }
        } catch (Exception e) {
          e.printStackTrace();
        } finally {
          try {
            if (con != null) {
              con.close();
            }
          } catch (SQLException e) {
            e.printStackTrace();
          }
        }
      }
    }
  }

