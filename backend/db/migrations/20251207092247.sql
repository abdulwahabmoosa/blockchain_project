-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "username" SET NOT NULL, ALTER COLUMN "email" SET NOT NULL, ALTER COLUMN "password_hash" SET NOT NULL, ADD CONSTRAINT "uni_users_email" UNIQUE ("email");
