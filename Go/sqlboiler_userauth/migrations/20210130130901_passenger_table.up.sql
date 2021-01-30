CREATE TABLE passenger
(
    ID int primary key,
    created_at timestamp,
    updated_at timestamp,
    email varchar(255),
    password varchar(255),
    fullname varchar(255),
    phone varchar(255),
    cellphone varchar(255),
    is_cellphone_verified tinyint(1) not null,
    photo_url varchar(255),
    is_email_verified tinyint(1) not null,
    locale varchar(255),
    referral_code varchar(255),
    registeration_ip varchar(255),
    is_registered_with_google tinyint(1) not null,
    email_verification_code varchar(255) not null,
    cellphone_verification_code varchar(255) not null,
    is_blocked tinyint(1) not null,
    adjust_fingerprint varchar(255) not null,
    comapny_id int(11) not null
);