@echo off
setlocal enabledelayedexpansion

if "%1"=="" (
    goto help_message
)

set type=%1
if "%2"=="" (
    set nr_clients=1
) else (
    set nr_clients=%2
)
set build=False
set clear=False
set log_client=False
set sequential=False

:parse_flags
if "%3"=="" goto end_parse
set flag=%3
if "%flag:~0,1%"=="-" (
    set flag=%flag:~1%
) else (
    goto help_message
)
for /l %%i in (0,1,2) do (
    set "char=!flag!"
    set "char=!char:~%%i,1!"
    if "!char!"=="" goto end_parse
    if "!char!"=="b" (
        set build=True
    ) else if "!char!"=="c" (
        set clear=True
    ) else if "!char!"=="s" (
        set sequential=True
    ) else if "!char!"=="l" (
        set log_client=True
    ) else (
        goto help_message
    )
)
shift
goto parse_flags
:end_parse

if %build%==True (
    if %clear%==True (
        cls && call build.bat %type% && cls
    ) else (
        call build.bat %type%
    )
) else (
    if %clear%==True (
        cls
    )
)

cd ../bin
echo Starting server...
start server_%type%.exe %log_client%
timeout /t 2 /nobreak
if %nr_clients%==1 (
    echo Starting one client...
) else (
    echo Starting %nr_clients% clients...
)
if %sequential%==True (
    for /L %%i in (1,1,%nr_clients%) do (
        call client_%type%.exe
    )
) else (
    call client_%type%.exe %nr_clients%
)
echo Press any key to stop the server...
pause >nul
echo Stopping server!
taskkill /F /IM server_%type%.exe >nul 2>&1
endlocal
exit /b 0

:help_message
echo Usage: %0 [api/grpc] [NR_CLIENTS] [PARAMS]
echo Param:
echo  -h     Display this help message
echo  -b     Build the project
echo  -c     Clear the console
echo  -s     Call the clients sequentially
echo  -l     Log client messages in the server console
exit /b
