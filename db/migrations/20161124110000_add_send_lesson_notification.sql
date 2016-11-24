-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user ADD COLUMN `send_lesson_notification` tinyint unsigned NOT NULL AFTER `followed_teacher_at`;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user DROP COLUMN `send_lesson_notification`;
