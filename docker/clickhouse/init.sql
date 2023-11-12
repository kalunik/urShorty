-- CREATE DATABASE IF NOT EXISTS urlshorty;

CREATE TABLE default.visit_history
(
    short_path  String,
    domain      String,
    created_at  DateTime,
    visited_at  DateTime,
    latitude    Float32,
    longitude   Float32,
    country     String,
    city        String,
    proxy       Boolean
)
    ENGINE = MergeTree()
        PRIMARY KEY (short_path, visited_at);