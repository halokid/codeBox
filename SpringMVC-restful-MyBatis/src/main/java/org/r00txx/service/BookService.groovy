package org.r00txx.service

import org.r00txx.dao.IBookDao
import org.r00txx.entity.Book
import org.springframework.stereotype.Service

import javax.annotation.Resource


@Service
class BookService {
    @Resource
    IBookDao bookDao;

    List<Book> getBooks() {
        bookDao.getBooks()
    }

    Book getBook(int id) {
        bookDao.getBook(id)
    }

}

