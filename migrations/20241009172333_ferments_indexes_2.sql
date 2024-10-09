-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_user_uuid ON fermentations (user_uuid);
-- +goose StatementEnd

-- +goose Down
DROP INDEX idx_user_uuid ON fermentations (user_uuid);
-- +goose StatementBegin
-- +goose StatementEnd
