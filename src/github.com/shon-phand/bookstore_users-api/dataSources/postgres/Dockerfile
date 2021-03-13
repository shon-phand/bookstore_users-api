  
FROM postgres

COPY ./init.sql/ /docker-entrypoint-initdb.d/
ENV POSTGRES_PASSWORD shon1234
ENV POSTGRES_USER postgres
ENV POSTGRES_DB postgres
