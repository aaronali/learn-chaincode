@ECHO  off

:QUERY

for /f "delims=" %%x in (serviceKey.txt) do set Keytext=%%x

@echo. 
curl -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"holdings\",\"args\": [\"user_type1_0\"]},\"secureContext\": \"dashboarduser_type1_0\"},\"id\": 1}" "https://22971498235e45ff9ca17cb7163e8205-vp0.us.blockchain.ibm.com:5001"
  @echo. 
curl -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"holdings\",\"args\": [\"user_type1_1\"]},\"secureContext\": \"dashboarduser_type1_0\"},\"id\": 1}" "https://22971498235e45ff9ca17cb7163e8205-vp0.us.blockchain.ibm.com:5001" 
@echo. 
curl -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"holdings\",\"args\": [\"Bank\"]},\"secureContext\": \"dashboarduser_type1_0\"},\"id\": 1}" "https://22971498235e45ff9ca17cb7163e8205-vp0.us.blockchain.ibm.com:5001"
@echo.  
curl -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"holdings\",\"args\": [\"Aaron\"]},\"secureContext\": \"dashboarduser_type1_0\"},\"id\": 1}" "https://22971498235e45ff9ca17cb7163e8205-vp0.us.blockchain.ibm.com:5001"
@echo.  
curl -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"%Keytext:~52,128%\"},\"ctorMsg\": {\"function\": \"holdings\",\"args\": [\"David\"]},\"secureContext\": \"dashboarduser_type1_0\"},\"id\": 1}" "https://22971498235e45ff9ca17cb7163e8205-vp0.us.blockchain.ibm.com:5001"
@echo. 
 

goto :QUERY