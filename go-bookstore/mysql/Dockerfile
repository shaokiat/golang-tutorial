# Use a MySQL 8.0 base image
FROM mysql:8.0

# Set the root password for the MySQL server
ENV MYSQL_ROOT_PASSWORD=password

# Copy the SQL file to initialize the database
COPY init.sql /docker-entrypoint-initdb.d/

# Set the character set and collation for the database
RUN echo "[mysqld]\n\
character-set-server=utf8mb4\n\
collation-server=utf8mb4_unicode_ci\n\
" >> /etc/mysql/conf.d/charset.cnf

# Run following Docker commands
# docker build -t my-mysql-image .
# docker run --name my-mysql-container -p 3306:3306 -d my-mysql-image
