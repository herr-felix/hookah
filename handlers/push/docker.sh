
image_tag=$(echo "$1" | tr '[A-Z]' '[a-z]')

docker build . -t $image_tag

# docker push -u "user" -p "secret" wherever/whatever

docker rmi $image_tag