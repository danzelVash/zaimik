create table users
(
    id           serial       not null unique,
    last_name    varchar(250)          default '',
    first_name   varchar(250)          default '',
    sur_name     varchar(250)          default '',
    email        varchar(250) not null unique,
    phone_number varchar(30)  not null default '',
    city         varchar(250)          default '',
    is_admin     boolean      not null default false
);

INSERT INTO users (last_name, first_name, sur_name, email, phone_number, city, is_admin)
VALUES ('admin', 'admin', 'admin', 'admin', 'admin', 'admin', true);

create table loan_requests
(
    id       serial                                      not null unique,
    user_id  int references users (id) on delete cascade not null,
    amount   int                                         not null,
    duration int                                         not null
);

create table subscriptions
(
    id                          serial                                              not null unique,
    user_id                     int references users (id) on delete cascade not null unique,
    loan_id                     int references loan_requests (id) on delete cascade not null unique,
    request_date                date                                                not null,
    first_pay_time              timestamp,
    first_pay_success           boolean,
    second_pay_appointment_date date not null ,
    second_pay_time             timestamp,
    second_pay_success          boolean,
    expired_date                date
);

create table sessions
(
    id           serial                                      not null unique,
    user_id      int references users (id) on delete cascade not null,
    session      varchar(250)                                not null,
    expired_date date                                        not null
);

-- create or replace function delete_expired_sessions()
-- returns trigger as $$
-- begin
--     if old.expired_date < current_date then
--         delete from sessions where id = old.id;
--     end if;
--     return old;
-- END;
-- $$ LANGUAGE plpgsql;
--
-- CREATE TRIGGER trg_delete_expired_sessions
--     BEFORE DELETE ON sessions
--     for each row
--     execute function delete_expired_sessions();


create table companies
(
    id                   serial       not null unique,
    name                 varchar(250) not null,
    logo_name_on_s3      varchar(250) not null,
    link_on_company_site varchar(250) not null,
    max_loan_amount      int          not null default 2147483647,
    max_loan_duration    int          not null default 2147483647,
    min_loan_percent     int          not null,
    priority             int          not null default 0
);

create table reviews
(
    id             serial       not null unique,
    reviewer_name  varchar(250) not null,
    reviewer_phone varchar(250) not null,
    review         text         not null,
    moderated      boolean      not null default false
);
