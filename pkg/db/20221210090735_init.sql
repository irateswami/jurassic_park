-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table dinosaurs (
    id integer not null primary key,
    species text,
    name text,
    herb_or_carn integer,
    cage integer,
    FOREIGN KEY(cage) REFERENCES cages(id)
);

create table cages (
    id integer not null primary key,
    species text,
    capacity integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table dinosaurs;
drop table cages;
-- +goose StatementEnd
