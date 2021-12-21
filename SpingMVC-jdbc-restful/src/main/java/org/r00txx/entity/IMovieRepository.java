package org.r00txx.entity;

import org.r00txx.models.Movie;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

/**
 * Created by r00xx<82049406@qq.com> on 2016/12/29.
 */
public interface IMovieRepository extends CrudRepository<Movie, Long> {
  List<Movie> findByYearLessThan(int year);
}
