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

-- -----------------------------------------------------
-- Table `informations`.`inquiries`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `informations`.`inquiries` (
  `id`          BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT, -- お問い合わせID
  `sender_id`   VARCHAR(36)         NULL     DEFAULT NULL,   -- 問い合わせ者ID
  `admin_id`    VARCHAR(36)         NULL     DEFAULT NULL,   -- 最終対応者ID
  `subject`     VARCHAR(64)         NOT NULL,                -- タイトル
  `description` TEXT(2000)          NOT NULL,                -- 詳細
  `email`       VARCHAR(256)        NOT NULL,                -- メールアドレス
  `is_replied`  TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,      -- 対応完了フラグ
  `created_at`  DATETIME            NOT NULL,                -- 作成日時
  `updated_at`  DATETIME            NOT NULL,                -- 更新日時
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE INDEX `idx_inquiries_updated_at` ON `informations`.`inquiries` (`updated_at` DESC) VISIBLE;

