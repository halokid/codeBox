package com.example.springbootmybatispostgres.repositories;
import com.example.springbootmybatispostgres.entities.Book;
import org.apache.ibatis.annotations.*;
import org.springframework.stereotype.Repository;

import java.util.List;

@Mapper
@Repository
public interface BookRepository {

    String SELECT_FROM_BOOK_WHERE_ID = "SELECT * FROM book WHERE id = #{id}";
    String SELECT_FROM_BOOK = "select * from book";

    @Select(SELECT_FROM_BOOK)
    List<Book> findAll();

    @Select(SELECT_FROM_BOOK_WHERE_ID)
    Book findById(long id);

    @Select("SELECT * FROM book WHERE title = #{title}")
    Book findByTitle(String title);

    @Delete("DELETE FROM book WHERE id = #{id}")
    boolean deleteById(long id);

    @Insert("INSERT INTO book(title, isbn, description, page, price) " +
            " VALUES (#{title}, #{isbn}, #{description}, #{page}, #{price})")
    void insert(Book book);

    @Update("Update book set title=#{title}, " +
            " isbn=#{isbn}, description=#{description}, page=#{page}, price=#{price} where id=#{id}")
    int update(Book book);
}