--- name: CreateUser
INSERT INTO "users" ("email", "password")
VALUES ($1, $2)
RETURNING "id", "email", "password", "created_at", "updated_at";

--- name: GetUserByEmail
SELECT * FROM "users" WHERE "email" = $1;

--- name: UpdatePassword
UPDATE "users" SET "password" = $1 WHERE "id" = $2;
