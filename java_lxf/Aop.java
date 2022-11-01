public class Aop {

}

class Book {

}

class Transaction {
  public void rollback() {}
  public void commit() {}
}

class BookService {
  public void createBook(Book book) {
    securityCheck();
    Transaction tx = startTransaction();
    try {
      tx.commit();
    } catch (RuntimeException e) {
      tx.rollback();
      throw e;
    }
    log("create book: " + book);
  }

  public void updateBook(Book book) {
    securityCheck();
    Transaction tx = startTransaction();
    try {
      // 核心业务逻辑
      tx.commit();
    } catch (RuntimeException e) {
      tx.rollback();
      throw e;
    }
    log("updated book: " + book);
  }

  private void securityCheck() {

  }

  Transaction startTransaction() {
    Transaction tx = new Transaction();
    return tx;
  }

  void log(String s) {}
}


/*
// todo: Proxy design pattern
public class SecurityCheckBookService implements BookService {
  private final BookService target;

  public SecurityCheckBookService(BookService target) {
    this.target = target;
  }

  public void createBook(Book book) {
    securityCheck();
    target.createBook(book);
  }

  public void updateBook(Book book) {
    securityCheck();
    target.updateBook(book);
  }

  public void deleteBook(Book book) {
    securityCheck();
    target.deleteBook(book);
  }

  private void securityCheck() {
        ...
  }
}
*/

// todo:
//如果我们以AOP的视角来编写上述业务，可以依次实现：
//  核心逻辑，即BookService；
//  切面逻辑，即：
//  权限检查的Aspect；
//  日志的Aspect；
//  事务的Aspect。
// 然后，以某种方式，让框架来把上述3个Aspect以Proxy的方式“织入”到BookService中，这样一来，就不必编写复杂而冗长的Proxy模式。







