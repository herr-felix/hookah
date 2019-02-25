FROM alpine

WORKDIR /opt/hookah

COPY ./ui/dist ./ui/dist

COPY ./hookah ./hookah

ENTRYPOINT ["/bin/sh", "./hookah"]
