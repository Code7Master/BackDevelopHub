CREATE TABLE question_table (
    idx             INT(11) UNSIGNED    NOT NULL    AUTO_INCREMENT,
    writer          VARCHAR(12)         NOT NULL,
    title           VARCHAR(100)        NOT NULL,
    content         TEXT                NOT NULL,
    vote_cnt        INT(11) UNSIGNED    NOT NULL,
    answer_cnt    INT(11) UNSIGNED      NOT NULL,
    view_cnt        INT(11) UNSIGNED    NOT NULL,
    PRIMARY KEY(idx)
)   