FROM scratch

ADD ./view /view/
COPY main.out /

EXPOSE 50052
ENTRYPOINT ["/main.out"]