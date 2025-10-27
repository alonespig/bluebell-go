-- 帖子统计表
CREATE TABLE post_stats (
    post_id BIGINT PRIMARY KEY,
    like_count BIGINT DEFAULT 0,
    dislike_count BIGINT DEFAULT 0,
    score BIGINT DEFAULT 0,
    last_update DATETIME(3)
);
-- post_id 对应主帖；
-- like_count、dislike_count 用于快速展示；
-- score 可用于排行榜排序；
-- 这张表的更新是RocketMQ 异步回写的目标。

-- 点赞关系表
CREATE TABLE post_votes (
    post_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    vote TINYINT NOT NULL,  -- 1=赞, -1=踩, 0=取消
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (post_id, user_id),
    INDEX idx_user_id (user_id)
);
-- 设计思路：
-- 每个用户对每个帖子只会有一条记录；
-- 想取消点赞，只需 UPDATE vote=0；
-- 想从踩变赞：UPDATE vote=-1 → 1；
-- 想从赞变踩：UPDATE vote=1 → -1；
-- 可以通过 COUNT(vote=1) 来计算点赞数。