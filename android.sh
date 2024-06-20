#！ /bin/bash
echo "package android so $1 ......"
rm -rf ./output/android
mkdir -p ./output/android
#gomobile bind -target android/arm64 -classpath ./ -v -o ./sotest_1.aar
gomobile bind -target android/arm64 -classpath ./ -o ./output/android/$1.aar
echo "package android so $1 finished......"