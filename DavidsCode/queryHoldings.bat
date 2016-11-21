@ECHO OFF

:QUERY

for /f "delims=" %%x in (serviceKey.txt) do set Keytext=%%x

@echo. 
curl.exe -X POST --insecure --header "Content-Type: application/json" --header "Accept: application/json" -d "{\"jsonrpc\": \"2.0\",\"method\": \"query\",\"params\": {\"type\": 1,\"chaincodeID\": {\"name\": \"30aa2479d7b3b4a6b3721b0b4112b65ad5b943d15c04dc9161bf0165e599fbe0fe10dcfde0fffd19aa36dfc8149a27ad2c353ba4bc3f4774cac94238b446ad09\"},\"ctorMsg\": {\"function\": \"holdings\",\"args\": [\"Aaron\"]},\"secureContext\": \"dashboarduser_type1_0\"},\"id\": 1}" "https://d48e4ff54a324330ac90c7ed2a4ddaa2-vp1.us.blockchain.ibm.com:444\Query"
@echo. 

pause

goto :QUERY