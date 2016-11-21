@ECHO OFF
 

:build
cls
ECHO Building.....
go build ./

rem IF NOT ERRORLEVEL 1 (
	go clean ./
	git commit -m "commit" -a
	git push
	curl.exe -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d @deploy.json "https://d48e4ff54a324330ac90c7ed2a4ddaa2-vp1.us.blockchain.ibm.com:444/chaincode" > serviceKey.txt
rem )


pause

goto :build
