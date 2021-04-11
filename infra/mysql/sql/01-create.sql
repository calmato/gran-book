-- MySQL Script generated by MySQL Workbench
-- Mon Apr 12 02:03:09 2021
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema users
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema users
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `users` DEFAULT CHARACTER SET utf8mb4 ;
-- -----------------------------------------------------
-- Schema books
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema books
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `books` DEFAULT CHARACTER SET utf8mb4 ;
-- -----------------------------------------------------
-- Schema stores
-- -----------------------------------------------------
USE `users` ;

-- -----------------------------------------------------
-- Table `users`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `users`.`users` (
  `id` VARCHAR(36) NOT NULL,
  `username` VARCHAR(32) NOT NULL DEFAULT '',
  `gender` TINYINT(2) UNSIGNED NOT NULL DEFAULT '0',
  `email` VARCHAR(256) NULL DEFAULT NULL,
  `role` TINYINT(2) UNSIGNED NOT NULL DEFAULT '0',
  `thumbnail_url` TEXT(8192) NULL DEFAULT NULL,
  `self_introduction` VARCHAR(256) NULL DEFAULT NULL,
  `last_name` VARCHAR(16) NULL DEFAULT NULL,
  `first_name` VARCHAR(16) NULL DEFAULT NULL,
  `last_name_kana` VARCHAR(32) NULL DEFAULT NULL,
  `first_name_kana` VARCHAR(32) NULL DEFAULT NULL,
  `postal_code` VARCHAR(16) NULL DEFAULT NULL,
  `prefecture` VARCHAR(32) NULL DEFAULT NULL,
  `city` VARCHAR(32) NULL DEFAULT NULL,
  `address_line1` VARCHAR(64) NULL DEFAULT NULL,
  `address_line2` VARCHAR(64) NULL DEFAULT NULL,
  `phone_number` VARCHAR(16) NULL DEFAULT NULL,
  `instance_id` VARCHAR(256) NULL DEFAULT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `users`.`relationships`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `users`.`relationships` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `follow_id` VARCHAR(36) NOT NULL,
  `follower_id` VARCHAR(36) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_relationships_01` (`follow_id` ASC) VISIBLE,
  INDEX `idx_relationships_02` (`follower_id` ASC) VISIBLE,
  UNIQUE INDEX `ui_relationships_01` (`follow_id` ASC, `follower_id` ASC) VISIBLE,
  UNIQUE INDEX `ui_relationships_02` (`follower_id` ASC, `follow_id` ASC) VISIBLE,
  CONSTRAINT `fk_relationships_users_01`
    FOREIGN KEY (`follow_id`)
    REFERENCES `users`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_relationships_users_02`
    FOREIGN KEY (`follower_id`)
    REFERENCES `users`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `books` ;

-- -----------------------------------------------------
-- Table `books`.`authors`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`authors` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL DEFAULT '',
  `name_kana` VARCHAR(64) NOT NULL DEFAULT '',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `books`.`books`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`books` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(64) NOT NULL DEFAULT '',
  `title_kana` VARCHAR(128) NOT NULL DEFAULT '',
  `description` TEXT(2000) NULL DEFAULT NULL,
  `isbn` VARCHAR(16) NOT NULL DEFAULT '',
  `publisher` VARCHAR(64) NOT NULL DEFAULT '',
  `published_on` DATE NOT NULL,
  `thumbnail_url` TEXT(8192) NULL DEFAULT NULL,
  `rakuten_url` TEXT(8192) NULL DEFAULT NULL,
  `rakuten_genre_id` VARCHAR(32) NULL DEFAULT '000',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `isbn_UNIQUE` (`isbn` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `books`.`authors_books`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`authors_books` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `book_id` BIGINT(20) UNSIGNED NOT NULL,
  `author_id` BIGINT(20) UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `ui_authors_books_01` (`book_id` ASC, `author_id` ASC) VISIBLE,
  UNIQUE INDEX `ui_authors_books_02` (`author_id` ASC, `book_id` ASC) VISIBLE,
  INDEX `idx_authors_books_01` (`book_id` ASC) VISIBLE,
  INDEX `idx_authors_books_02` (`author_id` ASC) VISIBLE,
  CONSTRAINT `fk_authors_books_authors_01`
    FOREIGN KEY (`author_id`)
    REFERENCES `books`.`authors` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_authors_books_books_01`
    FOREIGN KEY (`book_id`)
    REFERENCES `books`.`books` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `books`.`bookshelves`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`bookshelves` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `book_id` BIGINT(20) UNSIGNED NOT NULL,
  `user_id` VARCHAR(36) NOT NULL,
  `status` TINYINT(4) UNSIGNED NOT NULL DEFAULT '0',
  `read_on` DATE NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_bookshelves_01` (`user_id` ASC) VISIBLE,
  INDEX `idx_bookshelves_02` (`book_id` ASC) VISIBLE,
  UNIQUE INDEX `ui_bookshelves_01` (`book_id` ASC, `user_id` ASC) VISIBLE,
  UNIQUE INDEX `ui_bookshelves_02` (`user_id` ASC, `book_id` ASC) VISIBLE,
  CONSTRAINT `fk_bookshelves_books_01`
    FOREIGN KEY (`book_id`)
    REFERENCES `books`.`books` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_bookshelves_users_01`
    FOREIGN KEY (`user_id`)
    REFERENCES `users`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `books`.`reviews`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `books`.`reviews` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `book_id` BIGINT(20) UNSIGNED NOT NULL,
  `user_id` VARCHAR(36) NULL,
  `score` TINYINT(4) UNSIGNED NOT NULL DEFAULT '0',
  `impression` TEXT(2000) NULL DEFAULT NULL,
  `created_at` DATETIME NOT NULL COMMENT '	',
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_reviews_01` (`book_id` ASC) VISIBLE,
  INDEX `idx_reviews_02` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `ui_reviews_01` (`book_id` ASC, `user_id` ASC) VISIBLE,
  UNIQUE INDEX `ui_reviews_02` (`user_id` ASC, `book_id` ASC) VISIBLE,
  CONSTRAINT `fk_reviews_books_01`
    FOREIGN KEY (`book_id`)
    REFERENCES `books`.`books` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_reviews_users_01`
    FOREIGN KEY (`user_id`)
    REFERENCES `users`.`users` (`id`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
