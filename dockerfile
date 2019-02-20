FROM alpine

COPY ./ui/dist /ui/dist

COPY ./hookah /hookah

ENTRYPOINT ["/hookah"]
