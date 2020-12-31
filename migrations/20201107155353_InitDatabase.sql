-- migrate:up
create table if not exists oauth_token
(
    id            serial                              not null
        constraint oauth_token_pk
            primary key,
    user_id       integer                             not null,
    access_token  varchar                             not null,
    refresh_token varchar                             not null,
    scopes        varchar                             not null,
    revoked       boolean   default false             not null,
    expires_at    timestamp                           not null,
    created_at    timestamp default CURRENT_TIMESTAMP not null,
    client_id     integer   default 0                 not null
);

create unique index if not exists oauth_token_id_uindex
    on oauth_token (id);

create table if not exists oauth_client
(
    id         serial                              not null
        constraint oauth_client_pk
            primary key,
    user_id    integer                             not null,
    name       varchar                             not null,
    secret     varchar                             not null,
    redirect   varchar                             not null,
    revoked    boolean   default false             not null,
    created_at timestamp default CURRENT_TIMESTAMP not null
);

create unique index if not exists oauth_client_id_uindex
    on oauth_client (id);

create table if not exists users
(
    id                 serial                                                                                        not null
        constraint users_pk
            primary key,
    username           varchar                                                                                       not null,
    email              varchar                                                                                       not null,
    password_hash      varchar                                                                                       not null,
    last_visit         timestamp           default CURRENT_TIMESTAMP                                                 not null,
    created_at         timestamp           default CURRENT_TIMESTAMP                                                 not null,
    is_bot             boolean             default false                                                             not null,
    is_active          boolean             default true                                                              not null,
    is_supporter       boolean             default false                                                             not null,
    has_supported      boolean             default false                                                             not null,
    support_level      integer             default 0                                                                 not null,
    pm_friends_only    boolean             default false                                                             not null,
    avatar_url         varchar             default 'https://301222.selcdn.ru/akasi/avatars/1.png'::character varying not null,
    country_code       varchar             default '-'::character varying                                            not null,
    default_group      varchar             default 'osu'::character varying                                          not null,
    can_moderate       boolean             default false                                                             not null,
    interests          varchar,
    occupation         varchar             default ''::character varying                                             not null,
    title              varchar,
    location           varchar,
    twitter            varchar,
    lastfm             varchar,
    skype              varchar,
    website            varchar,
    discord            varchar,
    playstyle          character varying[] default '{}'::character varying[]                                         not null,
    playmode           varchar             default ''::character varying                                             not null,
    cover_url          varchar             default 'https://301222.selcdn.ru/akasi/bg/1.jpg'::character varying      not null,
    max_blocks         integer             default 50                                                                not null,
    max_friends        integer             default 100                                                               not null,
    support_expired_at timestamp           default CURRENT_TIMESTAMP                                                 not null,
    profile_colour     varchar             default ''::character varying
);

create unique index if not exists users_id_uindex
    on users (id);

create unique index if not exists users_username_uindex
    on users (username);

create unique index if not exists users_email_uindex
    on users (email);

create table if not exists countries
(
    id   serial  not null
        constraint countries_pk
            primary key,
    code varchar not null,
    name varchar not null
);

comment on table countries is 'contains all country codes and names ';

create unique index if not exists countries_code_uindex
    on countries (code);

create unique index if not exists countries_id_uindex
    on countries (id);

create unique index if not exists countries_name_uindex
    on countries (name);

create table if not exists channels
(
    id           serial                              not null
        constraint channels_pk
            primary key,
    name         varchar                             not null,
    description  varchar                             not null,
    type         varchar                             not null,
    icon         varchar,
    users        integer[] default '{}'::integer[]   not null,
    active_users integer[] default '{}'::integer[]   not null,
    created_at   timestamp default CURRENT_TIMESTAMP not null
);

create unique index if not exists channels_id_index
    on channels (id);

create index if not exists table_name_name_index
    on channels (name);

create table if not exists message
(
    id         serial                              not null
        constraint message_pk
            primary key,
    sender_id  integer                             not null
        constraint message_users_id_fk
            references users,
    channel_id integer                             not null
        constraint message_channels_id_fk
            references channels
            on delete cascade,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    content    varchar                             not null,
    is_action  boolean   default false             not null
);

create index if not exists message_channel_id_index
    on message (channel_id);

create unique index if not exists message_id_uindex
    on message (id);

create index if not exists message_sender_id_index
    on message (sender_id);

