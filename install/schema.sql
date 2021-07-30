CREATE TABLE IF NOT EXISTS "settings"(
key text primary key not null,
value text not null,
type text not null
);
CREATE TABLE `admins` (`id` integer,`password` text,`salt` text,`email` text UNIQUE,`description` text,`avatar_url` text,PRIMARY KEY (`id`));
