CREATE TYPE "ussr_requestor" AS ENUM (
  'user',
  'sender'
);

CREATE TYPE "ussr_status" AS ENUM (
  'requested',
  'accepted',
  'rejected'
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY NOT NULL,
  "email" TEXT UNIQUE NOT NULL,
  "password" TEXT NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "refresh_tokens" (
  "id" serial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "user_agent" TEXT NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "senders" (
  "id" serial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "device_id" TEXT UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "user_senders" (
  "user_id" int NOT NULL,
  "sender_id" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "sender_id")
);

CREATE TABLE "user_sender_reqs" (
  "id" serial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "sender_id" int NOT NULL,
  "requestor" ussr_requestor NOT NULL,
  "status" ussr_status NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "user_api_keys" (
  "api_key" TEXT PRIMARY KEY NOT NULL,
  "user" int NOT NULL,
  "api_secret" TEXT NOT NULL,
  "expiresAfter" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "batches" (
  "id" serial PRIMARY KEY NOT NULL,
  "queued_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "messages" (
  "id" serial PRIMARY KEY NOT NULL,
  "to" TEXT NOT NULL,
  "content" TEXT NOT NULL,
  "batch_id" int,
  "sent_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "user_id" int NOT NULL,
  "sender_id" int NOT NULL
);

ALTER TABLE "refresh_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "senders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_senders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_senders" ADD FOREIGN KEY ("sender_id") REFERENCES "senders" ("id");

ALTER TABLE "user_sender_reqs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_sender_reqs" ADD FOREIGN KEY ("sender_id") REFERENCES "senders" ("id");

ALTER TABLE "user_api_keys" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("batch_id") REFERENCES "batches" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("sender_id") REFERENCES "senders" ("id");
