CREATE TABLE user_table (
    idx             INT(11) UNSIGNED    NOT NULL    AUTO_INCREMENT,
    username        VARCHAR(12)         NOT NULL,
    email           VARCHAR(320)        NOT NULL,
    password        VARCHAR(320)        NOT NULL,
    PRIMARY KEY(idx)
)

