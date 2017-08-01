## Minikube Filesystem Bug Repro

This project is a repro of a bug that occurs when mounting volumes with Minikube/xhyve.

###Environment
- **OS**: Mac OS Sierra - 10.12.5
- **VM Driver**: xhyve - stable 0.3.3
- **Minikube version**: 0.21.0

###Steps to Repro
1. Make sure you have Minikube running and are using the `xhyve` VM driver
    
       $ brew install minikube
       $ brew install docker-machine-driver-xhyve # make sure to follow additional instructions if installing for the first time
       $ minikube start --vm-driver=xhyve
    
2. Run `./repro.sh` to build the Dockerfile. It copies the `watch-mounted-events-in-go`
binary into the image and runs it - it uses the [fsnotify project](https://github.com/howeyc/fsnotify) to detect
any changes to the `/app/mounted` directory in the container.
`repro.sh` also runs the image in a container and mounts the local `mounted`
directory to the `/app/mounted` directory in the container. It then runs `watch-mounted-events-in-go`
and attaches to the container's output.
3. In your local filesystem, create any file or directory inside of the `mounted` directory.

**What you expect to happen**: `fsnotify` should detect and log changes to the mounted volume
to STDOUT.

**What actually happens**: `fsnotify` doesn't detect any changes to the mounted volume coming from the host
filesystem. However, you can assert that the file changes *are* propagating through to the container by
bashing into the container:
    
    # create file locally
    touch mounted/foo.txt 
    
    # copy the container ID of the repro container
    docker ps | grep repro
    
    # bash into the container
    docker exec -it <CONTAINER_ID> /bin/bash
    
    # inside the container
    ls /app/mounted
    
Running `ls /app/mounted` inside the container, you should see that `foo.txt` exists inside
the container. However the event was not recorded.