#!/bin/sh

# 获取HOME路径
mkdir -p ~/isx
cd ~/isx
wget https://openfly.oss-cn-shanghai.aliyuncs.com/isx/isx_darwin_arm64
mv isx_darwin_arm64 isx
chmod a+x ~/isx/isx

tee -a ~/.zshrc <<-'EOF'
export PATH=$PATH:/Users/ispong/isx
EOF

source /Users/ispong/.zshrc

echo "安装成功"