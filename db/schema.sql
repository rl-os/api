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
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

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
-- Name: user_details; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_details (
    id integer NOT NULL,
    user_id integer NOT NULL,
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
    playstyle character varying[] DEFAULT '{}'::character varying[],
    playmode character varying DEFAULT ''::character varying NOT NULL,
    cover_url character varying DEFAULT ''::character varying,
    max_blocks integer DEFAULT 50 NOT NULL,
    max_friends integer DEFAULT 100
);


--
-- Name: user_details_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_details_id_seq OWNED BY public.user_details.id;


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
    avatar_url character varying DEFAULT ''::character varying NOT NULL,
    country_code character varying DEFAULT ''::character varying NOT NULL,
    default_group character varying DEFAULT 'osu'::character varying NOT NULL
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
-- Name: countries id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.countries ALTER COLUMN id SET DEFAULT nextval('public.countries_id_seq'::regclass);


--
-- Name: oauth_client id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_client ALTER COLUMN id SET DEFAULT nextval('public.oauth_client_id_seq'::regclass);


--
-- Name: oauth_token id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_token ALTER COLUMN id SET DEFAULT nextval('public.oauth_token_id_seq'::regclass);


--
-- Name: user_details id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_details ALTER COLUMN id SET DEFAULT nextval('public.user_details_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: countries countries_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT countries_pk PRIMARY KEY (id);


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
-- Name: user_details user_details_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT user_details_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


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
-- Name: oauth_client_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX oauth_client_id_uindex ON public.oauth_client USING btree (id);


--
-- Name: oauth_token_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX oauth_token_id_uindex ON public.oauth_token USING btree (id);


--
-- Name: user_details_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_details_id_uindex ON public.user_details USING btree (id);


--
-- Name: user_details_user_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_details_user_id_uindex ON public.user_details USING btree (user_id);


--
-- Name: users_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_id_uindex ON public.users USING btree (id);


--
-- Name: users_username_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX users_username_uindex ON public.users USING btree (username);


--
-- Name: user_details user_details_users_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_details
    ADD CONSTRAINT user_details_users_id_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


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
    ('20191219114442');
