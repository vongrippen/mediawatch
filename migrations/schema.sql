--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2 (Debian 12.2-1.pgdg90+1)
-- Dumped by pg_dump version 12.2 (Ubuntu 12.2-2.pgdg18.04+1)

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
-- Name: media_files; Type: TABLE; Schema: public; Owner: gomediawatch
--

CREATE TABLE public.media_files (
    id uuid NOT NULL,
    size integer NOT NULL,
    pathname character varying(255) NOT NULL,
    filename character varying(255) NOT NULL,
    content_type character varying(255) NOT NULL,
    video_codec character varying(255) NOT NULL,
    audio_codec character varying(255) NOT NULL,
    duration_seconds numeric NOT NULL,
    container_format character varying(255) NOT NULL,
    width integer NOT NULL,
    height integer NOT NULL,
    show character varying(255) NOT NULL,
    season character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.media_files OWNER TO gomediawatch;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: gomediawatch
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO gomediawatch;

--
-- Name: media_files media_files_pkey; Type: CONSTRAINT; Schema: public; Owner: gomediawatch
--

ALTER TABLE ONLY public.media_files
    ADD CONSTRAINT media_files_pkey PRIMARY KEY (id);


--
-- Name: media_files_audio_codec_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_audio_codec_idx ON public.media_files USING btree (audio_codec);


--
-- Name: media_files_container_format_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_container_format_idx ON public.media_files USING btree (container_format);


--
-- Name: media_files_content_type_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_content_type_idx ON public.media_files USING btree (content_type);


--
-- Name: media_files_duration_seconds_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_duration_seconds_idx ON public.media_files USING btree (duration_seconds);


--
-- Name: media_files_height_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_height_idx ON public.media_files USING btree (height);


--
-- Name: media_files_season_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_season_idx ON public.media_files USING btree (season);


--
-- Name: media_files_show_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_show_idx ON public.media_files USING btree (show);


--
-- Name: media_files_size_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_size_idx ON public.media_files USING btree (size);


--
-- Name: media_files_video_codec_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_video_codec_idx ON public.media_files USING btree (video_codec);


--
-- Name: media_files_width_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE INDEX media_files_width_idx ON public.media_files USING btree (width);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: gomediawatch
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: admin_patroni
--

GRANT CREATE ON SCHEMA public TO gomediawatch;


--
-- Name: DEFAULT PRIVILEGES FOR SEQUENCES; Type: DEFAULT ACL; Schema: public; Owner: vgplex_6018
--

ALTER DEFAULT PRIVILEGES FOR ROLE vgplex_6018 IN SCHEMA public REVOKE ALL ON SEQUENCES  FROM vgplex_6018;
ALTER DEFAULT PRIVILEGES FOR ROLE vgplex_6018 IN SCHEMA public GRANT ALL ON SEQUENCES  TO gomediawatch;


--
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: public; Owner: vgplex_6018
--

ALTER DEFAULT PRIVILEGES FOR ROLE vgplex_6018 IN SCHEMA public REVOKE ALL ON TABLES  FROM vgplex_6018;
ALTER DEFAULT PRIVILEGES FOR ROLE vgplex_6018 IN SCHEMA public GRANT ALL ON TABLES  TO gomediawatch;


--
-- PostgreSQL database dump complete
--

