## A Simple Demo for Cosmos Sdk

#### usage:

* get clone this project to your local repo 

* environment set:

    cd ~/

    vim .bashrc or vim .zshrc 在最后添加以下内容:

```$shell
        export GOPATH=$HOME/go
        export GOBIN=\$GOPATH/bin
        export PATH=\$PATH:\$GOBIN 
        export GO111MODULE=on （最重要）
```     
 
    保存退出后:
  
            source ~/.bashrc or ~/.zshrc



* 重启goland 并运行 make install 命令:
 
    when the command finished on success, there will generate two binary 
    files named "nsd" and "nscli" at $GOBIN 
    
* cd $GOBIN
* ./nsd help and ./nscli help to see what you do next
