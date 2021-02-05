FROM imroc.tencentcloudcr.com/test/ubuntu:latest
MAINTAINER rockerchen@tencent.com
COPY benchechoclient /benchechoclient
RUN chmod a+x /benchechoclient
CMD ["/benchechoclient", "echoserver:9000"]