create table if not exists user_relation
(
    id        serial                not null
        constraint user_relation_pk
            primary key,
    user_id   integer               not null
        constraint user_relation_users_id_id_fk
            references users
            on delete cascade,
    target_id integer               not null
        constraint user_relation_target_id_id_fk
            references users
            on delete cascade,
    friend    boolean default false not null
);

create unique index if not exists user_relation_id_uindex
    on user_relation (id);

create index if not exists user_relation_user_id_target_id_index
    on user_relation (user_id, target_id);

create table if not exists user_month_playcount
(
    id         serial  not null
        constraint user_month_playcount_pk
            primary key,
    user_id    integer not null
        constraint user_month_playcount_users_id_fk
            references users
            on delete cascade,
    playcount  integer not null,
    year_month varchar not null
);

comment on column user_month_playcount.year_month is '{year}-{month}-01';

create unique index if not exists user_month_playcount_id_uindex
    on user_month_playcount (id);

create table if not exists achievements
(
    id                 serial               not null
        constraint achievements_pk
            primary key,
    name               varchar              not null,
    description        varchar              not null,
    enabled            boolean default true not null,
    grouping           varchar              not null,
    image              varchar,
    mode               varchar default 'osu'::character varying,
    quest_instructions varchar,
    slug               varchar              not null
);

create unique index if not exists achievements_id_uindex
    on achievements (id);

create unique index if not exists achievements_slug_uindex
    on achievements (slug);

create table if not exists user_achievements
(
    id             serial  not null
        constraint user_achievements_pk
            primary key,
    achievement_id integer not null
        constraint user_achievements_achievements_fk
            references achievements
            on delete cascade,
    user_id        integer not null
        constraint user_achievements_users_id_fk
            references users
            on delete cascade,
    created_at     timestamp default CURRENT_TIMESTAMP
);

create unique index if not exists user_achievements_id_uindex
    on user_achievements (id);

create index if not exists user_achievements_user_id_index
    on user_achievements (user_id);

create table if not exists user_statistics
(
    id                        serial                                            not null
        constraint user_statistics_pk
            primary key,
    user_id                   integer                                           not null
        constraint user_statistics_users_id_fk
            references users
            on delete cascade,
    level_current             integer          default 1                        not null,
    level_progress            integer          default 0                        not null,
    pp                        double precision default 0.0                      not null,
    ranked_score              integer          default 0                        not null,
    hit_accuracy              double precision default 0.0                      not null,
    play_count                integer          default 0                        not null,
    play_time                 integer          default 0                        not null,
    total_score               integer          default 0                        not null,
    total_hits                integer          default 0                        not null,
    maximum_combo             integer          default 0                        not null,
    replays_watched_by_others integer          default 0                        not null,
    is_ranked                 boolean          default true                     not null,
    grade_counts_ss           integer          default 0                        not null,
    grade_counts_ssh          integer          default 0                        not null,
    grade_counts_s            integer          default 0                        not null,
    grade_counts_sh           integer          default 0                        not null,
    grade_counts_a            integer          default 0                        not null,
    mode                      varchar          default 'std'::character varying not null
);

create unique index if not exists user_statistics_id_uindex
    on user_statistics (id);

create index if not exists user_statistics_user_id_is_ranked_index
    on user_statistics (user_id, is_ranked);

create index if not exists user_statistics_mode_index
    on user_statistics (mode);

create table if not exists beatmap_set
(
    id                 serial                                                                      not null
        constraint beatmap_set_pk
            primary key,
    last_checked       timestamp with time zone default CURRENT_TIMESTAMP                          not null,
    title              varchar                  default ''::character varying                      not null,
    artist             varchar                  default ''::character varying                      not null,
    play_count         integer                  default 0                                          not null,
    favourite_count    integer                  default 0                                          not null,
    submitted_date     timestamp with time zone default CURRENT_TIMESTAMP                          not null,
    last_updated       timestamp with time zone default CURRENT_TIMESTAMP                          not null,
    ranked_date        timestamp with time zone,
    creator            varchar                  default ''::character varying                      not null,
    user_id            integer                                                                     not null,
    bpm                integer   default 150                                        not null,
    source             varchar                                                      not null,
    covers             json      default json_build_object()                        not null,
    preview_url        varchar                                                      not null,
    tags               varchar                                                      not null,
    video              boolean   default false                                      not null,
    storyboard         boolean   default false                                      not null,
    ranked             integer   default 0                                          not null,
    status             varchar                                                      not null,
    is_scoreable       boolean   default true                                       not null,
    discussion_enabled boolean   default true                                       not null,
    discussion_locked  boolean   default false                                      not null,
    can_be_hyped       boolean   default true                                       not null,
    availability       json      default json_build_object()                        not null,
    hype               json      default json_build_object()                        not null,
    nominations        json      default json_build_object()                        not null,
    legacy_thread_url  varchar   default ''::character varying                      not null,
    description        json      default json_build_object('description', '')       not null,
    genre              json      default json_build_object('id', 1, 'name', 'None'),
    language           json      default json_build_object('id', 1, 'name', 'None') not null,
    "user"             json      default json_build_object()                        not null,
    ratings            integer[] default '{}'::integer[]                            not null
);

