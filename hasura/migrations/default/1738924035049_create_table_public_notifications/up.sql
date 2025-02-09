CREATE TABLE "public"."notifications" ("id" serial NOT NULL, "userId" integer NOT NULL, "message" varchar NOT NULL, "is_read" boolean NOT NULL DEFAULT false, "created_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , FOREIGN KEY ("userId") REFERENCES "public"."users"("id") ON UPDATE cascade ON DELETE cascade);
