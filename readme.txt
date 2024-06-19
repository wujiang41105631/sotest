-- 打包安卓so 文件：
    gomobile bind -target android/arm64 -classpath ./ -v -o ./sotest_1.aar
-- 打包ios文件 :
    gomobile bind -target ios/arm64 -classpath ./ -v -o ./sotest.1.xcframework