comment on column beatmap_set.user_id is 'user in original bancho';

create unique index if not exists beatmap_set_id_uindex
    on beatmap_set (id);

create table if not exists beatmaps
(
    id                serial                                                       not null
        constraint beatmaps_pk
            primary key,
    beatmapset_id     integer                                                      not null
        constraint beatmaps_beatmap_set_id_fk
            references beatmap_set,
    mode              varchar                  default 'osu'::character varying    not null,
    mode_int          integer                  default 0                           not null,
    convert           boolean                  default false                       not null,
    difficulty_rating double precision         default 1.0                         not null,
    version           varchar                  default ''::character varying       not null,
    total_length      integer                  default 100                         not null,
    hit_length        integer                  default 100,
    bpm               integer                  default 100                         not null,
    cs                integer                  default 5                           not null,
    drain             integer                  default 5                           not null,
    accuracy          integer                  default 5                           not null,
    ar                integer                  default 5                           not null,
    playcount         integer                  default 0                           not null,
    passcount         integer                  default 0                           not null,
    count_circles     integer                  default 0                           not null,
    count_sliders     integer                  default 0                           not null,
    count_spinners    integer                  default 0                           not null,
    count_total       integer                  default 0                           not null,
    is_scoreable      boolean                  default true                        not null,
    last_updated      timestamp with time zone default CURRENT_TIMESTAMP           not null,
    ranked            integer                  default 0                           not null,
    status            varchar                  default 'ranked'::character varying not null,
    url               varchar                  default ''::character varying       not null,
    deleted_at        timestamp with time zone,
    max_combo         integer
);

create unique index if not exists beatmaps_id_uindex
    on beatmaps (id);

create table if not exists user_beatmapset_favourite
(
    id            serial                                             not null
        constraint favouritemaps_pk
            primary key,
    beatmapset_id integer                                            not null,
    user_id       integer                                            not null,
    created_at    timestamp with time zone default CURRENT_TIMESTAMP not null
);

create unique index if not exists favouritemaps_id_uindex
    on user_beatmapset_favourite (id);

create unique index if not exists favouritemaps_user_record_index
    on user_beatmapset_favourite (user_id, beatmapset_id);

create table if not exists user_performance_ranks
(
    id      serial                                     not null
        constraint user_performance_ranks_pk
            primary key,
    mode    varchar   default 'std'::character varying not null,
    data    integer[] default '{}'::integer[]          not null,
    user_id integer   default 0                        not null
        constraint user_performance_ranks_users_id_fk
            references users
);

create index if not exists user_performance_ranks_mode_index
    on user_performance_ranks (mode);

create unique index if not exists user_performance_ranks_id_uindex
    on user_performance_ranks (id);

create table if not exists schema_migrations
(
    version varchar(255) not null
        constraint schema_migrations_pkey
            primary key
);

create table if not exists beatmap_failtimes
(
    beatmap_id integer   not null
        constraint beatmap_failtimes_pk
            primary key
        constraint beatmap_failtimes_beatmaps_id_fk
            references beatmaps,
    fail       integer[] not null,
    exit       integer[] not null
);

create unique index if not exists beatmap_failtimes_beatmap_id_uindex
    on beatmap_failtimes (beatmap_id);


-- migrate:down
drop table schema_migrations;

drop table oauth_token;

drop table oauth_client;

drop table countries;

drop table message;

drop table channels;

drop table user_relation;

drop table user_month_playcount;

drop table user_achievements;

drop table achievements;

drop table user_statistics;

drop table beatmaps;

drop table beatmap_set;

drop table user_beatmapset_favourite;

drop table user_performance_ranks;

drop table users;

drop table beatmap_failtimes;


