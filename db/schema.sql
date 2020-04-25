SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: check_online(timestamp without time zone); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.check_online(val timestamp without time zone) RETURNS boolean
    LANGUAGE plpgsql IMMUTABLE
    AS $$
BEGIN
    RETURN (val > (CURRENT_TIMESTAMP - (10 ||' minutes')::interval));
END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: achievements; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.achievements (
    id integer NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    enabled boolean DEFAULT true NOT NULL,
    "grouping" character varying NOT NULL,
    image character varying,
    mode character varying DEFAULT 'osu'::character varying,
    quest_instructions character varying,
    slug character varying NOT NULL
);


--
-- Name: achievements_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.achievements_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: achievements_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.achievements_id_seq OWNED BY public.achievements.id;


--
-- Name: beatmap_set; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.beatmap_set (
    id integer NOT NULL,
    last_checked timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    title character varying DEFAULT ''::character varying NOT NULL,
    artist character varying DEFAULT ''::character varying NOT NULL,
    play_count integer DEFAULT 0 NOT NULL,
    favourite_count integer DEFAULT 0 NOT NULL,
    has_favourited boolean DEFAULT false,
    submitted_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_updated timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ranked_date timestamp with time zone,
    creator character varying DEFAULT ''::character varying NOT NULL,
    user_id integer NOT NULL,
    bpm integer DEFAULT 150 NOT NULL,
    source character varying NOT NULL,
    covers json DEFAULT json_build_object() NOT NULL,
    preview_url character varying NOT NULL,
    tags character varying NOT NULL,
    video boolean DEFAULT false NOT NULL,
    storyboard boolean DEFAULT false NOT NULL,
    ranked integer DEFAULT 0 NOT NULL,
    status character varying NOT NULL,
    is_scoreable boolean DEFAULT true NOT NULL,
    discussion_enabled boolean DEFAULT true NOT NULL,
    discussion_locked boolean DEFAULT false NOT NULL,
    can_be_hyped boolean DEFAULT true NOT NULL,
    availability json DEFAULT json_build_object() NOT NULL,
    hype json DEFAULT json_build_object() NOT NULL,
    nominations json DEFAULT json_build_object() NOT NULL,
    legacy_thread_url character varying DEFAULT ''::character varying NOT NULL,
    description json DEFAULT json_build_object('description', '') NOT NULL,
    genre json DEFAULT json_build_object('id', 1, 'name', 'None'),
    language json DEFAULT json_build_object('id', 1, 'name', 'None') NOT NULL,
    "user" json DEFAULT json_build_object() NOT NULL
);


--
-- Name: COLUMN beatmap_set.has_favourited; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.beatmap_set.has_favourited IS 'TODO THIS';


--
-- Name: COLUMN beatmap_set.user_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.beatmap_set.user_id IS 'user in original bancho';


--
-- Name: beatmap_set_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.beatmap_set_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: beatmap_set_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.beatmap_set_id_seq OWNED BY public.beatmap_set.id;


--
-- Name: beatmaps; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.beatmaps (
    id integer NOT NULL,
    beatmapset_id integer NOT NULL,
    mode character varying DEFAULT 'osu'::character varying NOT NULL,
    mode_int integer DEFAULT 0 NOT NULL,
    convert boolean DEFAULT false NOT NULL,
    difficulty_rating double precision DEFAULT 1.0 NOT NULL,
    version character varying DEFAULT ''::character varying NOT NULL,
    total_length integer DEFAULT 100 NOT NULL,
    hit_length integer DEFAULT 100,
    bpm integer DEFAULT 100 NOT NULL,
    cs integer DEFAULT 5 NOT NULL,
    drain integer DEFAULT 5 NOT NULL,
    accuracy integer DEFAULT 5 NOT NULL,
    ar integer DEFAULT 5 NOT NULL,
    playcount integer DEFAULT 0 NOT NULL,
    passcount integer DEFAULT 0 NOT NULL,
    count_circles integer DEFAULT 0 NOT NULL,
    count_sliders integer DEFAULT 0 NOT NULL,
    count_spinners integer DEFAULT 0 NOT NULL,
    count_total integer DEFAULT 0 NOT NULL,
    is_scoreable boolean DEFAULT true NOT NULL,
    last_updated timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ranked integer DEFAULT 0 NOT NULL,
    status character varying DEFAULT 'ranked'::character varying NOT NULL,
    url character varying DEFAULT ''::character varying NOT NULL,
    deleted_at timestamp with time zone,
    max_combo integer
);


