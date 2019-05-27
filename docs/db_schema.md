# db スキーマ

## ideas

| カラム名   | 型        | 属性 | 説明など                   |
| :--------- | :-------- | :--- | :------------------------- |
| id         | uint      |      | ただの ID 自動で更新される |
| created_at | time.Time |      | 作成日時                   |
| status     | string    |      | 「nowDoing」or「notYet」   |
| idea       | string    |      | アイデアの内容             |
