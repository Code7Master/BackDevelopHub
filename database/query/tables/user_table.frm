CREATE TABLE user_table (
    idx             INT(11) UNSIGNED    NOT NULL    AUTO_INCREMENT,
    username        VARCHAR(12)         NOT NULL,
    email           VARCHAR(320)        NOT NULL,
    password        VARBINARY(700)      NOT NULL,
    create_date     DATE                NOT NULL,
    PRIMARY KEY(idx)
)
