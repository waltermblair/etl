# ETL practice
## setup
used json-generator.com to create a bunch of random people and downloaded the result as data.json  
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
provide read access:  
* create read-only user with key/id creds
* send creds to user
