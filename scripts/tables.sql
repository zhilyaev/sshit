create table if not exists Log
(
    UUID UUID,
    SessionUUID UUID,
    Doc String,
    Created DateTime
)
    engine = MergeTree
    order by UUID;

create table if not exists Session
(
    UUID UUID,
    Created DateTime,
    Closed DateTime,
    User String,
    Remote String
)
    engine = MergeTree
        order by UUID;

