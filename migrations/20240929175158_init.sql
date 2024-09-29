-- +goose Up
-- +goose StatementBegin
CREATE TABLE fermentations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    start_at DATETIME NOT NULL,
    bottled_at DATETIME,
    recipe_notes TEXT,
    tasting_notes TEXT,
    deleted_at DATETIME DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE fermentations;
-- +goose StatementEnd
