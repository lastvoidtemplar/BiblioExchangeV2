-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Messages (
  message_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user1_id UUID NOT NULL,
  user2_id UUID NOT NULL,
  message_type INTEGER, --[note: '0 for text and 1 and attachment'] 
  content VARCHAR(1000)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Messages
-- +goose StatementEnd
