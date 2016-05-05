-- phpMyAdmin SQL Dump
-- version 4.3.11
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Erstellungszeit: 05. Mai 2016 um 13:49
-- Server-Version: 5.6.24
-- PHP-Version: 5.6.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Datenbank: `makershare`
--

--
-- Daten für Tabelle `category`
--

INSERT INTO `category` (`category_id`, `parent_id`, `name`, `description`) VALUES
(1, NULL, 'Transistor', 'Hi'),
(2, 1, 'n-Channel', 'Bye'),
(3, 1, 'p-Channel', 'sweg'),
(4, 2, 'nnnn', 'aaaa'),
(5, NULL, 'Switch', ''),
(6, 5, 'Push Button', 'push me'),
(7, 5, 'Toggle Switch', 'asdasd');

--
-- Daten für Tabelle `group`
--

INSERT INTO `group` (`group_id`, `name`, `description`) VALUES
('mEt?usta4e', 'First group', 'This is le first group'),
('y6prAf=uBa', 'Second group', 'this is le second group');

--
-- Daten für Tabelle `item`
--

INSERT INTO `item` (`item_id`, `coins`, `status`, `quantity`, `object_id`, `stock_id`) VALUES
(1, 5, 1, 0, 1, 1),
(2, 7, 1, 3, 2, 1),
(3, 4, 1, 20, 2, 2),
(4, 10, 1, 123, 2, 3),
(5, 15, 1, 5, 3, 1),
(6, 1, 1, 4, 1, 1),
(7, 7, 1, 2, 4, 2),
(8, 35, 1, 123, 3, 1),
(9, 3, 1, 4, 1, 1),
(10, 3, 1, 11, 1, 1),
(11, 3, 1, 4, 4, 2),
(12, 1, 1, 1, 1, 1),
(13, 5, 1, 3, 2, 1),
(14, 4, 1, 3, 1, 1);

--
-- Daten für Tabelle `object`
--

INSERT INTO `object` (`object_id`, `name`, `description`, `category_id`) VALUES
(1, '2n2222a', 'Transistor Central Semiconductor corp. 2N2222A NPN Gehäuseart TO-18 I(C) 800 mA Emitter-Sperrspannung U(CEO) 40 V', 2),
(2, 'BC547', 'lol', 2),
(3, '2n3906', 'Pnp transistor', 3),
(4, 'tact switch', 'push meee', 6);

--
-- Daten für Tabelle `stock`
--

INSERT INTO `stock` (`stock_id`, `name`, `location`) VALUES
(1, 'Matthias''s warehouse', '46.779620, 11.689619'),
(2, 'Gardena Stock 1', 'yolo'),
(3, 'Unnamed stock', 'Unknown location');

--
-- Daten für Tabelle `user`
--

INSERT INTO `user` (`user_id`, `username`, `password`, `email`, `name`, `surname`, `balance`, `type`, `verify_code`, `salt`, `group_id`) VALUES
(1, 'matthias', 'd2b344fae3e2c417d8d59afb3614f152c82d4652bed65425cec4da730a61579f', 'matpuz81@libero.it', 'Matthias', 'Moroder', 0, 1, NULL, 0x510c704c0aef13d722bf177bce26d98c, 'mEt?usta4e'),
(2, 'sepp', 'c7d099bcbf9101b0a49a7323c476d44d442699a9270b69d20f4191f29e1379fd', 'sepp@lopp.com', 'Sepp', 'Lopp', 0, 2, NULL, 0xb8e7992822651f2e04e0c37f898a41f6, 'mEt?usta4e'),
(3, 'Franz', '6d01ca5211bd8f81fbda66053f21d151256f488a760faf658ece206ed4358595', 'franztest@lol.it', 'Franz', 'Test', 0, 1, NULL, 0x86a7031f3508f737316e8a17bca30164, 'y6prAf=uBa'),
(4, 'seppl', '3193aa74d879fdb095af4ab0fd61b1cdbd90962123def149ab2771a596e41068', 'asd@asd.it', 'Seppl', 'Loppl', 0, 2, NULL, 0xcdcb1d90199191253eee45be109a708a, NULL),
(5, 'asfdasdf', '209bac24b721f7aa488b3e52e0d6bdce1b2f1d2be049c2ff620f5628f4d0ec8d', 'asd@asd.it', 'asdfasfds', 'sadfasdfa', 0, 2, NULL, 0x6169e81672bf8223667e5759025fc09b, NULL);

--
-- Daten für Tabelle `user_stock`
--

INSERT INTO `user_stock` (`user_id`, `stock_id`) VALUES
(1, 1),
(1, 2),
(2, 3);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
