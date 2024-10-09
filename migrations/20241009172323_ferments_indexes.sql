-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_uuid ON fermentations (uuid);
-- +goose StatementEnd

-- +goose Down
DROP INDEX idx_uuid ON fermentations (uuid);
-- +goose StatementBegin
-- +goose StatementEnd
