BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `Warenkorb` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER,
	`EquipmentID`	INTEGER,
	`EntleihDatum`	TEXT,
	`RueckgabeDatum`	TEXT,
	`Anzahl`	INTEGER
);
INSERT INTO `Warenkorb` VALUES (328,0,2,'20.06.2018','20.08.2018',1);
INSERT INTO `Warenkorb` VALUES (329,0,3,'20.06.2018','20.08.2018',1);
INSERT INTO `Warenkorb` VALUES (333,0,1,'21.06.2018','21.08.2018',1);
INSERT INTO `Warenkorb` VALUES (334,0,2,'21.06.2018','21.08.2018',1);
INSERT INTO `Warenkorb` VALUES (335,0,3,'21.06.2018','21.08.2018',1);
CREATE TABLE IF NOT EXISTS `Vorgemerkt` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER,
	`EquipmentID`	INTEGER,
	`RueckgabeDatum`	TEXT
);
CREATE TABLE IF NOT EXISTS `User` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`Benutzername`	TEXT,
	`Email`	TEXT,
	`Passwort`	REAL,
	`BenutzerTyp`	TEXT,
	`AktivBis`	TEXT,
	`Bild`	TEXT
);
INSERT INTO `User` VALUES (23,'jhoefker','jakob.hoefker@gmail.com','JDJhJDE0JG9qWmFucGI0aC5HcFgzN0FXc0RGMGVLcnZJcWNEcFlNRTVLNlBhbk8vRW5Xdm1Sek1nL2dx','Verleiher','immer','../../../static/images/jakob.jpg');
INSERT INTO `User` VALUES (32,'benutzer1','benutzer1@benutzer.de','JDJhJDE0JDMya203b3JySVNRVFNZb0ZwaTdGMnVGaDIxMjdnOVM1eWNxVUE2L1pzbDdPQVZPdldVVTA2','Benutzer','erstmal soweit','http://via.placeholder.com/350x350');
CREATE TABLE IF NOT EXISTS `Session` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER
);
INSERT INTO `Session` VALUES (2,12);
CREATE TABLE IF NOT EXISTS `MeineGeraete` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER,
	`EquipmentID`	INTEGER,
	`EntleihDatum`	TEXT,
	`RueckgabeDatum`	TEXT
);
INSERT INTO `MeineGeraete` VALUES (137,32,2,'20.06.2018','20.08.2018');
INSERT INTO `MeineGeraete` VALUES (138,32,3,'20.06.2018','20.12.2018');
INSERT INTO `MeineGeraete` VALUES (139,32,4,'20.06.2018','20.12.2018');
INSERT INTO `MeineGeraete` VALUES (140,32,1,'21.06.2018','21.08.2018');
INSERT INTO `MeineGeraete` VALUES (141,32,2,'21.06.2018','21.08.2018');
INSERT INTO `MeineGeraete` VALUES (142,32,3,'21.06.2018','21.08.2018');
CREATE TABLE IF NOT EXISTS `Lagerort` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`LagerortName`	TEXT
);
INSERT INTO `Lagerort` VALUES (1,'Regal 1');
INSERT INTO `Lagerort` VALUES (2,'Regal 2');
INSERT INTO `Lagerort` VALUES (3,'Regal 3');
CREATE TABLE IF NOT EXISTS `Kategorie` (
	`ID`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`KategorieName`	TEXT
);
INSERT INTO `Kategorie` VALUES (1,'Kamera');
INSERT INTO `Kategorie` VALUES (2,'Objektiv');
INSERT INTO `Kategorie` VALUES (3,'Stativ');
INSERT INTO `Kategorie` VALUES (4,'Mikrofon');
INSERT INTO `Kategorie` VALUES (5,'Drohne');
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
INSERT INTO `Equipment` VALUES (1,'Nikon D800',1,'14-156',1,'Kamera Body + extra Akku',498,'Eine Kamera zum Fotos machen und Filmen','../../../static/images/nikon_d800.jpg');
INSERT INTO `Equipment` VALUES (2,'Canon 50mm 1.2',2,'45-567',2,'Objektiv',496,'scharfe Linse','../../../static/images/canon_50_12.jpg');
INSERT INTO `Equipment` VALUES (3,'Red Epic',1,'67-567',3,'Brain + Bildschirm',496,'Filmkamera','../../../static/images/red_epic.jpg');
INSERT INTO `Equipment` VALUES (4,'Sennheiser MKE 600',3,'78-789',2,'Mikrofon',498,'man kann etwas damit aufnehmen','../../../static/images/sennheiser_mke_600.jpg');
INSERT INTO `Equipment` VALUES (23,'DJI Inspire 2',5,'789-789',3,'Eine Drohne',500,'Dies ist eine Drohne.','../../../static/images/inspire_2.jpg');
INSERT INTO `Equipment` VALUES (24,'Manfrotto 190',3,'678-678',2,'Ein Stativ',500,'','../../../static/images/Manfrotto-Aluminium-Dreibeinstativ-MT190GOA4TD.jpg');
COMMIT;
