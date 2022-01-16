-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
  id INT NOT NULL AUTO_INCREMENT,
  hash VARCHAR(20) NOT NULL,
  url VARCHAR(500) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT NOW(),
  updated_at DATETIME,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd
