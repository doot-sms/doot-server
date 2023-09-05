-- SQL SCHEMA

CREATE TYPE "ussr_requestor" AS ENUM (
  'user',
  'sender'
);

CREATE TYPE "ussr_status" AS ENUM (
  'requested',
  'accepted',
  'rejected'
);

CREATE TABLE "user" (
  "id" int PRIMARY KEY,
  "email" string UNIQUE,
  "password" string
);

CREATE TABLE "sender" (
  "id" int PRIMARY KEY,
  "user_id" int,
  "device_id" string UNIQUE
);

CREATE TABLE "user_sender" (
  "user_id" int,
  "sender_id" int,
  PRIMARY KEY ("user_id", "sender_id")
);

CREATE TABLE "user_sender_req" (
  "id" int PRIMARY KEY,
  "user_id" int,
  "sender_id" int,
  "requestor" ussr_requestor,
  "status" ussr_status
);

CREATE TABLE "user_api_key" (
  "api_key" string PRIMARY KEY,
  "user" int,
  "api_secret" string,
  "expiresAfter" timestamp
);

CREATE TABLE "batch" (
  "id" int PRIMARY KEY,
  "queued_at" timestamp
);

CREATE TABLE "message" (
  "to" string,
  "content" string,
  "batch_id" int,
  "sent_at" timestamp,
  "user_id" int,
  "sent_from" int
);

ALTER TABLE "sender" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_sender" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_sender" ADD FOREIGN KEY ("sender_id") REFERENCES "sender" ("id");

ALTER TABLE "user_sender_req" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_sender_req" ADD FOREIGN KEY ("sender_id") REFERENCES "sender" ("id");

ALTER TABLE "user_api_key" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("batch_id") REFERENCES "batch" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "message" ADD FOREIGN KEY ("sent_from") REFERENCES "sender" ("id");
