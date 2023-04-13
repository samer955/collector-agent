create table testdb.cpu
(
    uuid        varchar(100) not null
        primary key,
    hostname    text         null,
    model       text         null,
    utilization float        null,
    time        timestamp    null
);

create table testdb.memory
(
    uuid        varchar(100) not null
        primary key,
    hostname    text         null,
    total       float        null,
    utilization float        null,
    time        timestamp    null
);

create table testdb.`system`
(
    uuid         varchar(100) not null
        primary key,
    hostname     text         null,
    ip           text         null,
    os           text         null,
    architecture text         null,
    platform     text         null,
    version      text         null,
    online_users int          null,
    latency      int          null,
    time         timestamp    null
);

create table testdb.tcp
(
    uuid              varchar(100) not null
        primary key,
    hostname          text         null,
    queue_size        int          null,
    segments_received int          null,
    segments_sent     int          null,
    time              timestamp    null
);