## 1. Install Go version if don't have and set GOPATH and GOROOT
## 2.Go to directory
cd $GOPATH/Go/src/user_event_aggregation/handler

## 3. Run the commnad on terminal for part 1 
 go run . -i input.json -o output.json

 ## 4. Run the commnad on terminal for part 2
 go run . -i input.json -o output.json --update

# 5 check input.json and output.json file in directory 
$GOPATH/Go/src/user_event_aggregation/input.json
$GOPATH/Go/src/user_event_aggregation/output.json