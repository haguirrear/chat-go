CREATE TABLE "users" (
    "id" serial PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NULL,
    "is_active" boolean NOT NULL DEFAULT TRUE,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NULL
)