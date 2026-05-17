FROM localhost:5001/cmaestro-python:3.12

LABEL AUTHORS="Cactus Maestro (c) https://github.com/qerdinger/cactus-maestro - MIT License"
LABEL IMG_LICENSE="MIT"

WORKDIR /

COPY requirements.txt .

RUN ["pip3.12", "install", "-r", "requirements.txt"]


COPY serverless.py .

CMD [ "serverless.py" ]