--
-- Скрипт сгенерирован Devart dbForge Studio for MySQL, Версия 7.2.58.0
-- Домашняя страница продукта: http://www.devart.com/ru/dbforge/mysql/studio
-- Дата скрипта: 10.08.2018 17:31:00
-- Версия сервера: 5.5.5-10.1.26-MariaDB
-- Версия клиента: 4.1
--


-- 
-- Отключение внешних ключей
-- 
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;

-- 
-- Установить режим SQL (SQL mode)
-- 
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 
-- Установка кодировки, с использованием которой клиент будет посылать запросы на сервер
--
SET NAMES 'utf8';

-- 
-- Установка базы данных по умолчанию
--
USE slonworks;

--
-- Описание для таблицы fibonachi
--
DROP TABLE IF EXISTS fibonachi;
CREATE TABLE fibonachi (
  id INT(11) NOT NULL AUTO_INCREMENT,
  param INT(11) NOT NULL,
  result INT(11) NOT NULL,
  PRIMARY KEY (id)
)
ENGINE = INNODB
AUTO_INCREMENT = 21
AVG_ROW_LENGTH = 2340
CHARACTER SET utf8
COLLATE utf8_general_ci;

-- 
-- Вывод данных для таблицы fibonachi
--
INSERT INTO fibonachi VALUES
(1, 1, 1),
(2, 2, 1),
(3, 3, 2),
(4, 4, 3),
(5, 5, 5),
(6, 6, 8),
(7, 7, 13),
(8, 8, 21),
(9, 9, 34),
(10, 10, 55),
(11, 11, 89),
(12, 12, 144),
(13, 13, 233),
(14, 14, 377),
(15, 15, 610),
(16, 16, 987),
(17, 17, 1597),
(18, 18, 2584),
(19, 19, 4181),
(20, 20, 6765);

-- 
-- Восстановить предыдущий режим SQL (SQL mode)
-- 
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;

-- 
-- Включение внешних ключей
-- 
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;