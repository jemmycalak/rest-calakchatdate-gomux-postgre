create table t_user(
    iduser SERIAL PRIMARY KEY NOT NULL,
     firtsname VARCHAR(50) NOT NULL,
      lastname VARCHAR(50) NOT NULL,
       email VARCHAR(150) NOT NULL,
        password VARCHAR(500) NOT NULL,
         imageprofile VARCHAR(500),
          createAt TIMESTAMP WITH TIME ZONE,
           updateAt TIMESTAMP WITH TIME ZONE);