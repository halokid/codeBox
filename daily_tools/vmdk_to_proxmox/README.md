###迁移vmdk文件到proxmox私有云



1. xlsxtoini目录根据test.xlsx文件的格式生成虚机的ini配置文件

   ```shell
   cd ./xlstoini
   go test
   ```

2. 把源码目录放入gopath，进入目录

   ```shell
   go get
   go build
   ./vmdk_to_proxmox
   ```

3. 程序会自动读取  第一步生成的 ini配置文件，生成proxmox的虚机


有问题可以邮件我






