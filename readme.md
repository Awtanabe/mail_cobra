#### cobra

- 初期化
cobra-clii init

- バッチ追加


```
cobra-cli add <バッチ名>

ex)
cobra-cli add BulkMail

=> cmd/BulkMail.go が生成される
```

- 使い方
https://github.com/spf13/cobra/blob/main/site/content/user_guide.md

### フラグ

https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#working-with-flags


### bash


```
docker-compose exec app sh
```

## db 権限


```
mysql -u root -p
root

GRANT ALL PRIVILEGES ON *.* TO 'user'@'%' IDENTIFIED BY 'password';
FLUSH PRIVILEGES;
```