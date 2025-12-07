-- Modify "users" table
ALTER TABLE "users" DROP COLUMN "username", ADD COLUMN "first_name" text NOT NULL, ADD COLUMN "last_name" text NOT NULL, ADD COLUMN "phone" text NOT NULL;
