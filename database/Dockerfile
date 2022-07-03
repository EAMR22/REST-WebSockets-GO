#Crea un contenedor a partir de la imagen.
FROM postgres:10.3

#COPY copia un archivo de un directorio.
COPY up.sql /docker-entrypoint-initdb.d/1.sql

# Mantiene la base de datos corriendo para utilizarla.
CMD ["potsgres"]