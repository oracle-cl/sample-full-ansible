create user martin identified by Welcome#1#1#1 ;
GRANT CONNECT TO martin;
GRANT RESOURCE TO martin;
ALTER USER martin quota unlimited on data;
create table martin.datos (
   codigo number,
   descripcion varchar(30) );
insert into martin.datos values ( 1, 'uno' );
insert into martin.datos values ( 2, 'dos' );
insert into martin.datos values ( 3, 'tres' );
insert into martin.datos values ( 4, 'cuatro' );
insert into martin.datos values ( 5, 'cinco' );
commit ;
select * from martin.datos ;
