# Curl client aliases for Bonjour Software MyCollections API
alias authCurl='f(){ curl -i -H "Cookie: JWT=header_placeholder.payload_placeholder.signature_placeholder" -H "Content-Type:application/json" "$@"; }; f'
alias getProfile='f(){ authCurl localhost:8080/api/v1/profile; }; f'
alias getItem='f(){ authCurl localhost:8080/api/v1/item; }; f'
alias createItem='f(){ authCurl -X POST localhost:8080/api/v1/item -d "$@"; }; f'
alias updateItem='f(){ authCurl -X PUT localhost:8080/api/v1/item -d "$@"; }; f'
alias deleteItem='f(){ authCurl -X DELETE localhost:8080/api/v1/item -d "$@"; }; f'
