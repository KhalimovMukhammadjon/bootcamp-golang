package migrations

create type enum gender('male','female')

create table if not exists movies(
    id integer primary key autoincrement not null,
    title varchar not null,
    year date not null,
    description varchar not null,
    image_url varchar not null,
    rating float not null,
    created_at timestamps default current_timestamp not null,
    updated_at timestamps default current_timestamp not null
);

create table if not exists actors(
    id integer primary key autoincrement not null,
    name varchar not null,
    image_url varchar not null,
    gender gender,
    created_at timestamps default current_timestamp not null,
    updated_at timestamps default current_timestamp not null
);

create table if not exists genres(
    id integer primary key autoincrement not null,
    name varchar not null,
    created_at timestamps default current_timestamp not null,
    updated_at timestamps default current_timestamp not null
);

create table if not exists movies_genres(
    id integer primary key autoincrement not null,
    movie_id integer not null,
    genre_id integer not null,
    created_at timestamps default current_timestamp not null,
    updated_at timestamps default current_timestamp not null,
    foreign key (movie_id) references movies(id),
    foreign key (genre_id) references genres(id)
);

create table if not exists movies_actors(
    id integer primary key autoincrement not null,
    movie_id integer not null,
    actor_id integer not null,
    created_at timestamps default current_timestamp not null,
    updated_at timestamps default current_timestamp not null,
    foreign key (movie_id) references movies(id),
    foreign key (actor_id) references actors(id)
);