## Build Docker image
To build docker image with wrapped k6-playwright executable execute: 
```shell
docker build -f Dockerfile-distributed -t k6-pw-distributed:latest .
```
### Docker image usage
#### From local machine
Navigate to folder with k6-playwright spec files.  
To execute specific spec file run:
```shell
docker run -it -v $(pwd):/home/k6/scripts k6-pw-distributed:latest run scripts/<spec-name>.js
```
