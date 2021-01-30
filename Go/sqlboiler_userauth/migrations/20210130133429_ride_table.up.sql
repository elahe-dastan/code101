CREATE TABLE Ride
(
    id          int PRIMARY KEY,
    passenger_id int,
    FOREIGN KEY (passenger_id) REFERENCES passenger (id)
);