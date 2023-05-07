package com.example.springbootmybatispostgres.service;

import com.example.springbootmybatispostgres.entities.Book;

import java.util.List;

public interface BookService {

    Book create(Book book);

    List<Book> getAll();

    Book getOne(long id);

    void deleteById(long id);

    Book getByTitle(String title);
}
