
cd $GOPATH/src/github.com/aaronali/learn-chaincode/start 
go build ./ >buildLog.txt
call :buildifEmpty "buildLog.txt"
goto nonCompile

:buildifEmpty
pause
if %~z1 eq 0 goto commit
goto nonCompile
:commit 
git status 
git add --all 
git commit -m "Compiled my code" 
git push
exit /b

:nonCompile
echo non compite
type "buildLog.txt"