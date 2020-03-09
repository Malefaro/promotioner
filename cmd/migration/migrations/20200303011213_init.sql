-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
    user_id bigserial primary key,
    email varchar(255) unique not null,
    password varchar(64) not null,
    is_superuser bool not null default false
);

CREATE TYPE food_place_type_enum as enum('cafe', 'bar', 'restaurant', 'club', 'fast_food');
CREATE TYPE food_type_enum as enum('sushi', 'pizza', 'burgers', 'italian', 'asian', 'eastern', 'russian');

CREATE TABLE food_place (
    food_place_id bigserial primary key,
    name varchar(255) not null,
    average_price bigint not null,
    description text not null default '',
    food_place_type food_place_type_enum not null,
    food_type food_type_enum not null,
    rating real not null default 0.0
);

CREATE TABLE food_place_addresses (
    food_place_address_id bigserial primary key,
    latitude double precision not null,
    longitude double precision not null,
    address varchar(500) not null,
    food_place_id bigint references food_place(food_place_id) on delete CASCADE not null
);

CREATE TABLE food_place_likes (
    food_place_like_id bigserial primary key,
    rating smallint not null,
    food_place_id bigint references food_place(food_place_id) on delete CASCADE not null,
    user_id bigint references users(user_id) on delete CASCADE not null
);

CREATE TYPE promotion_type_enum as enum ('discount', 'gift', 'cumulative_bonus', 'several_products_buy_bonus');

CREATE TABLE promotions (
    promotion_id bigserial primary key,
    start_date date null,
    end_date date null,
    label_text varchar(255) not null,
    description text not null default '',
    banner varchar(255) not null default '',
    rating real not null default 0.0,
    promotion_type promotion_type_enum not null,

    -- creator_id bigint references promotion_creators(promotion_creator_id) on delete CASCADE not null,
    food_place_id bigint references food_place(food_place_id) on delete CASCADE not null
);

CREATE TABLE promotion_creators (
    promotion_creator_id bigserial primary key,
    is_custom bool,

    user_id bigint references users(user_id) on delete CASCADE not null,
    promotion_id bigint references promotions(promotion_id) on delete CASCADE not null unique
);

-- ALTER TABLE promotions ADD foreign key (creator_id) references promotion_creators(promotion_creator_id) on delete CASCADE;

CREATE TABLE promotion_addresses (
    promotion_address_id bigserial primary key,

    promotion_id bigint references promotions(promotion_id) on delete CASCADE not null,
    food_place_address_id bigint references food_place_addresses(food_place_address_id) on delete CASCADE not null
);

CREATE TABLE promotion_likes (
    promotion_like_id bigserial primary key,
    rating smallint not null,

    user_id bigint references users(user_id) on delete CASCADE not null,
    promotion_id bigint references promotions(promotion_id) on delete CASCADE not null
);

CREATE TABLE permissions (
    permission_id bigserial primary key,
    name varchar(255) not null,
    priority smallint not null
);

CREATE TABLE permission_groups (
    permission_group_id bigserial primary key,
    name varchar(255) not null
);

CREATE TABLE permission_group_permissions (
    permission_group_permission_id bigserial primary key,

    permission_id bigint references permissions(permission_id) on delete CASCADE not null,
    permission_group_id bigint references permission_groups(permission_group_id) on delete CASCADE not null
);

CREATE TABLE food_place_permission_groups (
    food_place_permission_group_id bigserial primary key,

    food_place_id bigint references food_place(food_place_id) on delete CASCADE not null,
    permission_group_id bigint references permission_groups(permission_group_id) on delete CASCADE not null
);

CREATE TABLE user_food_place_permission_groups (
    user_food_place_permission_group_id bigserial primary key,

    food_place_permission_group_id bigint references food_place_permission_groups(food_place_permission_group_id) on delete CASCADE not null,
    user_id bigint references users(user_id) on delete CASCADE not null
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE user_food_place_permission_groups;
DROP TABLE food_place_permission_groups;
DROP TABLE permission_group_permissions;
DROP TABLE permission_groups;
DROP TABLE permissions;
DROP TABLE promotion_likes;
DROP TABLE promotion_addresses;
DROP TABLE promotion_creators;
DROP TABLE promotions;
DROP TYPE promotion_type_enum;
DROP TABLE food_place_likes;
DROP TABLE food_place_addresses;
DROP TABLE food_place;
DROP TYPE food_type_enum;
DROP TYPE food_place_type_enum;
DROP TABLE users;

