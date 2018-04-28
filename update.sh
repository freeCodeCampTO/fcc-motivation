go build && \
zip fcc-tweet.zip fcc-tweet && \
rm fcc-tweet && \
aws lambda update-function-code --function-name fcc-tweet --zip-file fileb://fcc-tweet.zip && \
rm fcc-tweet.zip