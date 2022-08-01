-- +migrate Up

--
-- accounts
--
create table accounts (
  id            integer             primary key,
  username      string              not null unique,
  password      string              not null,
  token         string              not null unique,
  created_at    datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),
  updated_at    datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),

  constraint chk_username check (username <> ''),
  constraint chk_password check (username <> ''),
  constraint chk_token check(token <> '')
);

--
-- identities
--
create table identities (
    id          integer             primary key,
    account_id  integer             constraint fk_accounts_identities references accounts on delete cascade,
    ziti_id     string              not null unique,
    active      boolean             not null,
    created_at  datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),
    updated_at  datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),

    constraint chk_ziti_id check (ziti_id <> '')
);

--
-- services
--
create table services (
  id            integer             primary key,
  account_id    integer             constraint fk_accounts_services references accounts on delete cascade,
  ziti_id       string              not null unique,
  endpoint      string,
  active        boolean             not null,
  created_at    datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),
  updated_at    datetime            not null default(strftime('%Y-%m-%d %H:%M:%f', 'now')),

  constraint chk_ziti_id check (ziti_id <> '')
);