package org.r00txx.dao

import org.r00txx.entity.Book
import org.springframework.stereotype.Repository

@Repository
interface IBookDao {
    List<Book> getBooks()
    Book getBook(int id)
}


