package org.r00txx.controller;



import java.util.List;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.r00txx.api.Greeting;
import org.r00txx.dao.userDao;
import org.r00txx.entity.IMovieRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import org.r00txx.models.Movie;

@RestController
public class RestServiceController {

  @Autowired
  private IMovieRepository repo;

  static final Logger logger = LogManager.getLogger(RestServiceController.class.getName());

  // CREATE
  @RequestMapping("/movies/create")
  @ResponseBody
  public String createMovie(String title, int year) {
    Movie movie = new Movie(title, year);
    try {
      repo.save(movie);
    } catch (Exception e) {
      logger.error(e.getMessage());
      return e.getMessage();
    }
    return "creation successful: " + String.valueOf(movie.getId());
  }

  // READ， 返回普通字符串
  // 这里可以转化为 json， 那么整个程序就完整了，后期再加上自定义sql的用法
  @RequestMapping("/movies/read")
  @ResponseBody
  public String readMovie(long id) {
    Movie movie;
    try {
      movie = repo.findOne(id);
    } catch (Exception e) {
      logger.error(e.getMessage());
      return e.getMessage();
    }
    if (movie == null) {
      String errorMst = "no movie found for id " + id;
      logger.error(errorMst);
      return errorMst;
    } else {
      return movie.getTitle() + " : " + movie.getYear();
    }
  }


  // READ API，这个是返回json的例子
  // 这里可以转化为 json， 那么整个程序就完整了，后期再加上自定义sql的用法
  @RequestMapping("/movies/readapi")
  @ResponseBody
  public Greeting readMovieApi(long id) {

//    return  new Greeting(1, "jimmy");

//    /**

    Movie movie;
    movie = repo.findOne(id);

    System.out.println("------------------ output API --------------");
    return new Greeting(id, movie.getTitle());

//     */
  }


  // 返回空的对象
  @RequestMapping("/movies/testapi")
  @ResponseBody
  public Object[] testApi() {

//    return  new Greeting(1, "jimmy");

//    /**


    System.out.println("------------------ output API --------------");
    return new Object[] {100, "hello"};

//     */
  }



  @Autowired
  private userDao userdao;
  // 自定义sql语句查询的列子
  @RequestMapping("/movies/sql")
  @ResponseBody
  public Object[] testSql() {

//    return  new Greeting(1, "jimmy");

//    /**



    System.out.println("------------------ output API --------------");
//    return new Object[] {100, "hello"};
    String tmp = userdao.queryUsers();
    return new Object[] {100, tmp};
//     */
  }



//  @Autowired
//  private userDao userdao;
  // 返回多维度json的列子
  @RequestMapping("/movies/mutilapi")
  @ResponseBody
  public Object[] testMuli() {

//    return  new Greeting(1, "jimmy");

//    /**



    System.out.println("------------------ output API --------------");
//    return new Object[] {100, "hello"};
//    String tmp = userdao.queryUsers();
    Object a = new Object[] {1, "aaaa"};
    Object b = new Object[] {2, "bbbb"};

    return new Object[] {a, b};
//     */
  }


  // UPDATE
  @RequestMapping("/movies/update")
  @ResponseBody
  public String readMovie(long id, String title, int year) {
    Movie movie;
    try {
      movie = repo.findOne(id);
      movie.setTitle(title);
      movie.setYear(year);
      repo.save(movie);
    } catch (Exception e) {
      logger.error(e.getMessage());
      return e.getMessage();
    }
    return movie.getTitle() + " : " + movie.getYear();
  }

  // DELETE
  @RequestMapping("/movies/delete")
  @ResponseBody
  public String deleteMovie(long id) {
    try {
      repo.delete(id);
    } catch (Exception e) {
      logger.error(e.getMessage());
      return e.getMessage();
    }
    return "deletion successful";
  }

  @RequestMapping("/movies/readAllBeforeYear")
  public List<Movie> getMoviesBeforeYear(@RequestParam(value = "year") int year) {
    List<Movie> movies = repo.findByYearLessThan(year);
    return movies;
  }
}


