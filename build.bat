@echo off
setlocal enabledelayedexpansion

cd ./bin

set BUILD_COMMAND=go build ../src

echo Building for Windows...
%BUILD_COMMAND%
del OGCAT.exe
ren src.exe OGCAT.exe


rem zip it
set rar_path="C:\Program Files\WinRAR\WinRAR.exe"
set name=ogcat
cd ..
if not exist %name% mkdir %name%
cd ogcat
xcopy /s /e /y /i /q ..\bin\* .\ >nul 2>nul
if exist %name%.zip del %name%.zip
cd ..
if exist bin\%name%.zip del bin\%name%.zip >nul 2>nul
%rar_path% a %name%.zip %name%\* >nul 2>nul

rem clean up
rmdir /s /q ogcat >nul 2>nul
move /y %name%.zip bin\%name%.zip >nul 2>nul
