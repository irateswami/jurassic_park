-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists dinosaurs (
    id text not null primary key,
    name text,
    species text,
    cage text,
    FOREIGN KEY(cage) REFERENCES cages(id)
);

create table if not exists cages (
    id text not null primary key,
    species text,
    max_capacity integer not null,
    carnivore boolean,
    active boolean not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table dinosaurs;
drop table cages;
-- +goose StatementEnd
