rm -fr dist
cd cmd/cli

echo 'building linux-amd64...'
GOOS=linux GOARCH=amd64 go build -o ../../dist/linux-amd64/apm
echo 'building linux-arm64...'
GOOS=linux GOARCH=arm64 go build -o ../../dist/linux-arm64/apm

cd ../../dist
for dir in $(ls -d *); do
    tar cfzv "$dir".tgz $dir
    rm -rf $dir
done

echo 'done!'
