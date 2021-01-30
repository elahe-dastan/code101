CREATE TABLE Ride
(
    ID          int PRIMARY KEY,
    PassengerID int,
    FOREIGN KEY (PassengerID) REFERENCES passenger (ID)
);