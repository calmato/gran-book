-- -----------------------------------------------------
-- Schema users
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `users` DEFAULT CHARACTER SET utf8mb4 ;
USE `users` ;

-- -----------------------------------------------------
-- Table `users`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `users`.`users` (
  `id`                VARCHAR(36)         NOT NULL,              -- ユーザーID
  `username`          VARCHAR(32)         NOT NULL,              -- 表示名
  `gender`            TINYINT(2) UNSIGNED NOT NULL DEFAULT 0,    -- 性別
  `email`             VARCHAR(256)        NULL     DEFAULT NULL, -- メールアドレス
  `role`              TINYINT(2) UNSIGNED NOT NULL DEFAULT 0,    -- 権限
  `thumbnail_url`     TEXT(8192)          NULL,                  -- サムネイルURL
  `self_introduction` VARCHAR(256)        NULL     DEFAULT NULL, -- 自己紹介
  `last_name`         VARCHAR(16)         NULL     DEFAULT NULL, -- 姓
  `first_name`        VARCHAR(16)         NULL     DEFAULT NULL, -- 名
  `last_name_kana`    VARCHAR(32)         NULL     DEFAULT NULL, -- 姓(かな)
  `first_name_kana`   VARCHAR(32)         NULL     DEFAULT NULL, -- 名(かな)
  `postal_code`       VARCHAR(16)         NULL     DEFAULT NULL, -- 郵便番号
  `prefecture`        VARCHAR(32)         NULL     DEFAULT NULL, -- 都道府県
  `city`              VARCHAR(32)         NULL     DEFAULT NULL, -- 市区町村
  `address_line1`     VARCHAR(64)         NULL     DEFAULT NULL, -- 町名・番地
  `address_line2`     VARCHAR(64)         NULL     DEFAULT NULL, -- ビル名・号室など
  `phone_number`      VARCHAR(16)         NULL     DEFAULT NULL, -- 電話番号
  `instance_id`       VARCHAR(256)        NULL     DEFAULT NULL, -- デバイスID
  `created_at`        DATETIME            NOT NULL,              -- 作成日時
  `updated_at`        DATETIME            NOT NULL,              -- 更新日時
  `deleted_at`        DATETIME            NULL     DEFAULT NULL, -- 退会日時
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `email_UNIQUE` ON `users`.`users` (`email` ASC) VISIBLE;

-- -----------------------------------------------------
-- Table `users`.`relationships`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `users`.`relationships` (
  `id`          BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- リレーションシップID
  `follow_id`   VARCHAR(36)         NOT NULL,                -- フォローID
  `follower_id` VARCHAR(36)         NOT NULL,                -- フォロワーID
  `created_at`  DATETIME            NOT NULL,                -- 作成日時
  `updated_at`  DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_relationships_users_follow_id`
    FOREIGN KEY (`follow_id`)
    REFERENCES `users`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_relationships_users_follower_id`
    FOREIGN KEY (`follower_id`)
    REFERENCES `users`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `ui_relationships_follow_id_follower_id` ON `users`.`relationships` (`follow_id` ASC, `follower_id` ASC) VISIBLE;
CREATE UNIQUE INDEX `ui_relationships_follower_id_follow_id` ON `users`.`relationships` (`follower_id` ASC, `follow_id` ASC) VISIBLE;

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
    ON UPDATE CASCADE,
  CONSTRAINT `fk_bookshelves_reviews_review_id`
    FOREIGN KEY (`review_id`)
    REFERENCES `books`.`reviews` (`id`)
    ON DELETE CASCADE
    ON UPDATE SET NULL
) ENGINE = InnoDB;

CREATE INDEX `idx_bookshelves_book_id` ON `books`.`bookshelves` (`book_id` ASC) VISIBLE;
CREATE UNIQUE INDEX `review_id_UNIQUE` ON `books`.`bookshelves` (`review_id` DESC) VISIBLE;
CREATE UNIQUE INDEX `ui_bookshelves_book_id_user_id` ON `books`.`bookshelves` (`book_id` ASC, `user_id` ASC) VISIBLE;
CREATE UNIQUE INDEX `ui_bookshelves_user_id_book_id` ON `books`.`bookshelves` (`user_id` ASC, `book_id` ASC) VISIBLE;

-- -----------------------------------------------------
-- Schema informations
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `informations` DEFAULT CHARACTER SET utf8mb4 ;
USE `informations` ;

-- -----------------------------------------------------
-- Table `informations`.`categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `informations`.`categories` (
  `id`         BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- カテゴリID
  `name`       VARCHAR(64)         NOT NULL,                -- カテゴリ名
  `created_at` DATETIME            NOT NULL,                -- 作成日時
  `updated_at` DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `name_UNIQUE` ON `informations`.`categories` (`name` ASC) VISIBLE;

-- -----------------------------------------------------
-- Table `informations`.`notifications`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `informations`.`notifications` (
  `id`          BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- お知らせID
  `author_id`   VARCHAR(36)         NULL     DEFAULT NULL,   -- 作成者ID
  `editor_id`   VARCHAR(36)         NULL     DEFAULT NULL,   -- 更新者ID
  `category_id` BIGINT(20) UNSIGNED NULL     DEFAULT NULL,   -- カテゴリID
  `title`       VARCHAR(64)         NOT NULL,                -- タイトル
  `description` TEXT(2000)          NOT NULL,                -- 本文
  `importance`  TINYINT(4)          NOT NULL DEFAULT 0,      -- 重要度
  `created_at`  DATETIME            NOT NULL,                -- 作成日時
  `updated_at`  DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_notifications_categories_category_id`
    FOREIGN KEY (`category_id`)
    REFERENCES `informations`.`categories` (`id`)
    ON DELETE SET NULL
    ON UPDATE CASCADE
) ENGINE = InnoDB;

CREATE INDEX `idx_notifications_category_id` ON `informations`.`notifications` (`category_id` ASC) VISIBLE;
