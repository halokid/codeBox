package org.r00txx.controller

import org.r00txx.entity.Book
import org.r00txx.service.BookService
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

import javax.annotation.Resource
import javax.xml.ws.RequestWrapper

@RestController
class BookController {

    @Resource
    BookService service;

    @RequestMapping('/books')
    List<Book> getBooks() {
        service.getBooks()
    }

    @RequestMapping('/book/{id}')
    Book getBook(@PathVariable(name = 'id') int id) {
        service.getBook(id)
    }
}