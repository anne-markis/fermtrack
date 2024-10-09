-- +goose Up
-- +goose StatementBegin
ALTER TABLE fermentations
ADD COLUMN user_uuid CHAR(36);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE fermentations
DROP COLUMN user_uuid;
-- +goose StatementEnd
