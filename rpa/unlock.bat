@echo off
for /l %%i in (1,1,10000000000) do  (
	echo %%i
	timeout /t 59
	start http://127.0.0.1
) 

pause


