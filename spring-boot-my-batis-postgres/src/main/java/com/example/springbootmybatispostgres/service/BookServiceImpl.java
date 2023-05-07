package com.example.springbootmybatispostgres.service;

import com.example.springbootmybatispostgres.common.exception.BadRequestException;
import com.example.springbootmybatispostgres.common.exception.DataNotFoundException;
import com.example.springbootmybatispostgres.common.exception.DuplicateException;
import com.example.springbootmybatispostgres.entities.Book;
import com.example.springbootmybatispostgres.repositories.BookRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.util.ObjectUtils;

import java.text.MessageFormat;
import java.util.List;

@RequiredArgsConstructor
@Service
public class BookServiceImpl implements BookService {

    private final BookRepository repository;

    @Override
    public Book create(Book book) {
        Book bookById = getByTitle(book.getTitle());
        if(!ObjectUtils.isEmpty(bookById)){
          throw new DuplicateException(MessageFormat.format("Book {0} already exists in the system", book.getTitle()));
        }
        repository.insert(book);
        return getByTitle(book.getTitle());
    }

    @Override
    public List<Book> getAll(){
        return repository.findAll();
    }

    @Override
    public Book getOne(long id) {
        Book book = repository.findById(id);
        if(ObjectUtils.isEmpty(book)){
           throw new DataNotFoundException(MessageFormat.format("Book id {0} not found", String.valueOf(id)));
        }
        return book;
    }

    @Override
    public void deleteById(long id) {
        boolean isDeleted = repository.deleteById(id);
        if(!isDeleted){
            throw new BadRequestException("Delete error, please check ID and try again");
        }
    }

    @Override
    public Book getByTitle(String title) {
        return repository.findByTitle(title);
    }
}