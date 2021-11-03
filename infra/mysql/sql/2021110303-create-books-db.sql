-- -----------------------------------------------------
-- Schema books
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `books` DEFAULT CHARACTER SET utf8mb4 ;
USE `books` ;

-- -----------------------------------------------------
-- Table `books`.`authors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`authors` (
  `id`         BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- 著者ID
  `name`       VARCHAR(32)         NOT NULL,                -- 著者名
  `name_kana`  VARCHAR(64)         NOT NULL,                -- 著者名(かな)
  `created_at`  DATETIME           NOT NULL,                -- 作成日時
  `updated_at`  DATETIME           NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `name_UNIQUE` ON `books`.`authors` (`name` ASC) VISIBLE;

-- -----------------------------------------------------
-- Table `books`.`books`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`books` (
  `id`               BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- 書籍ID
  `title`            VARCHAR(64)         NOT NULL,                -- 書籍名
  `title_kana`       VARCHAR(128)        NOT NULL,                -- 書籍名(かな)
  `description`      TEXT(2000)          NULL,                    -- 説明
  `isbn`             VARCHAR(13)         NOT NULL,                -- ISBN
  `publisher`        VARCHAR(64)         NOT NULL,                -- 出版社
  `published_on`     VARCHAR(16)         NULL     DEFAULT NULL,   -- 出版日
  `thumbnail_url`    TEXT(8192)          NULL,                    -- サムネイルURL
  `rakuten_url`      TEXT(8192)          NULL,                    -- 楽天ショップURL
  `rakuten_size`     VARCHAR(64)         NULL     DEFAULT NULL,   -- 楽天ショップカテゴリ
  `rakuten_genre_id` VARCHAR(64)         NULL     DEFAULT '000',  -- 楽天ショップジャンルID
  `created_at`       DATETIME            NOT NULL,                -- 作成日時
  `updated_at`       DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `isbn_UNIQUE` ON `books`.`books` (`isbn` ASC) VISIBLE;

-- -----------------------------------------------------
-- Table `books`.`authors_books`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`authors_books` (
  `id`               BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- 中間テーブルID
  `book_id`          BIGINT(20) UNSIGNED NOT NULL,                -- 書籍ID
  `author_id`        BIGINT(20) UNSIGNED NOT NULL,                -- 著者ID
  `created_at`       DATETIME            NOT NULL,                -- 作成日時
  `updated_at`       DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_authors_books_book_id_author_id` ON `books`.`authors_books` (`book_id` ASC, `author_id` ASC) VISIBLE;

-- -----------------------------------------------------
-- Table `books`.`reviews`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`reviews` (
  `id`         BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- レビューID
  `book_id`    BIGINT(20) UNSIGNED NOT NULL,                -- 書籍ID
  `user_id`    VARCHAR(36)         NULL     DEFAULT NULL,   -- ユーザーID
  `score`      TINYINT(4) UNSIGNED NOT NULL DEFAULT 0,      -- 5段階評価
  `impression` TEXT(2000)          NULL,                    -- 感想
  `created_at` DATETIME            NOT NULL,                -- 作成日時
  `updated_at` DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_reviews_books_book_id`
    FOREIGN KEY (`book_id`)
    REFERENCES `books`.`books` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE INDEX `idx_reviews_book_id_created_at` ON `books`.`reviews` (`book_id` ASC, `created_at` DESC) VISIBLE;
CREATE UNIQUE INDEX `ui_reviews_book_id_user_id` ON `books`.`reviews` (`book_id` ASC, `user_id` ASC) VISIBLE;
CREATE UNIQUE INDEX `ui_reviews_user_id_book_id` ON `books`.`reviews` (`user_id` ASC, `book_id` ASC) VISIBLE;

-- -----------------------------------------------------
-- Table `books`.`bookshelves`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`bookshelves` (
  `id`         BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- 本棚ID
  `book_id`    BIGINT(20) UNSIGNED NOT NULL,                -- 書籍ID
  `user_id`    VARCHAR(36)         NOT NULL,                -- ユーザーID
  `review_id`  BIGINT(20) UNSIGNED NULL     DEFAULT 0,      -- レビューID
  `status`     TINYINT(4) UNSIGNED NOT NULL DEFAULT 0,      -- 読書ステータス
  `read_on`    DATETIME            NULL     DEFAULT NULL,   -- 読み終えた日
  `created_at` DATETIME            NOT NULL,                -- 作成日時
  `updated_at` DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_bookshelves_books_book_id`
    FOREIGN KEY (`book_id`)
    REFERENCES `books`.`books` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE INDEX `idx_bookshelves_book_id` ON `books`.`bookshelves` (`book_id` ASC) VISIBLE;
CREATE UNIQUE INDEX `ui_bookshelves_book_id_user_id` ON `books`.`bookshelves` (`book_id` ASC, `user_id` ASC) VISIBLE;
CREATE UNIQUE INDEX `ui_bookshelves_user_id_book_id` ON `books`.`bookshelves` (`user_id` ASC, `book_id` ASC) VISIBLE;
