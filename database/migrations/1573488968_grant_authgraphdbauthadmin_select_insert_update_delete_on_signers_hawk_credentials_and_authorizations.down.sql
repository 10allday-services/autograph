-- 1573488968_grant_authgraphdbauthadmin_select_insert_update_delete_on_signers_hawk_credentials_and_authorizations.down.sql
REVOKE SELECT, INSERT, UPDATE, DELETE ON signers, hawk_credentials, authorizations TO myautographdbauthadmin;
REVOKE USAGE ON authorizations_id_seq TO myautographdbauthadmin;
