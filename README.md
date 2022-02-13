# Avah webhook server

practice go programming project from zero

## go command

 - enable dependency tracking : go mod init $name
 - run code : 
	 - go run .
 - run code with env var : 
	 - $varName1 = $varValue1, $varName2 = $varValue2 go run .
- add missing and remove unused modules
	- go mod tidy

## Project Structure

best practice project structure
 **Flat structure**
 project structure for small applications (example: **[go-yaml/yaml](https://github.com/go-yaml/yaml)**)
||||
|-|-|-|
|application/|||
||main.go||
||main_test.go||
||utils.go||
||utils_test.go||

 **Modularization**
 project structure for medium or large sized applications (example: **[google/go-cloud](https://github.com/google/go-cloud)**)
|||||
|-|-|-|-|
|application/|||
||main.go||
||main_test.go||
||user/||
|||user.go|
|||login.go|
|||registration.go|
||articles/||
|||articles.go|
||utils/||
|||common_utils.go|

**Mature project structure**
traditional structure - feature 'internal' and 'pkg' folder which encapsulate some of inner workings of the projects (example: -   **[Kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)**)
|||||
|-|-|-|-|
|application/||||
||cmd|||
|||cmd.go||
||pkg|||
|||controller||
||||cronjob/cronjob_controller.go|
||||cronjob/cronjob_controller_test.go|
||||endpoint/endpoint_controller.go|
||||endpoint/endpoint_controller_test.go|
|||scheduler||
||||eventhandler.go|
||||eventhandler_test.go|
||||scheduler.go|
||||scheduler_test.go|

**key word**
|||
|-|-|
|defer|to delay the execution of the function until the nearby functions returns|

**tools**
- node
	- brew install node
- ngrok : public urls for exposing local web server (global install)
	- npm install ngrok -g
	- ngrok http $service_port