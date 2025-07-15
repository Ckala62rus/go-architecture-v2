@echo off
set timestamp=%date:~6,4%%date:~3,2%%date:~0,2%_%time:~0,2%%time:~3,2%%time:~6,2%
set timestamp=%timestamp: =0%
docker exec db_golang pg_dump -U postgres -d db > "docker\backup\db\backup_%timestamp%.sql"
echo Backup created: backup_%timestamp%.sql
pause