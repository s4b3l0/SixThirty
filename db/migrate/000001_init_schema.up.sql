CREATE SCHEMA FIVE_THIRTY;

CREATE TABLE FIVE_THIRTY.CREDENTIAL(
                                       USERNAME VARCHAR(50) UNIQUE PRIMARY KEY ,
                                       EMAIL VARCHAR(50),
                                       PASSWORD_HASH VARCHAR(255)
);

CREATE TABLE FIVE_THIRTY.SESSION(
  TOKEN VARCHAR(255),
  EFFECTIVE_FROM timestamp,
  EFFECTIVE_TO timestamp
);