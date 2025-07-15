@echo off
set /p filename="Enter backup filename (without path): "
type "docker\backup\db\%filename%" | docker exec -i db_golang psql -U postgres -d db
echo Restore completed!
pause