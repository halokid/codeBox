package com.example.springbootmybatispostgres.common.exception;

/**
 * trigger for bad request exception
 */
public class BadRequestException extends ServiceException {
    public BadRequestException() {
        super();
    }

    public BadRequestException(String message) {
        super(message);
    }

    public BadRequestException(String message, Throwable throwable) {
        super(message, throwable);
    }
}
