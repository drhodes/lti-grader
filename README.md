A Prototype grader tool that runs with jupyterhub that essentially
parses jupyter notebooks and responds with a set of form fields
automatically filled with answers and a submit button. Not sure if
this will work.

It uses the LTI protocol, so it needs two command line args, in
addition to an SSL certificate an keyfile.  Jupyterhub has good
integration with letsencrypt, but stores the cert/key in a json file
called acme.json, which it's Traefik proxy uses.

So, I found a short python script located at the following page:

https://techoverflow.net/2021/07/18/how-to-export-certificates-from-traefik-certificate-store/
(no license provided.)

That script is located in the acme-utils directory. (this may need
more documentation).

The server has two command line args that take the LTI-passport
secrets. The Makefile assumes that there is a bash file called
`env-secret.bash` that contains these two lines.

```bash
export GRADER_CONSUMER= ... 
export GRADER_SECRET= ...
```

Because they are secrets, the file `env-secret.bash` is not included
in this repository!
ls


