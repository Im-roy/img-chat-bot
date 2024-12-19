CREATE TABLE IF NOT EXISTS "chat-bot"."users_filepath_mappings" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    filepath VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);