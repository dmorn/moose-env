-- phpMyAdmin SQL Dump
-- version 4.3.11
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Erstellungszeit: 06. Jun 2016 um 13:29
-- Server-Version: 5.6.24
-- PHP-Version: 5.6.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Datenbank: `moosedb`
--

DELIMITER $$
--
-- Funktionen
--
CREATE DEFINER=`root`@`localhost` FUNCTION `isSubcategoryOf`(`icategory_id` INT, `iparent_id` INT) RETURNS int(11)
BEGIN
	
	DECLARE found_parent INT DEFAULT 0;

	IF icategory_id=iparent_id THEN
		RETURN 1;
	END IF;


	
	WHILE icategory_id > 0 DO
		SELECT parent_id INTO found_parent
		FROM category WHERE
		category_id=icategory_id;

		IF found_parent=iparent_id THEN
			RETURN 1;
		END IF;
		SET icategory_id = found_parent;

	END WHILE;

	RETURN 0;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `category`
--

CREATE TABLE IF NOT EXISTS `category` (
  `category_id` int(11) NOT NULL,
  `parent_id` int(11) NOT NULL,
  `name` varchar(40) NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

--
-- Daten für Tabelle `category`
--

INSERT INTO `category` (`category_id`, `parent_id`, `name`, `description`) VALUES
(1, 0, 'Transistor', 'Hi'),
(2, 1, 'n-Channel', 'Bye'),
(3, 1, 'p-Channel', 'sweg'),
(5, 0, 'Switch', ''),
(6, 5, 'Push Button', 'push me'),
(7, 5, 'Toggle Switch', 'asdasd'),
(8, 0, 'Led', 'Shine bright'),
(9, 8, 'Red Led', 'Reeed'),
(10, 8, 'Blue Led', 'Blueee');

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `group`
--

CREATE TABLE IF NOT EXISTS `group` (
  `group_id` int(10) NOT NULL,
  `name` varchar(40) NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Daten für Tabelle `group`
--

INSERT INTO `group` (`group_id`, `name`, `description`) VALUES
(1, 'First group', 'This is le first group');

-- --------------------------------------------------------

--
-- Stellvertreter-Struktur des Views `group_items`
--
CREATE TABLE IF NOT EXISTS `group_items` (
`item_id` int(11)
,`name` varchar(100)
,`description` text
,`category_id` int(11)
,`coins` int(11)
,`status` int(11)
,`quantity` int(11)
,`stock_name` varchar(100)
,`group_id` int(10)
,`stock_id` int(11)
);

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `group_request`
--

CREATE TABLE IF NOT EXISTS `group_request` (
  `group_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Stellvertreter-Struktur des Views `group_stocks`
--
CREATE TABLE IF NOT EXISTS `group_stocks` (
`stock_id` int(11)
,`name` varchar(100)
,`location` varchar(100)
,`group_id` int(10)
,`user_id` int(11)
);

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `item`
--

CREATE TABLE IF NOT EXISTS `item` (
  `item_id` int(11) NOT NULL,
  `coins` int(11) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `quantity` int(11) NOT NULL,
  `link` text NOT NULL,
  `object_id` int(11) NOT NULL,
  `stock_id` int(11) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `object`
--

CREATE TABLE IF NOT EXISTS `object` (
  `object_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text NOT NULL,
  `category_id` int(11) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

--
-- Daten für Tabelle `object`
--

INSERT INTO `object` (`object_id`, `name`, `description`, `category_id`) VALUES
(1, '2n2222a', 'Transistor Central Semiconductor corp. 2N2222A NPN Gehäuseart TO-18 I(C) 800 mA Emitter-Sperrspannung U(CEO) 40 V', 2),
(2, 'BC547', 'lol', 2),
(3, '2n3906', 'Pnp transistor', 3),
(4, 'tact switch', 'push meee', 6),
(5, 'Button', 'clickediclick', 7),
(6, '3mm Red Led', 'red led, 3mm diameter, 15mA', 9),
(7, '3mm blue led', 'blue led, 3mm, 16mA', 10);

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `stock`
--

CREATE TABLE IF NOT EXISTS `stock` (
  `stock_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `location` varchar(100) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

--
-- Daten für Tabelle `stock`
--

INSERT INTO `stock` (`stock_id`, `name`, `location`) VALUES
(1, 'Matthias''s warehouse', '46.779620, 11.689619');

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `user_id` int(11) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` char(64) NOT NULL,
  `email` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `surname` varchar(100) NOT NULL,
  `balance` int(11) NOT NULL,
  `type` int(11) NOT NULL DEFAULT '2',
  `group_id` int(10) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

--
-- Daten für Tabelle `user`
--

INSERT INTO `user` (`user_id`, `username`, `password`, `email`, `name`, `surname`, `balance`, `type`, `group_id`) VALUES
(1, 'ilenia', '6b1808a2ee5b1144c0853ad4beb7c717697a18f4c3292e778b91b4d23ded6f37', 'emai@mail.it', 'Ilenia', 'Fronza', 500, 1, 1),
(2, 'matthias', '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08', 'mamoroder@unibz.iz', 'Matthias', 'Moroder', 500, 2, 1);

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `user_stock`
--

CREATE TABLE IF NOT EXISTS `user_stock` (
  `user_id` int(11) NOT NULL,
  `stock_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Daten für Tabelle `user_stock`
--

INSERT INTO `user_stock` (`user_id`, `stock_id`) VALUES
(2, 1);

-- --------------------------------------------------------

--
-- Struktur des Views `group_items`
--
DROP TABLE IF EXISTS `group_items`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `group_items` AS select `item`.`item_id` AS `item_id`,`object`.`name` AS `name`,`object`.`description` AS `description`,`object`.`category_id` AS `category_id`,`item`.`coins` AS `coins`,`item`.`status` AS `status`,`item`.`quantity` AS `quantity`,`stock`.`name` AS `stock_name`,`user`.`group_id` AS `group_id`,`stock`.`stock_id` AS `stock_id` from ((((`item` join `object` on((`item`.`object_id` = `object`.`object_id`))) join `stock` on((`item`.`stock_id` = `stock`.`stock_id`))) join `user_stock` on((`item`.`stock_id` = `user_stock`.`stock_id`))) join `user` on((`user_stock`.`user_id` = `user`.`user_id`)));

-- --------------------------------------------------------

--
-- Struktur des Views `group_stocks`
--
DROP TABLE IF EXISTS `group_stocks`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `group_stocks` AS select `stock`.`stock_id` AS `stock_id`,`stock`.`name` AS `name`,`stock`.`location` AS `location`,`user`.`group_id` AS `group_id`,`user`.`user_id` AS `user_id` from ((`stock` join `user_stock` on((`stock`.`stock_id` = `user_stock`.`stock_id`))) join `user` on((`user_stock`.`user_id` = `user`.`user_id`)));

--
-- Indizes der exportierten Tabellen
--

--
-- Indizes für die Tabelle `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`category_id`);

--
-- Indizes für die Tabelle `group`
--
ALTER TABLE `group`
  ADD PRIMARY KEY (`group_id`);

--
-- Indizes für die Tabelle `item`
--
ALTER TABLE `item`
  ADD PRIMARY KEY (`item_id`);

--
-- Indizes für die Tabelle `object`
--
ALTER TABLE `object`
  ADD PRIMARY KEY (`object_id`);

--
-- Indizes für die Tabelle `stock`
--
ALTER TABLE `stock`
  ADD PRIMARY KEY (`stock_id`);

--
-- Indizes für die Tabelle `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`user_id`), ADD UNIQUE KEY `username` (`username`);

--
-- Indizes für die Tabelle `user_stock`
--
ALTER TABLE `user_stock`
  ADD PRIMARY KEY (`user_id`,`stock_id`);

--
-- AUTO_INCREMENT für exportierte Tabellen
--

--
-- AUTO_INCREMENT für Tabelle `category`
--
ALTER TABLE `category`
  MODIFY `category_id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=11;
--
-- AUTO_INCREMENT für Tabelle `item`
--
ALTER TABLE `item`
  MODIFY `item_id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=16;
--
-- AUTO_INCREMENT für Tabelle `object`
--
ALTER TABLE `object`
  MODIFY `object_id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=8;
--
-- AUTO_INCREMENT für Tabelle `stock`
--
ALTER TABLE `stock`
  MODIFY `stock_id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT für Tabelle `user`
--
ALTER TABLE `user`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=3;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
