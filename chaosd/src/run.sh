cp /bin/chaosd /var/lib/chaosd
mkdir /etc/chaosd/certs/
/var/lib/chaosd/chaosd server --cert /etc/chaosd/certs/chaosd.crt --key /etc/chaosd/certs/chaosd.key

