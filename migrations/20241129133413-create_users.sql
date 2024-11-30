
-- +migrate Up
CREATE TABLE users(
    id BINARY(16) PRIMARY KEY COMMENT "ユーザーID",
    name VARCHAR(255) UNIQUE COMMENT "ユーザー名",
    password BINARY(255) NOT NULL COMMENT "パスワード",
    created_at DATETIME NOT NULL COMMENT "作成日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時"
) COMMENT = "ユーザー";

-- +migrate Down
DROP TABLE users;