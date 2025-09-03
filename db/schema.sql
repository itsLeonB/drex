--
-- PostgreSQL database dump
--

\restrict 0qVZcZE3U2r3C4eNRR9OhX0pqgA2CdoQi5AMeSBEngQ7OD4dWEXbbcJKFzoMbk3

-- Dumped from database version 17.5 (1b53132)
-- Dumped by pg_dump version 17.6 (Ubuntu 17.6-1.pgdg24.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE IF EXISTS ONLY public.debt_transactions DROP CONSTRAINT IF EXISTS debt_transactions_transfer_method_id_fkey;
DROP INDEX IF EXISTS public.debt_transactions_transfer_method_id_idx;
DROP INDEX IF EXISTS public.debt_transactions_lender_profile_id_idx;
DROP INDEX IF EXISTS public.debt_transactions_created_at_idx;
DROP INDEX IF EXISTS public.debt_transactions_borrower_profile_id_idx;
ALTER TABLE IF EXISTS ONLY public.transfer_methods DROP CONSTRAINT IF EXISTS transfer_methods_pkey;
ALTER TABLE IF EXISTS ONLY public.debt_transactions DROP CONSTRAINT IF EXISTS debt_transactions_pkey;
DROP TABLE IF EXISTS public.transfer_methods;
DROP TABLE IF EXISTS public.debt_transactions;
DROP TYPE IF EXISTS public.debt_transaction_type;
DROP TYPE IF EXISTS public.debt_transaction_action;
--
-- Name: debt_transaction_action; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.debt_transaction_action AS ENUM (
    'LEND',
    'BORROW',
    'RECEIVE',
    'RETURN'
);


--
-- Name: debt_transaction_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.debt_transaction_type AS ENUM (
    'LEND',
    'REPAY'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: debt_transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.debt_transactions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    lender_profile_id uuid NOT NULL,
    borrower_profile_id uuid NOT NULL,
    type public.debt_transaction_type NOT NULL,
    action public.debt_transaction_action NOT NULL,
    amount numeric(20,2) NOT NULL,
    transfer_method_id uuid NOT NULL,
    description text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: transfer_methods; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transfer_methods (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    display text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);


--
-- Name: debt_transactions debt_transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.debt_transactions
    ADD CONSTRAINT debt_transactions_pkey PRIMARY KEY (id);


--
-- Name: transfer_methods transfer_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transfer_methods
    ADD CONSTRAINT transfer_methods_pkey PRIMARY KEY (id);


--
-- Name: debt_transactions_borrower_profile_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX debt_transactions_borrower_profile_id_idx ON public.debt_transactions USING btree (borrower_profile_id);


--
-- Name: debt_transactions_created_at_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX debt_transactions_created_at_idx ON public.debt_transactions USING btree (created_at);


--
-- Name: debt_transactions_lender_profile_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX debt_transactions_lender_profile_id_idx ON public.debt_transactions USING btree (lender_profile_id);


--
-- Name: debt_transactions_transfer_method_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX debt_transactions_transfer_method_id_idx ON public.debt_transactions USING btree (transfer_method_id);


--
-- Name: debt_transactions debt_transactions_transfer_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.debt_transactions
    ADD CONSTRAINT debt_transactions_transfer_method_id_fkey FOREIGN KEY (transfer_method_id) REFERENCES public.transfer_methods(id);


--
-- PostgreSQL database dump complete
--

\unrestrict 0qVZcZE3U2r3C4eNRR9OhX0pqgA2CdoQi5AMeSBEngQ7OD4dWEXbbcJKFzoMbk3

