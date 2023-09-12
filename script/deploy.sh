

HOST=7_root
#HOST=8_root

make build; ssh $HOST "mkdir -p /usr/http_download/static;";  scp ./main $HOST:/usr/http_download/http_download ;

cat << EOF > ./init_d_sh
#!/bin/sh
cd /usr/http_download
./http_download
EOF
scp ./init_d_sh $HOST:/etc/init.d/http_download

cat << EOF > ./x.sh
chmod +x /etc/init.d/http_download
ln -s /etc/init.d/http_download /etc/rc2.d/S99http_download
EOF

scp ./x.sh  $HOST:/usr/http_download/
ssh $HOST "sh /usr/http_download/x.sh"


# apt install net-tools ;
# netstat -plant  | grep 2080