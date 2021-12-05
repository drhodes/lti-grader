A Prototype grader tool that runs with jupyterhub that essentially
parses jupyter notebooks and responds with a set of form fields
automatically filled with answers and a submit button. Not sure if
this will work.

It uses the LTI protocol, so it needs two command line args, in
addition to an SSL certificate an keyfile.  Jupyterhub has good
integration with letsencrypt, but stores the cert/key in a json file
called acme.json, which it's Traefik proxy uses.

So, I found a short python script located at the following page:

# https://techoverflow.net/2021/07/18/how-to-export-certificates-from-traefik-certificate-store/
# no license provided.

and have included it here






```bash
export GRADER_CONSUMER= ...
export GRADER_SECRET= ...
```
