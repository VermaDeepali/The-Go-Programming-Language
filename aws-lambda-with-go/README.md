// Steps to deploy and revoke AWS Lambda via console

1. aws iam create-role \
  --role-name lambda-ex \
  --assume-role-policy-document '{"Version": "2012-10-17",		 	 	 "Statement": [{ "Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'

2. aws iam create-role \
  --role-name lambda-ex \
  --assume-role-policy-document file://trust-policy.json

3. aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

4. aws lambda create-function \
  --function-name aws-lambda-with-go \
  --zip-file fileb://function.zip\
  --handler main \
  --runtime go1.x \
  --zip-file fileb://function.zip
  --role arn:aws:iam::<ACCOUNT_ID>:role/lambda-basic-role \

5. aws lambda invoke \
  --function-name aws-lambda-with-go \
  --cli-binary-format raw-in-base64-out \
  --paylaod '{"what is your name?": "Deep", "How old are you?": 25}'
  output.json