package com.example.springbootmybatispostgres.entities;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.validation.constraints.NotBlank;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Author {

    private Long id;

    @NotBlank
    private String lastname;

    private String firstname;

}