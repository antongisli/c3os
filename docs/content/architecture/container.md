+++
title = "Container based OS"
date = 2022-02-09T17:56:26+01:00
weight = 1
pre = "<b>- </b>"
+++

c3os is a container-based OS. 

A container based OS is an OS which is shipped via containers. Indeed, if it happens to be based on Linux (most probably) you can run the container image as well on your docker daemon. The image being being booted is the container, which contains all the required pieces in order to boot (Kernel, Initrd, Init system).

## Benefits of using containers to distribute upgrades

Container registries are already widely supported and used by anyone. 

If you are operating a Kubernetes cluster and deploying apps on top, chances are that you already have a container registry deployed somewhere and configured to store them or manage your infrastructure stack. By using container images lets you re-use the same infrastructure to propagate upgrades to the nodes and handle customizations.

Containers images can be extended after build by using standard container building practices, and seamlessly plug into your existing pipelines. c3OS allows to seamlessly upgrade to container images that are derived from other versions.
 This means that customizing a c3OS version requires just to build a standard container image with a plain `Dockerfile`, plus the bits that are actually needed.
 
If you are familiar with Dockerfiles, then you are good to go to roll your own custom OS version to provision in the nodes. That removes any friction to questions like "How do I add this package to my nodes?", or more complex ones as "How can I replace with my own Kernel?".
