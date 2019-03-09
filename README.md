# ETL practice
Goal is to write an application that does the following:
1) fetch json files from s3 bucket marshall-test
2) store original files in couchdb
3) call lambda function to process each file
4) store results in mysql 

Ulimate goal is to have everything deployed in the cloud (including containerized application and databases)

Skills/Tools:
* AWS:
  * API
  * cli
  * IAM (optional)
* RESTful services
* Docker
* git, github
* Postman (optional)

## setup (already done)
### data
use json-generator.com to create a bunch of random people and downloaded the result as data.json  
split data.json sequence of objects into one file per object:  
```
mkdir ./people
jq -c '.[]' data.json | split -l 1 --additional-suffix=.json - people/person_
```
upload files to s3:  
* create admin user  
* install and configure aws cli with admin creds as default profile  
* use admin credentials to create s3 bucket and upload files  
```
aws s3api create-bucket --bucket marshall-test --region us-east-1
aws s3 cp people s3://marshall-test/ --recursive
aws s3 ls marshall-test/
``` 

### function
create lambda funcion:
* wrote Go program extractBalances
* deployed function @ https://uglmid0m9c.execute-api.us-east-1.amazonaws.com/beta

### AWS user
provide user read/invoke access to s3 bucket and lambda function:  
* create read-only user with key/id creds
* send creds to user

## project (in progress)
