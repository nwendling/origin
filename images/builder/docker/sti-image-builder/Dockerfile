FROM centos:centos7

RUN yum install -y --enablerepo=centosplus epel-release gettext tar automake make git docker

ADD https://github.com/openshift/source-to-image/releases/download/v0.3/source-to-image-v0.3-957c66b-linux-amd64.tar.gz /usr/bin/sti.tar.gz
RUN cd /usr/bin && tar xzvf /usr/bin/sti.tar.gz && rm -f /usr/bin/sti.tar.gz

ADD bin/build.sh /buildroot/build.sh

WORKDIR /buildroot
CMD ["/buildroot/build.sh"]
