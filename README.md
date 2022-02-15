A server that listens for requests from openedx lti_consumer
xblocks. The request has three pieces of data:

- edx_anon_id 
  This is a 32 character anonymized id that openedx uses for each user
  
- staff_answers
  This is json map that 

- notebook_name
  this is the name of the jupyter notebook to be graded.

----

The server has two command line args that take the LTI-passport
secrets. The Makefile assumes that there is a bash file called
`env-secret.bash` that contains these two lines.


```bash
export GRADER_CONSUMER= ... 
export GRADER_SECRET= ...
```

the file `env-secret.bash` is not included in this repository!  

----

note: eventually, `env-secret.bash` will be generated automatically at
some point.

