INSERT INTO articles (title, contents, username, nice, created_at) VALUES
('Go言語の基本', 'Go言語は、Googleによって開発されたオープンソースのプログラミング言語です。シンプルで効率的な構文と、並行処理のサポートが特徴です。', 'saki', 10, NOW()),
('Docker入門', 'Dockerは、アプリケーションをコンテナ化するためのプラットフォームです。これにより、開発環境と本番環境の差異を減らし、アプリケーションのデプロイを簡素化できます。', 'ken', 20, NOW()),
('MySQLの使い方', 'MySQLは、オープンソースのリレーショナルデータベース管理システムです。データベースの作成、テーブルの操作、クエリの実行など、基本的な使い方を学ぶことができます。', 'manaka', 15, NOW());

INSERT INTO articles (title, contents, username, nice) VALUES
('Go言語の基本', 'Go言語は、Googleによって開発されたオープンソースのプログラミング言語です。シンプルで効率的な構文と、並行処理のサポートが特徴です。', 'saki', 10),
('Docker入門', 'Dockerは、アプリケーションをコンテナ化するためのプラットフォームです。これにより、開発環境と本番環境の差異を減らし、アプリケーションのデプロイを簡素化できます。', 'ken', 20),
('MySQLの使い方', 'MySQLは、オープンソースのリレーショナルデータベース管理システムです。データベースの作成、テーブルの操作、クエリの実行など、基本的な使い方を学ぶことができます。', 'manaka', 15);

INSERT INTO comments (article_id, message, created_at) VALUES
(1, 'Go言語はシンプルで効率的な構文が魅力です。', NOW()),
(1, '並行処理のサポートが便利ですね。', NOW()),
(2, 'Dockerを使うとデプロイが簡単になります。', NOW()),
(3, 'MySQLの基本操作を学ぶのに役立ちます。', NOW());

INSERT INTO comments (article_id, message) VALUES
(1, 'GoよりRustだろ!'),
(2, 'たれぞうです');
