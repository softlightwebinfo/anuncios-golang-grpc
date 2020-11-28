create table roles
(
    role varchar(100) not null
        constraint roles_pk
            primary key
);

alter table roles owner to postgres;

create table users
(
    id bigserial not null
        constraint users_pk
            primary key,
    user_name varchar(100) not null,
    email varchar(200) not null,
    password varchar(255) not null,
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null,
    deleted_at timestamp with time zone default now() not null,
    fk_role varchar(100)
        constraint users_roles_role_fk
            references roles
            on update cascade on delete set null
);

alter table users owner to postgres;

create table permissions
(
    token varchar(100) not null
        constraint permissions_pk
            primary key,
    description text not null
);

alter table permissions owner to postgres;

create table roles_permissions
(
    fk_role varchar(100) not null
        constraint roles_permissions_roles_role_fk
            references roles
            on update cascade on delete restrict,
    fk_permission varchar(100) not null
        constraint roles_permissions_permissions_token_fk
            references permissions
            on update cascade on delete restrict,
    constraint roles_permissions_pk
        primary key (fk_role, fk_permission)
);

alter table roles_permissions owner to postgres;

create index roles_permissions_fk_role_fk_permission_index
    on roles_permissions (fk_role, fk_permission);

create table articles
(
    id bigserial not null
        constraint articles_pk
            primary key,
    title varchar(150) not null,
    description text not null,
    fk_user bigint not null
        constraint articles_users_id_fk
            references users
            on update cascade on delete restrict,
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null,
    deleted_at timestamp with time zone
);

alter table articles owner to postgres;

create index articles_fk_user_index
    on articles (fk_user);

create table categories
(
    id serial not null
        constraint categories_pk
            primary key,
    category varchar(100) not null,
    parent integer
        constraint categories_categories_id_fk
            references categories
            on update cascade on delete set null
);

alter table categories owner to postgres;

create table articles_images
(
    fk_article bigint not null
        constraint articles_images_articles_id_fk
            references articles
            on update cascade on delete restrict,
    image varchar(100) not null,
    constraint articles_images_pk
        primary key (fk_article, image)
);

alter table articles_images owner to postgres;

create index articles_images_fk_article_index
    on articles_images (fk_article);

create table users_login
(
    fk_user integer not null
        constraint users_login_users_id_fk
            references users
            on update cascade on delete restrict,
    created_at timestamp with time zone default now() not null,
    ip varchar(20),
    constraint users_login_pk
        primary key (fk_user, created_at)
);

alter table users_login owner to postgres;

create index users_login_fk_user_index
    on users_login (fk_user);

create table notifications
(
    id bigserial not null
        constraint notifications_pk
            primary key,
    fk_user integer not null
        constraint notifications_users_id_fk
            references users
            on update cascade on delete restrict,
    title varchar(100) not null,
    description text not null,
    created_at timestamp with time zone default now() not null
);

alter table notifications owner to postgres;

create index notifications_fk_user_index
    on notifications (fk_user);

create table users_notifications
(
    fk_notification bigint not null
        constraint users_notifications_notifications_id_fk
            references notifications
            on update cascade on delete restrict,
    fk_user integer not null
        constraint users_notifications_users_id_fk
            references users
            on update cascade on delete restrict,
    created_at timestamp with time zone default now() not null,
    constraint users_notifications_pk
        primary key (fk_notification, fk_user)
);

alter table users_notifications owner to postgres;

