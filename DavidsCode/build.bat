@ECHO OFF
 

:build
cls
ECHO Building.....
go build ./

rem IF NOT ERRORLEVEL 1 (
	go clean ./
	git commit -m "commit" -a
	git push
	curl.exe -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d @deploy.json "https://22971498235e45ff9ca17cb7163e8205-vp0.us.blockchain.ibm.com:5001/chaincode" > serviceKey.txt



pause

goto :build
