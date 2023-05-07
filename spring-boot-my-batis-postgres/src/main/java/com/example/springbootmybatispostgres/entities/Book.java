package com.example.springbootmybatispostgres.entities;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Size;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Book {

    private Long id;

    @NotBlank
    @Size(max = 100)
    private String title;

    private int page;

    private String isbn;

    private String description;

    private double price;
}