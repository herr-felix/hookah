FROM alpine

WORKDIR /opt/hookah

COPY ./ui/dist ./ui/dist

COPY ./ui/views ./ui/views

COPY ./hookah ./hookah

ENTRYPOINT ["./hookah"]
