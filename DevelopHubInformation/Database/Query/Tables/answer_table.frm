CREATE TABLE answer_table (
    idx             INT(11) UNSIGNED    NOT NULL    AUTO_INCREMENT,
    question_idx    INT(11) UNSIGNED    NOT NULL,
    writer          VARCHAR(12)         NOT NULL,
    content         TEXT                NOT NULL,
    vote_cnt        INT(11) UNSIGNED    NOT NULL,
    PRIMARY KEY(idx)
)