-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo (
   id serial PRIMARY KEY,
   name text not null,
   finished bool not null default FALSE,
   created_at TIMESTAMP not null default CURRENT_TIMESTAMP,
   updated TIMESTAMP null 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop TABLE todo;
-- +goose StatementEnd
