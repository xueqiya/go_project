docker build -t go_project . && \
docker run -d -p 8888:8888 --rm --name go_project go_project