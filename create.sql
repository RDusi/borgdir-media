BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `Equipment` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`Bezeichnung`	TEXT,
	`Kategorie`	INTEGER,
	`InventarNr`	TEXT,
	`Lagerort`	INTEGER,
	`Inhalt`	TEXT,
	`Anzahl`	INTEGER,
	`Hinweise`	TEXT,
	`Bild`	TEXT
);
INSERT INTO `Equipment` VALUES (1,'Nikon D800','Kamera','14-156','Regal 5','Kamera Body + extra Akku',20,'Eine Kamera zum Fotos machen und Filmen','../../../static/images/nikon_d800.jpg');
INSERT INTO `Equipment` VALUES (2,'Canon 50mm 1.2','Objektiv','45-567','Regal 6','Objektiv',45,'scharfe Linse','../../../static/images/canon_50_12.jpg');
INSERT INTO `Equipment` VALUES (3,'Red Epic','Kamera','67-567','Regal 5','Brain + Bildschirm',3,'Filmkamera','../../../static/images/red_epic.jpg');
INSERT INTO `Equipment` VALUES (4,'Sennheiser MKE 600','Mikrofon','78-789','Regal 4','Mikrofon',5,'man kann etwas damit aufnehmen','../../../static/images/sennheiser_mke_600.jpg');
CREATE TABLE IF NOT EXISTS `User` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`Benutzername`	TEXT,
	`Email`	TEXT,
	`Passwort`	TEXT,
	`BenutzerTyp`	TEXT,
	`AktivBis`	TEXT,
	`Bild`	TEXT,
	`Created`	DATE
);
INSERT INTO `User` VALUES (2,'Erica Meier','erica.meier@gmail.com','gutesPasswort','Verleiher','gesperrt','../../../static/images/avatar_frau_150x150.jpg',NULL);
INSERT INTO `User` VALUES (8,'Test','','','Benutzer','','../../../static/images/avatar_150x150.jpg',NULL);
INSERT INTO `User` VALUES (9,'Meier','test@meier.de','TestPW','Benutzer','','../../../static/images/avatar_frau_150x150.jpg',NULL);
INSERT INTO `User` VALUES (11,'Neuer Kunde','kunde@kunde.de','123','Benutzer','aktiv','../../../static/images/avatar_150x150.jpg',NULL);
CREATE TABLE IF NOT EXISTS `Warenkorb` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER,
	`EquipmentID`	INTEGER,
	`EntleihDatum`	DATE,
	`RueckgabeDatum`	DATE,
  `Anzahl` INTEGER
);
CREATE TABLE IF NOT EXISTS `MeineGeraete` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER,
	`EquipmentID`	INTEGER,
	`EntleihDatum`	DATE,
	`RueckgabeDatum`	DATE
);
CREATE TABLE IF NOT EXISTS `Vorgemerkt` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER,
	`EquipmentID`	INTEGER,
	`RueckgabeDatum`	DATE,
  `Anzahl` INTEGER
);
CREATE TABLE IF NOT EXISTS `Kategorie` (
  `ID` INTEGER PRIMARY KEY AUTOINCREMENT,
  `KategorieName` TEXT
);
INSERT INTO `Kategorie` VALUES (1, 'Kameras');
INSERT INTO `Kategorie` VALUES (2, 'Objektive');
INSERT INTO `Kategorie` VALUES (3, 'Stative');
INSERT INTO `Kategorie` VALUES (4, 'Mikrofone');
CREATE TABLE IF NOT EXISTS `Lagerort` (
  `ID` INTEGER PRIMARY KEY AUTOINCREMENT,
  `LagerortName` TEXT
);
INSERT INTO `Lagerort` VALUES (1, 'Regal 4');
INSERT INTO `Lagerort` VALUES (2, 'Regal 5');
INSERT INTO `Lagerort` VALUES (3, 'Regal 6');
COMMIT;
