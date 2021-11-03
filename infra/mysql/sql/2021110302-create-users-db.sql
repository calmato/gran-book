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
