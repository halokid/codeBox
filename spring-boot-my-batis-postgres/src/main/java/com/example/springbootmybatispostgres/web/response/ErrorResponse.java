package com.example.springbootmybatispostgres.web.response;

import lombok.*;

import java.time.Instant;

/**
 * <h2>ErrorResponse</h2>
 *
 * @author aek
 * <p>
 * Description:
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
@Builder
public class ErrorResponse {

    /**
     * error code
     */
    private String code;
    /**
     * short error message
     */
    private String message;

    /**
     * error cause timestamp
     */
    private Instant timestamp;
}
