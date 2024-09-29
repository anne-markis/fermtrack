-- +goose Up
-- +goose StatementBegin
INSERT INTO fermentations (uuid, nickname, start_at, bottled_at, recipe_notes, tasting_notes, deleted_at) 
VALUES
('550e8400-e29b-41d4-a716-446655440001', 'Syrah 2024', '2024-10-12 12:00:00', NULL, 'TODO long recipe', NULL, NULL),
('550e8400-e29b-41d4-a716-446655440010', 'watermelon 2023', '2023-05-05 12:00:00', NULL, 'winged it', NULL, NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM fermentations where uuid like '550e8400-e29b-41d4-a716%';
-- +goose StatementEnd
