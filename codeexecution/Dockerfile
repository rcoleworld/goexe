#publicly available docker image "python" on docker hub will be pulled

FROM python

#creating directory helloworld in container (linux machine)

RUN mkdir /home/test

#copying helloworld.py from local directory to container's helloworld folder

COPY ./tests/test.py /home/test/test.py

#running helloworld.py in container

CMD python /home/test/test.py