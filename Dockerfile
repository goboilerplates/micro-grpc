FROM scratch

COPY main.out /

EXPOSE 50052
ENTRYPOINT ["/main.out"]