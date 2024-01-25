## Config database


At the command line, log into your DBMS, as in the following example for MySQL.
```bash
mysql -u root -p
```

At the mysql command prompt, create a database.
```bash
create database recordings;
```

Change to the database you just created so you can add tables.
```bash
use recordings;
```

From the mysql command prompt, run the script **create-tables.sql**

```bash
source /path/to/create-tables.sql
```

At your DBMS command prompt, use a SELECT statement to verify you’ve successfully created the table with data.
```bash
select * from album;
```
