CREATE SCHEMA FIVE_THIRTY;

CREATE TABLE FIVE_THIRTY.CREDENTIAL(
                                       USER_NAME VARCHAR(50) UNIQUE PRIMARY KEY ,
                                       EMAIL VARCHAR(50),
                                       USER_PASSWORD VARCHAR(50)
);