--
-- Name: beatmaps_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.beatmaps_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: beatmaps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.beatmaps_id_seq OWNED BY public.beatmaps.id;


--
-- Name: channels; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.channels (
    id integer NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    type character varying NOT NULL,
    icon character varying,
    users integer[] DEFAULT '{}'::integer[] NOT NULL,
    active_users integer[] DEFAULT '{}'::integer[] NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: channels_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.channels_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: channels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.channels_id_seq OWNED BY public.channels.id;


--
-- Name: countries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.countries (
    id integer NOT NULL,
    code character varying NOT NULL,
    name character varying NOT NULL
);


--
-- Name: TABLE countries; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.countries IS 'contains all country codes and names ';


--
-- Name: countries_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.countries_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: countries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.countries_id_seq OWNED BY public.countries.id;


--
-- Name: message; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.message (
    id integer NOT NULL,
    sender_id integer NOT NULL,
    channel_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    content character varying NOT NULL,
    is_action boolean DEFAULT false NOT NULL
);


--
-- Name: message_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.message_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.message_id_seq OWNED BY public.message.id;


--
-- Name: oauth_client; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.oauth_client (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying NOT NULL,
    secret character varying NOT NULL,
    redirect character varying NOT NULL,
    revoked boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: oauth_client_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.oauth_client_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: oauth_client_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.oauth_client_id_seq OWNED BY public.oauth_client.id;


--
-- Name: oauth_token; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.oauth_token (
    id integer NOT NULL,
    user_id integer NOT NULL,
    access_token character varying NOT NULL,
    refresh_token character varying NOT NULL,
    scopes character varying NOT NULL,
    revoked boolean DEFAULT false NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    client_id integer DEFAULT 0 NOT NULL
);


--
-- Name: oauth_token_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.oauth_token_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: oauth_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.oauth_token_id_seq OWNED BY public.oauth_token.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: user_achievements; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_achievements (
    id integer NOT NULL,
    achievement_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: user_achievements_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_achievements_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_achievements_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_achievements_id_seq OWNED BY public.user_achievements.id;


--
-- Name: user_month_playcount; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_month_playcount (
    id integer NOT NULL,
    user_id integer NOT NULL,
    playcount integer NOT NULL,
    year_month character varying NOT NULL
);


--
-- Name: COLUMN user_month_playcount.year_month; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.user_month_playcount.year_month IS '{year}-{month}-01';


--
-- Name: user_month_playcount_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_month_playcount_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_month_playcount_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_month_playcount_id_seq OWNED BY public.user_month_playcount.id;


--
-- Name: user_relation; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_relation (
    id integer NOT NULL,
    user_id integer NOT NULL,
    target_id integer NOT NULL,
    friend boolean DEFAULT false NOT NULL
);


--
-- Name: user_relation_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_relation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_relation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_relation_id_seq OWNED BY public.user_relation.id;


--
-- Name: user_statistics; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_statistics (
    id integer NOT NULL,
    user_id integer NOT NULL,
    level_current integer DEFAULT 1 NOT NULL,
    level_progress integer DEFAULT 0 NOT NULL,
    pp double precision DEFAULT 0.0 NOT NULL,
    ranked_score integer DEFAULT 0 NOT NULL,
    hit_accuracy double precision DEFAULT 0.0 NOT NULL,
    play_count integer DEFAULT 0 NOT NULL,
    play_time integer DEFAULT 0 NOT NULL,
    total_score integer DEFAULT 0 NOT NULL,
    total_hits integer DEFAULT 0 NOT NULL,
    maximum_combo integer DEFAULT 0 NOT NULL,
    replays_watched_by_others integer DEFAULT 0 NOT NULL,
    is_ranked boolean DEFAULT true NOT NULL,
    grade_counts_ss integer DEFAULT 0 NOT NULL,
    grade_counts_ssh integer DEFAULT 0 NOT NULL,
    grade_counts_s integer DEFAULT 0 NOT NULL,
    grade_counts_sh integer DEFAULT 0 NOT NULL,
    grade_counts_a integer DEFAULT 0 NOT NULL
);


--
-- Name: user_statistics_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_statistics_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_statistics_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_statistics_id_seq OWNED BY public.user_statistics.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying NOT NULL,
    email character varying NOT NULL,
    password_hash character varying NOT NULL,
    last_visit timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    is_bot boolean DEFAULT false NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    is_supporter boolean DEFAULT false NOT NULL,
    has_supported boolean DEFAULT false NOT NULL,
    support_level integer DEFAULT 0 NOT NULL,
    pm_friends_only boolean DEFAULT false NOT NULL,
    avatar_url character varying DEFAULT 'https://301222.selcdn.ru/akasi/avatars/1.png'::character varying NOT NULL,
    country_code character varying DEFAULT '-'::character varying NOT NULL,
    default_group character varying DEFAULT 'osu'::character varying NOT NULL,
    can_moderate boolean DEFAULT false NOT NULL,
    interests character varying,
    occupation character varying DEFAULT ''::character varying NOT NULL,
    title character varying,
    location character varying,
    twitter character varying,
    lastfm character varying,
    skype character varying,
    website character varying,
    discord character varying,
    playstyle character varying[] DEFAULT '{}'::character varying[] NOT NULL,
    playmode character varying DEFAULT ''::character varying NOT NULL,
    cover_url character varying DEFAULT 'https://301222.selcdn.ru/akasi/bg/1.jpg'::character varying NOT NULL,
    max_blocks integer DEFAULT 50 NOT NULL,
    max_friends integer DEFAULT 100 NOT NULL,
    support_expired_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: achievements id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.achievements ALTER COLUMN id SET DEFAULT nextval('public.achievements_id_seq'::regclass);


--
-- Name: beatmap_set id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.beatmap_set ALTER COLUMN id SET DEFAULT nextval('public.beatmap_set_id_seq'::regclass);


--
-- Name: beatmaps id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.beatmaps ALTER COLUMN id SET DEFAULT nextval('public.beatmaps_id_seq'::regclass);


--
-- Name: channels id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.channels ALTER COLUMN id SET DEFAULT nextval('public.channels_id_seq'::regclass);


--
-- Name: countries id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.countries ALTER COLUMN id SET DEFAULT nextval('public.countries_id_seq'::regclass);


--
-- Name: message id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message ALTER COLUMN id SET DEFAULT nextval('public.message_id_seq'::regclass);


--
-- Name: oauth_client id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_client ALTER COLUMN id SET DEFAULT nextval('public.oauth_client_id_seq'::regclass);


--
-- Name: oauth_token id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_token ALTER COLUMN id SET DEFAULT nextval('public.oauth_token_id_seq'::regclass);


--
-- Name: user_achievements id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements ALTER COLUMN id SET DEFAULT nextval('public.user_achievements_id_seq'::regclass);


--
-- Name: user_month_playcount id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_month_playcount ALTER COLUMN id SET DEFAULT nextval('public.user_month_playcount_id_seq'::regclass);


--
-- Name: user_relation id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_relation ALTER COLUMN id SET DEFAULT nextval('public.user_relation_id_seq'::regclass);


--
-- Name: user_statistics id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_statistics ALTER COLUMN id SET DEFAULT nextval('public.user_statistics_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: achievements achievements_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.achievements
    ADD CONSTRAINT achievements_pk PRIMARY KEY (id);


--
-- Name: beatmap_set beatmap_set_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.beatmap_set
    ADD CONSTRAINT beatmap_set_pk PRIMARY KEY (id);


--
-- Name: beatmaps beatmaps_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.beatmaps
    ADD CONSTRAINT beatmaps_pk PRIMARY KEY (id);


--
-- Name: channels channels_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.channels
    ADD CONSTRAINT channels_pk PRIMARY KEY (id);


--
-- Name: countries countries_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_pk PRIMARY KEY (id);


--
-- Name: message message_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_pk PRIMARY KEY (id);


--
-- Name: oauth_client oauth_client_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_client
    ADD CONSTRAINT oauth_client_pk PRIMARY KEY (id);


--
-- Name: oauth_token oauth_token_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_token
    ADD CONSTRAINT oauth_token_pk PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: user_achievements user_achievements_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements
    ADD CONSTRAINT user_achievements_pk PRIMARY KEY (id);


--
-- Name: user_month_playcount user_month_playcount_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_month_playcount
    ADD CONSTRAINT user_month_playcount_pk PRIMARY KEY (id);


--
-- Name: user_relation user_relation_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_relation
    ADD CONSTRAINT user_relation_pk PRIMARY KEY (id);


--
-- Name: user_statistics user_statistics_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_statistics
    ADD CONSTRAINT user_statistics_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: achievements_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX achievements_id_uindex ON public.achievements USING btree (id);


--
-- Name: achievements_slug_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX achievements_slug_uindex ON public.achievements USING btree (slug);


--
-- Name: beatmap_set_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX beatmap_set_id_uindex ON public.beatmap_set USING btree (id);


--
-- Name: beatmaps_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX beatmaps_id_uindex ON public.beatmaps USING btree (id);


--
-- Name: channels_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX channels_id_index ON public.channels USING btree (id);


--
-- Name: countries_code_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX countries_code_uindex ON public.countries USING btree (code);


--
-- Name: countries_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX countries_id_uindex ON public.countries USING btree (id);


--
-- Name: countries_name_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX countries_name_uindex ON public.countries USING btree (name);


--
-- Name: message_channel_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX message_channel_id_index ON public.message USING btree (channel_id);


--
-- Name: message_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX message_id_uindex ON public.message USING btree (id);


--
-- Name: message_sender_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX message_sender_id_index ON public.message USING btree (sender_id);


--
-- Name: oauth_client_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX oauth_client_id_uindex ON public.oauth_client USING btree (id);


--
-- Name: oauth_token_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX oauth_token_id_uindex ON public.oauth_token USING btree (id);


--
-- Name: table_name_name_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX table_name_name_index ON public.channels USING btree (name);


--
-- Name: user_achievements_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_achievements_id_uindex ON public.user_achievements USING btree (id);


--
-- Name: user_achievements_user_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_achievements_user_id_index ON public.user_achievements USING btree (user_id);


--
-- Name: user_month_playcount_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_month_playcount_id_uindex ON public.user_month_playcount USING btree (id);


--
-- Name: user_relation_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_relation_id_uindex ON public.user_relation USING btree (id);


--
-- Name: user_relation_user_id_target_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_relation_user_id_target_id_index ON public.user_relation USING btree (user_id, target_id);


--
-- Name: user_statistics_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_statistics_id_uindex ON public.user_statistics USING btree (id);


--
-- Name: user_statistics_user_id_is_ranked_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_statistics_user_id_is_ranked_index ON public.user_statistics USING btree (user_id, is_ranked);


--
-- Name: users_email_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_email_uindex ON public.users USING btree (email);


--
-- Name: users_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_id_uindex ON public.users USING btree (id);


--
-- Name: users_username_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_username_uindex ON public.users USING btree (username);


--
-- Name: beatmaps beatmaps_beatmap_set_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.beatmaps
    ADD CONSTRAINT beatmaps_beatmap_set_id_fk FOREIGN KEY (beatmapset_id) REFERENCES public.beatmap_set(id);


--
-- Name: message message_channels_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_channels_id_fk FOREIGN KEY (channel_id) REFERENCES public.channels(id) ON DELETE CASCADE;


--
-- Name: message message_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_users_id_fk FOREIGN KEY (sender_id) REFERENCES public.users(id);


--
-- Name: user_achievements user_achievements_achievements_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements
    ADD CONSTRAINT user_achievements_achievements_fk FOREIGN KEY (achievement_id) REFERENCES public.achievements(id) ON DELETE CASCADE;


--
-- Name: user_achievements user_achievements_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_achievements
    ADD CONSTRAINT user_achievements_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_month_playcount user_month_playcount_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_month_playcount
    ADD CONSTRAINT user_month_playcount_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_relation user_relation_target_id_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_relation
    ADD CONSTRAINT user_relation_target_id_id_fk FOREIGN KEY (target_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_relation user_relation_users_id_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_relation
    ADD CONSTRAINT user_relation_users_id_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_statistics user_statistics_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_statistics
    ADD CONSTRAINT user_statistics_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20191202100404'),
    ('20191204172445'),
    ('20191207083235'),
    ('20191207174450'),
    ('20191207194321'),
    ('20191207194641'),
    ('20191219081947'),
    ('20200127085610'),
    ('20200127091253'),
    ('20200127092843'),
    ('20200127093220'),
    ('20200127094841'),
    ('20200128070351'),
    ('20200130072128'),
    ('20200130133637'),
    ('20200201131358'),
    ('20200201135712'),
    ('20200202081814'),
    ('20200224085752'),
    ('20200224101725'),
    ('20200224101740'),
    ('20200225094357'),
    ('20200313155053'),
    ('20200425094246');
