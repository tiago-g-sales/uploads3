



awslocal s3api put-object --bucket business-time-nonprod --key some/folder/name/here/





sudo aws  s3 ls s3://business-time-nonprod/some/folder/name/here -region us-east-1 --endpoint-url=http://localhost:4566 --recursive --human-readable --summarize

aws s3api create-bucket --bucket bucket-mensagens --region us-east-1 --endpoint-url=http://localhost:4566

aws s3api list-buckets --query "Buckets[].Name" --region us-east-1 --endpoint-url=http://localhost:4566

aws s3 cp file0.txt s3://bucket-mensagens --region us-east-1 --endpoint-url=http://localhost:4566
