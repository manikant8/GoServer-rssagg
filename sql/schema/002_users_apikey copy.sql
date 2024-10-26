-- migrations

-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    -- random() generates number b/1/w 0 & 1. convert it to text then to byte as sha256 operates on byte
    --  encode function converts the binary output of the sha256 function into a hexadecimal string representation. The second argument 'hex' specifies the format.
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;