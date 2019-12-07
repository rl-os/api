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
-- Name: oauth_client id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_client ALTER COLUMN id SET DEFAULT nextval('public.oauth_client_id_seq'::regclass);


--
-- Name: oauth_token id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_token ALTER COLUMN id SET DEFAULT nextval('public.oauth_token_id_seq'::regclass);


--
-- Name: oauth_client oauth_client_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.oauth_client
    ADD CONSTRAINT oauth_client_pk PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: oauth_client_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX oauth_client_id_uindex ON public.oauth_client USING btree (id);


--
-- Name: oauth_token_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX oauth_token_id_uindex ON public.oauth_token USING btree (id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20191202100404'),
    ('20191204172445'),
    ('20191207083235');
