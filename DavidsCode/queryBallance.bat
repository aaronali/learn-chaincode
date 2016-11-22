@ECHO OFF

:QUERY

for /f "delims=" %%x in (serviceKey.txt) do set Keytext=%%x

@echo. 
curl -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"ballance\",\"args\": [\"Aaron\"]},\"secureContext\": \"user_type1_0\"},\"id\": 1}" "https://22971498235e45ff9ca17cb7163e8205-vp0.us.blockchain.ibm.com:5001/chaincode"
@echo. 

 

goto :QUERY