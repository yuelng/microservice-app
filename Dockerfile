FROM golang:1.7

# add custom base package
RUN rm -rf /go/src/base
ADD ./src/base /go/src/base/

ADD ./src/guestbook /go/src/guestbook/
ADD ./src/greeter /go/src/greeter/

RUN cd ./src/guestbook && go get -v
RUN cd ./src/greeter && go get -v

# ADD 的注意事项
# 如果源路径是文件,且目标路径以/结尾,则直接将文件拷到目标路径下
# 如果源路径是文件,且目标路径不是以/结尾,则会把目标路径当做一个文件,如果目标文件是存在的文件,会用源文件覆盖它,内容是源文件,文件名是目标文件名
# 如果源路径是目录,且目标路径不存在,则docker会自动创建目录,并将目录拷贝进来
# 如果源文件是归档文件,则docker会自动帮解压

# COPY 类似但是不会自动解压