create table if not exists "users" (
	"id" serial not null primary key,
	"name" varchar(50) not null,
	"username" varchar(30) not null,
	"email" varchar not null,
	"password" varchar not null,
	"created_at" timestamp not null default current_timestamp,
	"updated_at" timestamp,
	"deleted_at" timestamp
);