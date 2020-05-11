proxy_host=127.0.0.1
proxy_port=20000
redirect_port=20080
iptables -t nat -A OUTPUT -p tcp --dport $redirect_port -j REDIRECT --to-port $proxy_port