-- Sync the database contents with the fixtures in models/data.go.
-- data.go uses time.Now(), so timestamps are generated at migration time.

START TRANSACTION;

DELETE FROM comments;
DELETE FROM articles;

INSERT INTO articles (
    article_id,
    title,
    contents,
    username,
    nice,
    created_at
) VALUES
    (1, 'first article', 'This is the test article.', 'saki', 1, NOW()),
    (2, 'second article', 'This is the test article.', 'saki', 2, NOW());

INSERT INTO comments (
    comment_id,
    article_id,
    message,
    created_at
) VALUES
    (1, 1, 'test comment1', NOW()),
    (2, 1, 'second comment', NOW());

ALTER TABLE articles AUTO_INCREMENT = 3;
ALTER TABLE comments AUTO_INCREMENT = 3;

COMMIT;
