@ECHO OFF

:QUERY

for /f "delims=" %%x in (serviceKey.txt) do set Keytext=%%x

@echo. 
curl -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"securities\",\"args\": [\"a\"]},\"secureContext\": \"dashboarduser_type1_0\"},\"id\": 1}" "https://d48e4ff54a324330ac90c7ed2a4ddaa2-vp1.us.blockchain.ibm.com:444/chaincode"
@echo. 

pause

goto :QUERY