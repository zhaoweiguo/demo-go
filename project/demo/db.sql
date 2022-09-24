
create db if not exists videos
use videos
create table video (
    id int primary key,
    user_id int,
    video_type string,
    video_name string,
    isdeleted bool,
    updated_at timestamp,
    created_at timestamp
)

