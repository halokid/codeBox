package com.example.springbootmybatispostgres.web.controller;

import com.example.springbootmybatispostgres.common.exception.BadRequestException;
import com.example.springbootmybatispostgres.entities.Book;
import com.example.springbootmybatispostgres.service.BookService;
import com.example.springbootmybatispostgres.web.response.SuccessResponse;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.util.ObjectUtils;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.text.MessageFormat;
import java.util.List;

@RequiredArgsConstructor
@RestController
@RequestMapping("/api/book")
public class BookController {

    private final BookService bookService;

     @PostMapping()
     public ResponseEntity<SuccessResponse> create(@RequestBody @Valid Book book) {
         if (!ObjectUtils.isEmpty(book.getId())) {
             throw new BadRequestException("A new data cannot already have an ID");
         }

         return new ResponseEntity<>(
                 new SuccessResponse(bookService.create(book), "Successful registration"),
                 HttpStatus.CREATED);
     }

    @GetMapping
    public ResponseEntity<SuccessResponse> getAll() {
        List<Book> books = bookService.getAll();

        return new ResponseEntity<>(new SuccessResponse(books, MessageFormat.format("{0} Results found", books.size())), HttpStatus.OK);
    }

    @GetMapping("/{id}")
    public ResponseEntity<SuccessResponse> getOne(@PathVariable("id") Long id) {
        Book book = bookService.getOne(id);
        return new ResponseEntity<>(
                new SuccessResponse(book, "Result found"), HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<SuccessResponse> delete(@PathVariable("id") Long id) {
        bookService.deleteById(id);
        return new ResponseEntity<>(
                new SuccessResponse(null, "Deletion completed successfully"), HttpStatus.OK);
    }

}