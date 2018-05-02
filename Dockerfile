FROM alpine:3.7

# Godep for vendoring
RUN mkdir /opt
ADD sysmon.tar.gz /opt

RUN echo "#!/bin/sh" > /opt/run.sh && echo "/opt/sysmon > /var/log/sysmon.log &" >> /opt/run.sh && echo "/bin/sh" >> /opt/run.sh && chmod a+x /opt/run.sh

WORKDIR /opt
ENTRYPOINT [ "/opt/run.sh" ]

EXPOSE 2048
