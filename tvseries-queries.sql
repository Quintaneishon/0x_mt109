-- create tables
CREATE TABLE Serie (
  ID INTEGER PRIMARY KEY,
  Name TEXT NOT NULL,
  Type TEXT NOT NULL
);
CREATE TABLE Actor (
  ID INTEGER PRIMARY KEY,
  Name TEXT NOT NULL,
  Phone TEXT NOT NULL,
  Email TEXT NOT NULL
);
CREATE TABLE Director (
  ID INTEGER PRIMARY KEY,
  Name TEXT NOT NULL,
  Phone TEXT NOT NULL,
  Email TEXT NOT NULL
);

CREATE TABLE Episode(
    ID INTEGER PRIMARY KEY,
    Name TEXT NOT NULL,
    Typ TEXT NOT NULL,
    ID_Director INTEGER NOT NULL,
    ID_Serie INTEGER NOT NULL,
    CONSTRAINT fk_Episode_Director FOREIGN KEY (ID_Director) REFERENCES Director(ID),
    CONSTRAINT fk_Episode_Serie FOREIGN KEY (ID_Serie) REFERENCES Serie(ID)
);
CREATE TABLE Actor_Serie(
    ID_Actor INTEGER NOT NULL,
    ID_Serie INTEGER NOT NULL,
    PRIMARY KEY (ID_Actor, ID_Serie),
    CONSTRAINT fk_Actor_Serie_Actor FOREIGN KEY (ID_Actor) REFERENCES Actor(ID),
    CONSTRAINT fk_Actor_Serie_Serie FOREIGN KEY (ID_Serie) REFERENCES Serie(ID)
);

-- GET THE ESPECIFIC INFORMATION
-- Which actors play in the series Big Sister?
SELECT Name 
FROM Actor 
WHERE ID IN (  
    SELECT ID_Actor 
    FROM Actor_Serie 
    WHERE ID_Serie = (
        SELECT ID 
        FROM Serie 
        WHERE Name = 'Big Sister'
    )
);

-- Which director has directed the greatest number of episodes?
SELECT Name 
FROM Director
WHERE ID = (
    SELECT ID_Director
    FROM Episode
    GROUP BY ID_Director
    ORDER BY COUNT(ID_Director) DESC
    LIMIT 1
);
