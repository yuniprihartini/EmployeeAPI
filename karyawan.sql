create database karyawan;
use karyawan;
select database();
create table kartap(
	kartap_id varchar(100),
    nama_lengkap varchar (100),
    alamat varchar (100),
    tempat_lahir varchar(100),
    tanggal_lahir date,
    gaji_pokok int, 
    tunjangan int,
    status_karyawan varchar (5),
     primary key(kartap_id));
    select * from kartap;
    
create table karkontrak(
	karkontrak_id varchar(100),
    nama_lengkap varchar (100),
    alamat varchar (100),
    tempat_lahir varchar(100),
    tanggal_lahir date,
    gaji_pokok int, 
    status_karyawan varchar (5),
     primary key(karkontrak_id));
    select * from karkontrak;
    
    insert into kartap (kartap_id, nama_lengkap, alamat, tempat_lahir, tanggal_lahir, gaji_pokok, tunjangan, status_karyawan) values
    ("KRTP1", "Muhammad Ahmadin", "Jalan Sidosermo Surabaya", "Surabaya", "1997-06-05",
    "5000000", "1000000", "aktif"),
     ("KRTP2", "Siti Zahidah", "Jalan Ketintang Madya Surabaya", "Lamongan", "1980-04-03",
    "6000000", "1000000", "aktif");
    
    insert into karkontrak (karkontrak_id, nama_lengkap, alamat, tempat_lahir, tanggal_lahir, gaji_pokok, status_karyawan) values
    ("KRKON1", "Firaun Amal", "Jalan Panjang Jiwo Surabaya", "Jombang", "1989-06-05",
    "3500000", "aktif"),
     ("KRKON2", "Zulaekha", "Jalan Tugu Pahlawan Surabaya", "Madura", "1990-04-03",
    "3250000","aktif");
     