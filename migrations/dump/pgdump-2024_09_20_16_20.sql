--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8
-- Dumped by pg_dump version 14.8

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO user;

--
-- Name: sku; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.sku (
    sku character varying(180) NOT NULL,
    name character varying(180) NOT NULL,
    size character varying(180) NOT NULL
);


ALTER TABLE public.sku OWNER TO user;

--
-- Name: sku_store; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.sku_store (
    id integer NOT NULL,
    sku character varying(180) NOT NULL,
    store_id integer NOT NULL,
    quantity integer NOT NULL,
    reserved integer NOT NULL,
    CONSTRAINT sku_store_quantity_check CHECK ((quantity >= 0)),
    CONSTRAINT sku_store_reserved_check CHECK ((reserved >= 0))
);


ALTER TABLE public.sku_store OWNER TO user;

--
-- Name: store; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.store (
    id integer NOT NULL,
    name character varying(180) NOT NULL,
    is_available boolean NOT NULL
);


ALTER TABLE public.store OWNER TO user;

--
-- Name: stores_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.stores_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.stores_id_seq OWNER TO user;

--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: sku; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.sku (sku, name, size) FROM stdin;
NikeClassicM	T-shirt nike classic	M
NikeClassicL	T-shirt nike classic	L
AdidasPoloM	T-shirt adidas polo	M
ShoesAdidas45	Shoes Adidas	45
\.


--
-- Data for Name: sku_store; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.sku_store (id, sku, store_id, quantity, reserved) FROM stdin;
2	NikeClassicM	2	5	1
3	AdidasPoloM	1	3	0
1	NikeClassicM	1	10	0
\.


--
-- Data for Name: store; Type: TABLE DATA; Schema: public; Owner: user
--

COPY public.store (id, name, is_available) FROM stdin;
3	thirdStore	f
2	secondStore	f
1	mainStore	t
\.


--
-- Name: stores_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.stores_id_seq', 1, false);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sku sku_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.sku
    ADD CONSTRAINT sku_pkey PRIMARY KEY (sku);


--
-- Name: sku_store sku_store_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.sku_store
    ADD CONSTRAINT sku_store_pkey PRIMARY KEY (id);


--
-- Name: sku_store sku_store_sku_store_id_key; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.sku_store
    ADD CONSTRAINT sku_store_sku_store_id_key UNIQUE (sku, store_id);


--
-- Name: store store_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.store
    ADD CONSTRAINT store_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

